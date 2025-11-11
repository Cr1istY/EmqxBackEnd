package main

import (
	"EmqxBackEnd/router"
)

func main() {
	r := router.Setup()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
