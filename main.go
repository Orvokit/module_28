package main

import (
	"module_28/pkg/app"
	"module_28/pkg/storage"
)

func main() {
	//repository := &StubStorage{}
	repository := storage.NewMemStorage()
	app := app.App{Repository: repository}
	app.Run()
}
