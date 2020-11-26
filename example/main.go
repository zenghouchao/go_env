package main

import (
	"fmt"
	env "github.com/zenghouchao/go_env"
)

func main() {
	var params env.EnvParams
	params.ParseEnvFile(".env")
	database := params.GetSection("DATABASE", "DATABASE")
	fmt.Println(database)

}