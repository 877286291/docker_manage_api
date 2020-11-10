package main

import (
	"DockerUI/app/router"
)

func main() {
	r := router.InitRouter()
	_ = r.Run(":8080")
}
