package main

import (
	"os"
	"prakerja4/configs"
	"prakerja4/routes"
)

func init() {
	configs.LoadEnv()
	configs.ConnectDatabase()
}

func main() {
	e := routes.Init()
	e.Start(":" + getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // port default jika tidak ada PORT yang ditentukan
	}
	return port
}

