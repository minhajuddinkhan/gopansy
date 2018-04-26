package constants

type key int

const (
	//DbKey DB context key
	DbKey key = iota
	//DbType type of database
	DbType string = "postgres"

	//JwtSecret JwtSecret
	JwtSecret string = "tryingtobeagopher"

	//Authorization Authorization
	Authorization string = "auth"

	//HashRounds HashRounds
	HashRounds int = 10
)
