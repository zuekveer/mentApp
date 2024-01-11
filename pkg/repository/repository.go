package repository

type Repository struct {
	config *Config
}

type Config struct {
	Password string `env:"POSTGRES_PASSWORD"`
	User     string `env:"POSTGRES_USER"`
	Database string `env:"POSTGRES_DATABASE"`
	Port     int8   `env:"POSTGRES_PORT"`
	Host     string `env:"POSTGRES_HOST"`
}

func New(cfg *Config) *Repository {
	return &Repository{
		config: cfg,
	}
}

func (strt *Repository) Start() error {
	return nil
}

func (stp *Repository) Shutdown() error {
	return nil
}