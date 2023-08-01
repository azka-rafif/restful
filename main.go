package main

import "go-mini/internal"

// go:generate go run github.com/swaggo/swag/cmd/swag init
func main() {
	internal.SetupAndServe()
}
