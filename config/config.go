package config

//Configuration Configuration
type Configuration struct {
	Port             string
	ConnectionString string
	Addr             string
}

//GetEnvs get environments
func GetEnvs() map[string]string {

	envs := make(map[string]string)
	envs["dev"] = "dev.json"
	return envs

}
