package config

var DBConnStr = "postgres://" +
	"admin" + ":" + "root" +
	"@" + "localhost" + ":" + "5432" +
	"/" + "postgres" + "?sslmode=disable"
