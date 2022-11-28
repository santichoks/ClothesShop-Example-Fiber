package configs

type Configs struct {
	App        Fiber
	PostgreSQL PostgreSQL
}

type Fiber struct {
	Host string
	Port string
}

type PostgreSQL struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	SSLMode  string
}
