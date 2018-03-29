package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/alastairruhm/go-rixcloud/rixcloud"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("rixcloud")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("$HOME/.rixcloud")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(fmt.Errorf("Fatal error config file: %s", err))
		os.Exit(1)
	}

	username := viper.GetString("username")
	password := viper.GetString("password")
	serviceid := viper.GetString("serviceid")

	client := rixcloud.NewClient(&http.Client{}, username, password)

	v, _, err := client.Profiles.GetTrafficData(serviceid)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	upload, _ := strconv.Atoi(v.Upload)
	fmt.Printf("upload  : %5d M\n", upload>>20)
	dowload, _ := strconv.Atoi(v.Download)
	fmt.Printf("download: %5d M\n", dowload>>20)
	fmt.Printf("used    : %5d M\n", (upload+dowload)>>20)
	total, _ := strconv.Atoi(v.Total)
	fmt.Printf("total   : %5d M\n", total>>20)
}
