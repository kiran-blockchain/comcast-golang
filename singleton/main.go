package main

import (
    "fmt"
    "sync"
)

// Config represents the configuration settings
type Config struct {
    Setting1 string
    Setting2 int
}

var (
    instance *Config
    once     sync.Once
)

// GetConfig returns the singleton instance of Config
func GetConfig() *Config {
    once.Do(func() {
        instance = &Config{
            Setting1: "default_value",
            Setting2: 42,
        }
    })
    return instance
}

func main() {
    config1 := GetConfig()
    config2 := GetConfig()

    fmt.Printf("Config1: %+v\n", config1)
    fmt.Printf("Config2: %+v\n", config2)

    // Verify that config1 and config2 are indeed the same instance
    if config1 == config2 {
        fmt.Println("Both instances are the same.")
    } else {
        fmt.Println("Instances are different.")
    }
}
