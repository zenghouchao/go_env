package main

import (
	"fmt"
	env "github.com/zenghouchao/go_env"
	"os"
	"runtime"
)

var osType = runtime.GOOS

func main() {
	dir, _ := os.Getwd()
	if osType == "windows" {
		dir = dir + "\\.env"
	} else if osType == "linux" {
		dir = dir + "/.env"
	}
	var params env.EnvParams
	params.ParseEnvFile(dir)
	database := params.GetSection("DATABASE", "DATABASE")
	fmt.Println(database)

}
