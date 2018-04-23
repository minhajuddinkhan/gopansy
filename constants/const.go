package constants

type key int

const (
	//DbKey DB context key
	DbKey key = iota
	//DbType type of database
	DbType string = "postgres"
)
