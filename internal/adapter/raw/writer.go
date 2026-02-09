package raw

// Path: internal/adapter/raw/writer.go
//
// Raw adapter implementing READ and WRITE packets.
// RAM behavior only. No semantics. No policy.
//
// Packet format (big-endian):
//
// WRITE
//   [ OP=0x01 ][ MID ][ START(4) ][ COUNT(4) ][ DATA... ]
//
// READ
//   [ OP=0x02 ][ MID ][ START(4) ][ COUNT(4) ]
//
// READ RESPONSE
//   [ OP=0x02 ][ MID ][ START(4) ][ COUNT(4) ][ DATA... ]

import (
	"encoding/binary"
	"io"
	"log"
	"net"

	"github.com/tamzrod/rma/internal/core/memory"
)

const (
	opWrite = 0x01
	opRead  = 0x02
)

type Writer struct {
	Core *memory.Core
}

func New(core *memory.Core) *Writer {
	return &Writer{Core: core}
}

func (w *Writer) Handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 64*1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("[raw-adapter] read error: %v", err)
			}
			return
		}

		if n < 10 {
			// OP(1) + MID(1) + START(4) + COUNT(4)
			continue
		}

		op := buf[0]
		memID := memory.MemoryID(buf[1])
		start := binary.BigEndian.Uint32(buf[2:6])
		count := binary.BigEndian.Uint32(buf[6:10])

		switch op {

		case opWrite:
			data := buf[10:n]
			if err := w.Core.WriteUnits(memID, uint64(start), uint32(count), data); err != nil {
				log.Printf("[raw-adapter] WRITE error: %v", err)
				return
			}

			log.Printf(
				"[raw-adapter] WRITE mem=%d start=%d count=%d bytes=%d",
				memID, start, count, len(data),
			)

		case opRead:
			data, err := w.Core.ReadUnits(memID, uint64(start), uint32(count))
			if err != nil {
				log.Printf("[raw-adapter] READ error: %v", err)
				return
			}

			respLen := 10 + len(data)
			resp := make([]byte, respLen)

			resp[0] = opRead
			resp[1] = byte(memID)
			binary.BigEndian.PutUint32(resp[2:6], start)
			binary.BigEndian.PutUint32(resp[6:10], count)
			copy(resp[10:], data)

			if _, err := conn.Write(resp); err != nil {
				log.Printf("[raw-adapter] READ response write error: %v", err)
				return
			}

			log.Printf(
				"[raw-adapter] READ mem=%d start=%d count=%d bytes=%d",
				memID, start, count, len(data),
			)

		default:
			// Unknown OP  ignore (RAM doesn't complain)
			continue
		}
	}
}
