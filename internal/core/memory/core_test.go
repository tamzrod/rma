package memory

import (
	"bytes"
	"errors"
	"math"
	"sync"
	"testing"
)

func TestReadWriteUnalignedAndIsolation(t *testing.T) {
	c := NewCore()
	if err := c.BuildFromSpecs([]BankSpec{
		{ID: 1, UnitWidthBits: 3, Units: 5},
		{ID: 2, UnitWidthBits: 3, Units: 5},
	}); err != nil {
		t.Fatal(err)
	}
	c.Seal()
	if err := c.WriteUnits(1, 1, 3, []byte{0xA5, 0x01}); err != nil {
		t.Fatal(err)
	}
	got, err := c.ReadUnits(1, 1, 3)
	if err != nil || !bytes.Equal(got, []byte{0xA5, 0x01}) {
		t.Fatalf("unaligned read=%v err=%v", got, err)
	}
	got[0] = 0
	again, err := c.ReadUnits(1, 1, 3)
	if err != nil || again[0] != 0xA5 {
		t.Fatalf("read returned aliased memory: %v %v", again, err)
	}
	other, err := c.ReadUnits(2, 0, 5)
	if err != nil || !bytes.Equal(other, []byte{0, 0}) {
		t.Fatalf("bank isolation failed: %v %v", other, err)
	}
	if err := c.BuildFromSpecs([]BankSpec{{ID: 3, UnitWidthBits: 8, Units: 1}}); !errors.Is(err, ErrSealed) {
		t.Fatalf("sealed layout change: %v", err)
	}
}

func TestInvalidAccessAndNoMutation(t *testing.T) {
	c := NewCore()
	if err := c.BuildFromSpecs([]BankSpec{{ID: 1, UnitWidthBits: 16, Units: 2}}); err != nil {
		t.Fatal(err)
	}
	if err := c.WriteUnits(1, 0, 2, []byte{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}
	checks := []struct {
		name string
		err  error
		want error
	}{
		{"unknown bank", c.WriteUnits(3, 0, 1, []byte{1, 2}), ErrBankNotFound},
		{"zero count", c.WriteUnits(1, 0, 0, nil), ErrInvalidCount},
		{"out of range", c.WriteUnits(1, 1, 2, []byte{1, 2, 3, 4}), ErrOutOfBounds},
		{"address overflow", c.WriteUnits(1, math.MaxUint64, 2, []byte{1, 2, 3, 4}), ErrOutOfBounds},
		{"start at units", c.WriteUnits(1, 2, 1, []byte{1, 2}), ErrOutOfBounds},
		{"short payload", c.WriteUnits(1, 0, 2, []byte{1}), ErrPayloadSizeMismatch},
		{"long payload", c.WriteUnits(1, 0, 2, []byte{1, 2, 3, 4, 5}), ErrPayloadSizeMismatch},
	}
	for _, test := range checks {
		if !errors.Is(test.err, test.want) {
			t.Errorf("%s: got %v want %v", test.name, test.err, test.want)
		}
	}
	if _, err := c.ReadUnits(3, 0, 1); !errors.Is(err, ErrBankNotFound) {
		t.Fatalf("missing read bank: %v", err)
	}
	if _, err := c.ReadUnits(1, 0, 0); !errors.Is(err, ErrInvalidCount) {
		t.Fatalf("zero read count: %v", err)
	}
	if _, err := c.ReadUnits(1, math.MaxUint64, 2); !errors.Is(err, ErrOutOfBounds) {
		t.Fatalf("overflow read: %v", err)
	}
	data, err := c.ReadUnits(1, 0, 2)
	if err != nil || !bytes.Equal(data, []byte{1, 2, 3, 4}) {
		t.Fatalf("rejected writes mutated state: %v %v", data, err)
	}
}

func TestInvalidSpecsAndAtomicBuild(t *testing.T) {
	c := NewCore()
	cases := []struct {
		spec BankSpec
		want error
	}{
		{BankSpec{ID: 3, UnitWidthBits: 0, Units: 1}, ErrInvalidUnitWidth},
		{BankSpec{ID: 3, UnitWidthBits: 8, Units: 0}, ErrInvalidUnits},
		{BankSpec{ID: 3, UnitWidthBits: math.MaxUint32, Units: math.MaxUint64}, ErrMemoryTooLarge},
		{BankSpec{ID: 3, UnitWidthBits: 8, Units: maxBankBytes + 1}, ErrMemoryTooLarge},
	}
	for _, tc := range cases {
		if err := c.BuildFromSpecs([]BankSpec{tc.spec}); !errors.Is(err, tc.want) {
			t.Fatalf("spec %+v got %v want %v", tc.spec, err, tc.want)
		}
	}
	err := c.BuildFromSpecs([]BankSpec{
		{ID: 10, UnitWidthBits: 8, Units: 4},
		{ID: 10, UnitWidthBits: 8, Units: 4},
	})
	if !errors.Is(err, ErrDuplicateBank) {
		t.Fatalf("duplicate batch: %v", err)
	}
	if _, err := c.ReadUnits(10, 0, 1); !errors.Is(err, ErrBankNotFound) {
		t.Fatalf("partial batch leaked: %v", err)
	}
	if err := c.BuildFromSpecs([]BankSpec{{ID: 10, UnitWidthBits: 8, Units: 4}}); err != nil {
		t.Fatal(err)
	}
	err = c.BuildFromSpecs([]BankSpec{
		{ID: 11, UnitWidthBits: 8, Units: 4},
		{ID: 10, UnitWidthBits: 8, Units: 4},
	})
	if !errors.Is(err, ErrDuplicateBank) {
		t.Fatalf("existing duplicate: %v", err)
	}
	if _, err := c.ReadUnits(11, 0, 1); !errors.Is(err, ErrBankNotFound) {
		t.Fatalf("partial existing-id batch leaked: %v", err)
	}
}

func TestConcurrentAtomicReadWrite(t *testing.T) {
	c := NewCore()
	if err := c.BuildFromSpecs([]BankSpec{{ID: 1, UnitWidthBits: 8, Units: 64}}); err != nil {
		t.Fatal(err)
	}
	c.Seal()
	const writers, readers, iterations = 4, 4, 300
	var wg sync.WaitGroup
	errCh := make(chan error, writers+readers)
	for i := 0; i < writers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			data := bytes.Repeat([]byte{byte(i + 1)}, 64)
			for j := 0; j < iterations; j++ {
				if err := c.WriteUnits(1, 0, 64, data); err != nil {
					errCh <- err
					return
				}
			}
		}(i)
	}
	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				got, err := c.ReadUnits(1, 0, 64)
				if err != nil {
					errCh <- err
					return
				}
				for _, v := range got {
					if v != got[0] {
						errCh <- errors.New("torn read")
						return
					}
				}
			}
		}()
	}
	wg.Wait()
	close(errCh)
	for err := range errCh {
		t.Error(err)
	}
}

