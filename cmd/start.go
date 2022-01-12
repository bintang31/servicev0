package cmd

import (
	"fmt"
	"go/build"
	"os"
	"servicev0/routes"
	log "servicev0/src/infrastructure/logger"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// AppPath application path
	AppPath string
)

func init() {
	rootCmd.AddCommand(startCmd)
	loadEnvVars()
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start mobileloket http service",
	Run: func(cmd *cobra.Command, args []string) {
		routes.API()
		log.InfoLogHandler("Start Service")
	},
}

func loadEnvVars() {

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	// Bind OS environment variable
	viper.SetEnvPrefix("mobileloket")
	viper.BindEnv("env")
	viper.BindEnv("app_path") // bind LOKET_APP_PATH varibale

	if viper.Get("env") == "development" {
		viper.SetConfigName("staging")
		dir, _ := os.Getwd()
		AppPath = dir

	} else if viper.Get("env") == "production" {
		viper.SetConfigName("production")
		dir, _ := os.Getwd()
		AppPath = dir

	} else if viper.Get("env") == "testing" {
		viper.SetConfigName("testing")
		AppPath = viper.GetString("app_path")

	} else if viper.Get("env") == "production" {
		viper.SetConfigName("production")
		dir, _ := os.Getwd()
		AppPath = dir

	} else {
		viper.SetConfigName("config")
		dir, _ := os.Getwd()
		AppPath = dir
	}

	viper.SetConfigType("json")
	viper.AddConfigPath(AppPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

}
