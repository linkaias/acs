package bootstrap

import (
	"APICallerStats/model"
	"github.com/spf13/viper"
	"log"
)

func NewEnv() *model.Env {
	//_ = gotenv.Load()
	env := &model.Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return env
}
