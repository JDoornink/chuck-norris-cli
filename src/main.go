/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"chuck-norris/cmd"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	Auth AuthOpts
}

type AuthOpts struct {
	Audience  string
	Disabled  bool
	EngineURL string
	Issuer    string
}

func main() {
	// https://api.chucknorris.io/jokes/random
	// https://geek-jokes.sameerkumar.website/api?format=json

	fmt.Println("main.go call")

	config, err := initConfig()
	if err != nil {
		//log.Fatal().Err(err)
		log.Fatal(err)
	}

	fmt.Println("config auth audience = ", config.Auth.Audience)
	cmd.Execute()

}

func initConfig() (Configuration, error) {
	fmt.Println("initConfig in main.yaml")
	v := viper.New()
	config := Configuration{}

	// v.AddConfigPath("./configs"): Use this if your config is in the src folder
	v.AddConfigPath("../configs")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	//v.AddConfigPath("..")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		// If the config file is not found it's not necessarily an error.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
		fmt.Println("Something isnt right")
	}

	config.Auth.Audience = v.GetString("auth.audience")

	fmt.Println("config.Auth.Audience = ", config.Auth.Audience)
	return config, nil
}
