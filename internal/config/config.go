package config

import (
	"encoding/json"
	"example.com/receiptcheck/internal/model"
	"example.com/receiptcheck/internal/rules"
	"fmt"
	"os"
)

type Config struct {
	Database string                 `json:"database"`
	Rules    []model.ValidationRule `json:"rules"`
	Strict   bool                   `json:"strict"`
}

func Default() Config {
	return Config{Database: "receipts.db", Rules: rules.DefaultRules(), Strict: true}
}
func Load(path string) (Config, error) {
	b, e := os.ReadFile(path)
	if e != nil {
		return Config{}, e
	}
	var c Config
	if e = json.Unmarshal(b, &c); e != nil {
		return Config{}, e
	}
	return c, Validate(c)
}
func Validate(c Config) error {
	if c.Database == "" {
		return fmt.Errorf("database required")
	}
	if len(c.Rules) == 0 {
		return fmt.Errorf("rules required")
	}
	return nil
}
func Save(path string, c Config) error {
	b, e := json.MarshalIndent(c, "", "  ")
	if e != nil {
		return e
	}
	return os.WriteFile(path, b, 0644)
}
