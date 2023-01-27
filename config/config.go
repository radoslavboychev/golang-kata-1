package config

import (
	liberror "github.com/echocat/golang-kata-1/v1/errors"
	"github.com/spf13/viper"
)

type config struct {
	MagazinesFile string `mastructure:"MAGAZINES_FILE"`
	BooksFile     string `mapstructure:"BOOKS_FILE"`
	AuthorEmail   string `mapstructure:"AUTHOR_EMAIL"`
	FindByISBN    string `mapstructure:"ISBN"`
	FindByTitle   string `mapstructure:"FIND_TITLE"`
}

var conf *config

// LoadConfig reads a .env file and loads the environment variables
func LoadConfig() (config, error) {

	viper.AddConfigPath("../.././config")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return config{}, liberror.ErrFailedToOpenFile
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return config{}, liberror.ErrFailedToOpenFile
	}

	return *conf, nil
}
