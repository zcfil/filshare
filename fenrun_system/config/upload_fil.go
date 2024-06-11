package config

import "github.com/spf13/viper"

type UpLoadFile struct {
	UploadPath string
	MaxSize    int64
	URL        string
	Total      int
}

func InItUpLoadFile(cfg *viper.Viper) *UpLoadFile {
	return &UpLoadFile{
		UploadPath: cfg.GetString("UploadPath"),
		MaxSize:    cfg.GetInt64("MaxSize"),
		URL:        cfg.GetString("URL"),
		Total:      cfg.GetInt("Total"),
	}
}

var UpdateFile = new(UpLoadFile)
