package config

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

var defaultConfig = []byte(`
environment: D
app: golang-base-structure
http_address: 10001
db_oracle:
  host: 127.0.0.1
  port: 1521
  username: abc
  password: Abv@123
  database: 
    - etl
db_sql_server:
  host: 127.0.0.1
  port: 1433
  username: bdf
  password: Aaa@123
  database: 
    - ApplicationGateway
db_postgres:
  host: 127.0.0.1
  port: 5432
  username: venusvc
  password: Vc@123
  database: 
    - cmsvc
`)

type (
	Config struct {
		Base `mapstructure:",squash"`
	}

	Base struct {
		App         string   `mapstructure:"app"`
		Environment string   `mapstructure:"environment"`
		HTTPAddress int      `mapstructure:"http_address"`
		DBOracle    DBServer `mapstructure:"db_oracle"`
		DBSQLServer DBServer `mapstructure:"db_sql_server"`
		DBPostgres  DBServer `mapstructure:"db_postgres"`
		MGPSURL     string   `mapstructure:"mgps_url"`
	}

	DBServer struct {
		Host     string   `mapstructure:"host"`
		Port     int      `mapstructure:"port"`
		Username string   `mapstructure:"username"`
		Password string   `mapstructure:"password"`
		Database []string `mapstructure:"database"`
	}
)

// Load config
func Load() (*Config, error) {
	var cfg = &Config{}
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		return nil, err
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
