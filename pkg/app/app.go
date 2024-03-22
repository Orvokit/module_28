package app

import (
	"fmt"
	"io"
	"strconv"
)

type Storage interface {
	Put(string) bool
	Get(string) string
	Students()
}

type App struct {
	Repository Storage
}

func (a *App) Run() {
	for {
		a.printStudents()
		if student, ok := a.inputNextStudent(); ok {
			a.storeStudent(student)
		} else {
			break
		}
	}
}

func (a *App) printStudents() {
	fmt.Println("Список введенных студентов с оценками:")
	a.Repository.Students()
}

func (a *App) inputNextStudent() (string, bool) {
	for {
		fmt.Print("Введите данные студента или Ctrl + Z для завершения: ")
		var inputName string
		var inputAge, inputGrade int
		_, err := fmt.Scanln(&inputName, &inputAge, &inputGrade)

		inputInfo := inputName + " " + strconv.Itoa(inputAge) + " " + strconv.Itoa(inputGrade)

		if err == io.EOF {
			var getStudent, name string
			fmt.Println("Если хотите найти студента в списке, введите 'find', если нет, еще раз введите Ctrl + Z")
			fmt.Scan(&getStudent)
			if getStudent == "find" {
				fmt.Print("Введите имя студента: ")
				fmt.Scan(&name)
				fmt.Println(a.Repository.Get(name))
				return "end", false
			} else {
				a.printStudents()
				return "end", false
			}
		} else {
			return inputInfo, true
		}
	}
}

func (a *App) storeStudent(student string) {
	msg := "Студент уже присутствует в коллекции\n"
	if ok := a.Repository.Put(student); ok {
		msg = "Студент успешно добавлен\n"
	}
	fmt.Printf(msg, student)
}
