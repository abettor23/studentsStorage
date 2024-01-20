package main

import (
	"studentsApp/pkg/app"
	"studentsApp/pkg/storage"
	"studentsApp/pkg/student"
)

// main инициализирует приложение и запускает его.
func main() {
	repository := &storage.StudentsStorage{Students: make(map[string]*student.Student)}
	app := app.NewApp(repository)
	app.Run()
}
