package mylogger

type Config struct {
	isPersistent   bool
	persistentPath string
	debugMode      bool
}

func NewConfig(isPersistent bool, persistentPath string, debugMode bool) *Config {
	return &Config{
		isPersistent:   isPersistent,
		persistentPath: persistentPath,
		debugMode:      debugMode,
	}
}

func DefaultConfig() *Config {
	return &Config{
		isPersistent:   false,
		persistentPath: "",
		debugMode:      true,
	}
}
