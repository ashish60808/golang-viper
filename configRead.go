package main

import (
	"github.com/spf13/viper"
	"fmt"

	)
func main() {
viper.SetConfigName("config1") // name of config1.yaml file (without extension)
viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
viper.AddConfigPath(".")      // path to look for the config1.yaml file in

err := viper.ReadInConfig()
if err != nil {
fmt.Println("Config not found...")
} else {
name := viper.GetString("name")
fmt.Println("Config found, name = ", name)


position := viper.GetString("position")
fmt.Println("Config found, position = ", position)

	auth := viper.GetStringMapString("auth")
	fmt.Println("Config found, auth = ", auth["password"])
	fmt.Printf("Reading config for auth = %#v\n", auth)
	fmt.Println("Config found, auth = ", auth["usernamebm "])


}
}