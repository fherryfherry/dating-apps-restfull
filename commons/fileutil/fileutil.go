package fileutil

import (
	"github.com/spf13/viper"
)

func GetFullUrl(fileName string) string {
	if fileName == "" {
		return ""
	}
	return viper.GetString("base_url_upload") + "/" + fileName
}
