package config

// DatabaseConfiguration describes DB options in config.yml
type DatabaseConfiguration struct {
	Dbname   string
	Username string
	Password string
	Host     string
	Port     int32
}
