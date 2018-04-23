package config

//Configuration Configuration
type Configuration struct {
	Port             string
	ConnectionString string
	Addr             string
}

var (
	envPaths = map[string]string{
		"dev": "dev.json",
	}
)

//GetEnvPath get environments
func GetEnvPath(env string) string {
	return envPaths[env]
}

var config Configuration

//SetConfig SetConfig
func SetConfig(conf Configuration) {

	if &config != nil {
		config = conf
	}
}

//GetConfig GetConfig
func GetConfig() Configuration {

	return config
}
