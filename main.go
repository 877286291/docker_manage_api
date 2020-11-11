package main

import (
	"DockerUI/app/router"
	_ "fmt"
)

func main() {
	r := router.InitRouter()
	_ = r.Run(":8080")
}
