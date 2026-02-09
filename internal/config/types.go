package config

// Path: internal/config/types.go

type Config struct {
Memory    MemoryConfig     `yaml:"memory"`
Listeners []ListenerConfig `yaml:"listeners"`
}

type MemoryConfig struct {
Banks []MemoryBank `yaml:"banks"`
}

type MemoryBank struct {
ID            uint64     `yaml:"id"`
Name          string     `yaml:"name"`
UnitWidthBits uint32     `yaml:"unit_width_bits"`
Units         uint64     `yaml:"units"`
Init          *CIDRBlock `yaml:"init,omitempty"`
Access        *CIDRBlock `yaml:"access,omitempty"`
}

type CIDRBlock struct {
AllowCIDR []string `yaml:"allow_cidr"`
}

type ListenerConfig struct {
Name string `yaml:"name"`
Bind string `yaml:"bind"`
}
