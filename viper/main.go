package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("")
	viper.SetConfigType("toml")
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	fmt.Println(viper.GetBool("year"))
	fmt.Println(viper.GetStringSlice("byear"))
	viper.Set("year",false)
	fmt.Println(viper.GetBool("year"))
	viper.WriteConfig()
}