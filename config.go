package main

type configuration struct {
	Port             string
	ConnectionString string
	Addr             string
}

var (
	//ConfDev Configuration
	ConfDev = configuration{
		Port:             "8080",
		ConnectionString: "host=db user=pansy-user dbname=pansy-go password=s3cr3tp4ssw0rd sslmode=disable",
		Addr:             "0.0.0.0:8080",
	}
)
