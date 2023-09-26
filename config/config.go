package config

import (
	"github.com/jinzhu/configor"
	"path/filepath"
	"runtime"
	"sync"
)

var config Config
var once sync.Once

// Config represent application configuration
type Config struct {
	Layerfolders          []string
	ColorVariations       []string
	Desiredpicturesnumber int
	Outputfolder          string
}

// LoadConfig Load configuration from environment
func LoadConfig() {
	_, b, _, _ := runtime.Caller(0)
	str := filepath.Dir(b)

	// singleton pattern
	once.Do(
		func() {
			configor.New(&configor.Config{Verbose: true}).Load(&config, str+"/../resources/config.yml")
		})
}

// GetConfig enforce getting config checking that it was been previously initialized
func GetConfig() Config {

	if nil == &config {
		panic("Before you can access config you need to initialize it calling LoadConfig()")
	}
	return config
}
