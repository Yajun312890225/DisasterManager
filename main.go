package main

import (
	"DisasterManager/conf"
	_ "DisasterManager/docs"
	"DisasterManager/router"
	"os"
)

func main() {
	conf.Init()
	r := router.InitRouter()
	r.Run(os.Getenv("PORT"))
}
