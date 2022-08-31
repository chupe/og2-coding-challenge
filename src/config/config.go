package config

type Config struct {
	DB           DB
	MongoExpress MongoExpress
	Factories    Factories
}

type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type MongoExpress struct {
	User     string
	Password string
}

type Factories struct {
	Iron   []Level
	Copper []Level
	Gold   []Level
}

type Level struct {
	Level           int
	Production      int // per minute
	UpgradeDuration int // duration in seconds
	Cost            cost
}

type cost struct {
	Iron   int
	Copper int
	Gold   int
}
