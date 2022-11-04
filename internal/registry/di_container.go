package registry

import (
	"golang-base-structure/config"
	"golang-base-structure/internal/adapter"
	"golang-base-structure/internal/api"
	"golang-base-structure/internal/helper/memcache"
	"golang-base-structure/internal/usecase"
	"sync"

	goCache "github.com/patrickmn/go-cache"
	"github.com/sarulabs/di"
)

// DI Name
const (
	// config
	ConfigDIName string = "Config"
	// helper
	oracleDBDIName    string = "OracleDB"
	sqlServerDBDIName string = "SQLServerDB"
	postgresDBDIName  string = "PostgresDB"
	memCacheDIName    string = "MemCacheHelper"
	// repository
	oracleRepositoryDIName    string = "OracleRepository"
	postgresRepositoryDIName  string = "PostgresRepository"
	sqlServerRepositoryDIName string = "SQLServerRepository"
	// adapter
	mgpsAdapterDIName string = "MGPSAdapter"
	// useCase
	applicationUseCaseDIName string = "ApplicationUseCase"
	userUseCaseDIName        string = "UserUseCase"
	transactionUseCaseDIName string = "TransactionUseCase"
	// api
	APIDIName string = "API"
)

var (
	buildOnce sync.Once
	builder   *di.Builder
	container di.Container
)

// BuildDIContainer build DI container
func BuildDIContainer() {
	buildOnce.Do(func() {
		builder, _ = di.NewBuilder()
		if err := buildConfigs(); err != nil {
			panic(err)
		}
		if err := buildHelpers(); err != nil {
			panic(err)
		}
		if err := buildRepositories(); err != nil {
			panic(err)
		}
		if err := buildAdapters(); err != nil {
			panic(err)
		}
		if err := buildUsecases(); err != nil {
			panic(err)
		}
		if err := buildAPIs(); err != nil {
			panic(err)
		}
		container = builder.Build()
	})
}

// GetDependency gets dependency from DI container
func GetDependency(dependencyName string) interface{} {
	return container.Get(dependencyName)
}

// CleanDependency cleans dependency
func CleanDependency() error {
	return container.Clean()
}

