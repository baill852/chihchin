package lib

type Config struct {
	Ws    Ws
	Redis Redis
}

type Ws struct {
	Host string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	Db       int
}

//TODO load from config.json
func LoadConfig() *Config {
	return &Config{
		Ws{
			Host: "wss://stream.yshyqxx.com/stream",
		},
		Redis{
			Host:     "127.0.0.1",
			Port:     6379,
			Password: "",
			Db:       0,
		},
	}
}
