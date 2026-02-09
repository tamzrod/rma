package config

// Path: internal/config/loader.go

import (
"os"

"gopkg.in/yaml.v3"
)

func Load(path string) (*Config, error) {
raw, err := os.ReadFile(path)
if err != nil {
return nil, err
}

var cfg Config
if err := yaml.Unmarshal(raw, &cfg); err != nil {
return nil, err
}

if err := Validate(&cfg); err != nil {
return nil, err
}

return &cfg, nil
}
