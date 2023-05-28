package main

import (
	"github.com/NetfluxESIR/backend-client/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
