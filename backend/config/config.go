package config

type MainConfig struct {
	ServiceName string
	Server      ServerConfig    `fig:"server"`
	Snowflake   SnowflakeConfig `fig:"snowflake"`
	DBMigration DbMigrate       `fig:"dbmigrate"`
	Rdbms       RdbmsConfig     `fig:"rdbms"`
	Redis       RedisConfig     `fig:"redis"`
	Cron        CronConfig      `fig:"cron"`
	PubSub      PubSubConfig    `fig:"pubsub"`
	Rpc         RpcConfig       `fig:"rpc"`
	Tracer      TracerConfig    `fig:"tracer"`
	Keycloak    KeycloakConfig  `fig:"keycloak"`
	App         AppConfig       `fig:"app"`
}

type (
	ServerConfig struct {
		Grpc struct {
			Port int `fig:"port"`
		}

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

	SnowflakeConfig struct {
		Epoch int64 `fig:"epoch"`
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

	RedisConfig struct {
		Host     string `fig:"host"`
		Port     int    `fig:"port"`
		Username string `fig:"username"`
		Password string `fig:"password"`
		DB       int    `fig:"db"`
	}

	CronConfig struct {
		Jobs CronJobsConfig `yaml:"jobs"`
	}

	CronJobsConfig struct {
		UpdateStatusOrderToPartner CronJobDetailConfig `yaml:"updateStatusOrderToPartner"`
	}

	CronJobDetailConfig struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Spec        string `yaml:"spec"`
		Interval    int    `yaml:"interval"`
	}

	/*
		one publisher can use more service
		publisher name prefix: service-domain-action
		example: ScStockUpdatePub
		Sc = service that publishes messages (service in app or service another app)
		stock = domain in service
		Update = action
	*/
	PublisherList struct {
	}

	/*
		one subscriber can onliy use one service and app
		subcriber name prefix: servicepub-domain-action-servicesub
		example: MdProductCreateCompletedScSub
		Md = service that publishes messages
		Product = domain in service
		CreateCompleted = action
		Sc = the service that receives the message
	*/

	SubscriberList struct {
	}

	PubSubConfig struct {
		ProjectId      string         `fig:"projectId"`
		AuthJsonPath   string         `fig:"authJsonPath"`
		PublisherList  PublisherList  `fig:"publisherList"`
		SubscriberList SubscriberList `fig:"subscriberList"`
		Enabled        bool           `fig:"enabled"`
	}

	RpcConfig struct {
		TmsMasterDataServiceClientConfig struct {
			Host string `fig:"host"`
			Port int    `fig:"port"`
		} `fig:"tmsMasterDataServiceClientConfig"`
		CustomerPortalServiceClientConfig struct {
			Host string `fig:"host"`
			Port int    `fig:"port"`
		} `fig:"customerPortalServiceClientConfig"`
	}

	TracerConfig struct {
		Jaeger JaegerConfig `fig:"jaeger"`
	}

	JaegerConfig struct {
		CollectorUrl string `fig:"collectorUrl"`
	}

	KeycloakConfig struct {
		BaseUrl                      string `yaml:"baseUrl"`
		InternalEmployeeRealm        string `yaml:"internalEmployeeRealm"`
		InternalEmployeeClientId     string `yaml:"internalEmployeeClientId"`
		InternalEmployeeClientSecret string `yaml:"internalEmployeeClientSecret"`
		ExternalEmployeeRealm        string `yaml:"externalEmployeeRealm"`
		ExternalEmployeeClientId     string `yaml:"externalEmployeeClientId"`
		ExternalEmployeeClientSecret string `yaml:"externalEmployeeClientSecret"`
		CustomerPortalRealm          string `yaml:"customerPortalRealm"`
		CustomerPortalClientId       string `yaml:"customerPortalClientId"`
		CustomerPortalClientSecret   string `yaml:"customerPortalClientSecret"`
	}

	AppConfig struct {
		TmsUrl            string `yaml:"tmsUrl"`
		CustomerPortalUrl string `yaml:"customerPortalUrl"`
	}
)
