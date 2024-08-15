package config

type MainConfig struct {
	ServiceName   string
	Server        ServerConfig        `fig:"server"`
	DBMigration   DbMigrate           `fig:"dbmigrate"`
	Rdbms         RdbmsConfig         `fig:"rdbms"`
	Authorization AuthorizationConfig `fig:"authConfig"`
}

type (
	ServerConfig struct {
		Rest struct {
			ListenAddress  string `fig:"listenAddress"`
			Port           int    `fig:"port"`
			DefaultTimeout int    `fig:"defaultTimeout"`
			ReadTimeout    int    `fig:"readTimeout"`
			WriteTimeout   int    `fig:"writeTimeout"`
			EnableSwagger  bool   `fig:"enableSwagger"`
			APIKey         string `fig:"APIKey" json:"APIKey"`
		}
	}

	DbMigrate struct {
		MigrationConfig MigrationConfig `fig:"app"`
	}

	MigrationConfig struct {
		Dsn    string `fig:"dsn"`
		Driver string `fig:"driver"`
	}

	RdbmsConfig struct {
		DBConfig DBConfig `fig:"app"`
	}

	DBConfig struct {
		DSN             string `fig:"dsn"`
		Driver          string `fig:"driver"`
		MaxOpenConns    int    `fig:"maxOpenConns"`
		MaxIdleConns    int    `fig:"maxIdleConns"`
		ConnMaxLifetime int    `fig:"connMaxLifetime"`
		Retry           int    `fig:"retry"`
	}

	AuthorizationConfig struct {
		EnableBasicAuth bool   `fig:"enableBasicAuth"`
		Username        string `fig:"username"`
		Password        string `fig:"password"`
	}
)
