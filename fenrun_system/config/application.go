package config

import "github.com/spf13/viper"

type Application struct {
	IsInit        bool
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          string
	Name          string
	JwtSecret     string
	LogPath       string
	Env           string
	EnvMsg        string

	RuntimeRootPath string
	ExportSavePath  string

	CustomerRatio float64
	SalesmanRatio float64

	DataCrawlerPeriod int64
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		IsInit:            cfg.GetBool("isInit"),
		ReadTimeout:       cfg.GetInt("readTimeout"),
		WriterTimeout:     cfg.GetInt("writerTimeout"),
		Host:              cfg.GetString("host"),
		Port:              cfg.GetString("port"),
		Name:              cfg.GetString("name"),
		JwtSecret:         cfg.GetString("jwtSecret"),
		LogPath:           cfg.GetString("logPath"),
		Env:               cfg.GetString("env"),
		EnvMsg:            cfg.GetString("envMsg"),
		RuntimeRootPath:   cfg.GetString("runTimeRootPath"),
		ExportSavePath:    cfg.GetString("exportSavePath"),
		CustomerRatio:     cfg.GetFloat64("customerratio"),
		SalesmanRatio:     cfg.GetFloat64("salesmanratio"),
		DataCrawlerPeriod: cfg.GetInt64("datacrawlerperiod"),
	}
}

var ApplicationConfig = new(Application)
