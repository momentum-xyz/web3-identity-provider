package database

// Config : structure to hold MYSQL configuration
type Config struct {
	Name      string `yaml:"database" envconfig:"DB_DATABASE"`
	Host      string `yaml:"host" envconfig:"DB_HOST"`
	Port      uint   `yaml:"port" envconfig:"DB_PORT"`
	Username  string `yaml:"username" envconfig:"DB_USERNAME"`
	Password  string `yaml:"password" envconfig:"DB_PASSWORD"`
	Migrate   bool   `yaml:"migrate" envconfig:"DB_MIGRATE"`
	Dialect   string `yaml:"dialect" envconfig:"DB_DIALECT"`
	DSNParams string `yaml:"dsn-params" envconfig:"DB_DSN_PARAMS"`
}

func (x *Config) Init() {
	x.Name = "web3_idp_dev"
	x.Host = "localhost"
	x.Password = ""
	x.Username = "root"
	x.Port = 5432
	x.Dialect = "postgres"
	x.DSNParams = ""
}
