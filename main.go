package main

import (
	"boilerplate.com/v1/cmd/dependency"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed Load ENV")
	}
	fx.New(dependency.Module, fx.NopLogger).Run()
}