func buildConfigs() error {
	defs := []di.Def{}

	configDef := di.Def{
		Name:  ConfigDIName,
		Scope: di.App,
		Build: func(di di.Container) (interface{}, error) {
			return config.Load()
		},
		Close: func(obj interface{}) error {
			return nil
		},
	}
	defs = append(defs, configDef)
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildHelpers() error {
	defs := []di.Def{}
	// Please read in README to implement Oracle DB
	/*
		dbOracleDef := di.Def{
			Name:  oracleDBDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				return db.NewOracleDBHelper(cfg.DBOracle.Host, cfg.DBOracle.Port, cfg.DBOracle.Username,
					cfg.DBOracle.Password, cfg.DBOracle.Database[0]), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}
		defs = append(defs, dbOracleDef)
	*/
	/*
		dbPostgresDef := di.Def{
			Name:  postgresDBDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				return db.NewPostgresDBHelper(cfg.DBPostgres.Host, cfg.DBPostgres.Port, cfg.DBPostgres.Username,
				cfg.DBPostgres.Password, cfg.DBPostgres.Database[0]), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}
		defs = append(defs, dbPostgresDef)

		dbSQLServerDef := di.Def{
			Name:  sqlServerDBDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				return db.NewSQLServerDBHelper(cfg.DBSQLServer.Host, cfg.DBSQLServer.Port, cfg.DBSQLServer.Username,
					cfg.DBSQLServer.Password, cfg.DBSQLServer.Database[0]), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}
		defs = append(defs, dbSQLServerDef)
	*/
	memCacheDef := di.Def{
		Name:  memCacheDIName,
		Scope: di.App,
		Build: func(di di.Container) (interface{}, error) {
			return memcache.NewMemCacheHelper(), nil
		},
		Close: func(obj interface{}) error {
			return nil
		},
	}
	defs = append(defs, memCacheDef)

	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildRepositories() error {
	defs := []di.Def{}

	// Please read in README to implement Oracle DB
	/*
		oracleRepositoryDef := di.Def{
			Name:  oracleRepositoryDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				oracleDB := di.Get(oracleDBDIName).(db.DBHelper)
				return repository.NewOracleRepository(oracleDB), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}
		defs = append(defs, oracleRepositoryDef)
	*/
	/*
		postgresRepositoryDef := di.Def{
			Name:  postgresRepositoryDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				postgresDB := di.Get(postgresDBDIName).(db.DBHelper)
				return repository.NewPostgresRepository(postgresDB), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}
		defs = append(defs, postgresRepositoryDef)

		sqlServerRepositoryDef := di.Def{
			Name:  sqlServerRepositoryDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				sqlServerDB := di.Get(sqlServerDBDIName).(db.DBHelper)
				return repository.NewSQLServerRepository(sqlServerDB), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}
		defs = append(defs, sqlServerRepositoryDef)
	*/
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildAdapters() error {
	defs := []di.Def{}

	mgpsAdapterDef := di.Def{
		Name:  mgpsAdapterDIName,
		Scope: di.App,
		Build: func(di di.Container) (interface{}, error) {
			cfg := di.Get(ConfigDIName).(*config.Config)
			return adapter.NewMGPSAdapter(cfg), nil
		},
		Close: func(obj interface{}) error {
			return nil
		},
	}
	defs = append(defs, mgpsAdapterDef)

	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildUsecases() error {
	defs := []di.Def{}

	applicationUseCaseDef := di.Def{
		Name:  applicationUseCaseDIName,
		Scope: di.App,
		Build: func(di di.Container) (interface{}, error) {
			cfg := di.Get(ConfigDIName).(*config.Config)
			memCache := di.Get(memCacheDIName).(*goCache.Cache)
			//oracleRepository := di.Get(oracleDBDIName).(repository.OracleRepository)
			return usecase.NewApplicationUseCase(cfg, memCache), nil
		},
		Close: func(obj interface{}) error {
			return nil
		},
	}
	defs = append(defs, applicationUseCaseDef)

	transactionUsecaseDef := di.Def{
		Name:  transactionUseCaseDIName,
		Scope: di.App,
		Build: func(di di.Container) (interface{}, error) {
			cfg := di.Get(ConfigDIName).(*config.Config)
			//sqlRepository := di.Get(sqlServerDBDIName).(repository.SQLServerRepository)
			//return usecase.NewTransactionUseCase(cfg, sqlRepository), nil
			return usecase.NewTransactionUseCase(cfg, nil), nil
		},
		Close: func(obj interface{}) error {
			return nil
		},
	}
	defs = append(defs, transactionUsecaseDef)

	userUsecaseDef := di.Def{
		Name:  userUseCaseDIName,
		Scope: di.App,
		Build: func(di di.Container) (interface{}, error) {
			cfg := di.Get(ConfigDIName).(*config.Config)
			//postgresRepository := di.Get(postgresDBDIName).(repository.PostgresRepository)
			//return usecase.NewUserUseCase(cfg, postgresRepository), nil
			return usecase.NewUserUseCase(cfg, nil), nil
		},
		Close: func(obj interface{}) error {
			return nil
		},
	}
	defs = append(defs, userUsecaseDef)

	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildAPIs() error {
	defs := []di.Def{}

	apiDef := di.Def{
		Name:  APIDIName,
		Scope: di.App,
		Build: func(di di.Container) (interface{}, error) {
			cfg := di.Get(ConfigDIName).(*config.Config)
			applicationUsecase := di.Get(applicationUseCaseDIName).(usecase.ApplicationUseCase)
			userUsecase := di.Get(userUseCaseDIName).(usecase.UserUseCase)
			transactionUsecase := di.Get(transactionUseCaseDIName).(usecase.TransactionUseCase)

			return api.NewAPI(cfg, applicationUsecase, userUsecase, transactionUsecase), nil
		},
		Close: func(obj interface{}) error {
			return nil
		},
	}
	defs = append(defs, apiDef)

	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}