func TestUnalignedUnitsKeepNeighborBits(t *testing.T) {
	c := NewCore()
	if err := c.BuildFromSpecs([]BankSpec{{ID: 1, UnitWidthBits: 3, Units: 4}}); err != nil {
		t.Fatal(err)
	}
	for _, w := range []struct {
		start uint64
		data  byte
	}{{0, 0x07}, {2, 0x01}} {
		if err := c.WriteUnits(1, w.start, 1, []byte{w.data}); err != nil {
			t.Fatal(err)
		}
	}
	full, err := c.ReadUnits(1, 0, 4)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(full, []byte{0x47, 0x00}) {
		t.Fatalf("packed bits=%08b %08b want 01000111 00000000", full[0], full[1])
	}
	const mask = 0x07
	for unit, want := range []byte{0x07, 0x00, 0x01, 0x00} {
		got, err := c.ReadUnits(1, uint64(unit), 1)
		if err != nil {
			t.Fatal(err)
		}
		if got[0]&mask != want {
			t.Fatalf("unit %d=%03b want %03b", unit, got[0]&mask, want)
		}
	}
}

func TestSingleBitWidthUnits(t *testing.T) {
	c := NewCore()
	if err := c.BuildFromSpecs([]BankSpec{{ID: 1, UnitWidthBits: 1, Units: 16}}); err != nil {
		t.Fatal(err)
	}
	if err := c.WriteUnits(1, 4, 4, []byte{0x0F}); err != nil {
		t.Fatal(err)
	}
	got, err := c.ReadUnits(1, 0, 16)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, []byte{0xF0, 0x00}) {
		t.Fatalf("single-bit write=%08b %08b want 11110000 00000000", got[0], got[1])
	}
	if err := c.WriteUnits(1, 3, 1, []byte{0x01}); err != nil {
		t.Fatal(err)
	}
	got, err = c.ReadUnits(1, 3, 1)
	if err != nil || got[0]&0x01 != 0x01 {
		t.Fatalf("single bit read=%v err=%v", got, err)
	}
	got, err = c.ReadUnits(1, 4, 4)
	if err != nil || got[0]&0x0F != 0x0F {
		t.Fatalf("neighbor bits disturbed=%08b err=%v", got[0], err)
	}
}

func TestAggregateLimitFailsBeforeAllocation(t *testing.T) {
	c := NewCore()
	var specs []BankSpec
	for i := 0; i < 5; i++ {
		specs = append(specs, BankSpec{
			ID: MemoryID(i + 1), UnitWidthBits: 8, Units: maxBankBytes,
		})
	}
	if err := c.BuildFromSpecs(specs); !errors.Is(err, ErrMemoryTooLarge) {
		t.Fatalf("aggregate cap: %v", err)
	}
	for _, spec := range specs {
		if _, err := c.ReadUnits(spec.ID, 0, 1); !errors.Is(err, ErrBankNotFound) {
			t.Fatalf("partial aggregate build: id=%d err=%v", spec.ID, err)
		}
	}
	if got := bitsToBytes(math.MaxUint64); got != math.MaxUint64/8+1 {
		t.Fatalf("bitsToBytes overflow: %d", got)
	}
}
