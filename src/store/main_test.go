package store

var dbConfig = PostgresConfig{
	Host:     "127.0.0.1",
	Port:     5432,
	Schema:   "pairswap",
	User:     "root",
	Password: "password",
	Option:   "charset=utf8&parseTime=True&loc=Local",
}
