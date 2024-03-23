package config

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	Secrets  SecretConfig   `json:"secrets"`
	Stage    string         `json:"stage"`
	SMTP     SMTPConfig     `json:"smtp"`
}

type PostgreConfig struct {
	Host   string `json:"host"`
	DBName string `json:"db_name"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type ServerConfig struct {
	GRPC PortConfig `json:"grpc"`
	Rest PortConfig `json:"rest"`
}

type PortConfig struct {
	Port string `json:"port"`
	Name string `json:"name"`
}

type DatabaseConfig struct {
	PostgreDB PostgreConfig `json:"postgre"`
	RedisDB   RedisConfig   `json:"redis"`
}

type SecretConfig struct {
	JWTSecret string `json:"jwt_secret"`
}

type SMTPConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}
