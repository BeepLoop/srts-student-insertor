package insertor

type Config struct {
	Source     string
	Output     string
	Database   string
	Limit      bool
	LimitValue int
	Program    string
}

func Initialize(config *Config) {
	//
}
