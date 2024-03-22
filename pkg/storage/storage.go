package storage

import (
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

type MemStorage struct {
	students map[string]*Student
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		students: make(map[string]*Student),
	}
}

func (ms *MemStorage) Put(inputInfo string) bool {
	inputInfoSplited := strings.Split(inputInfo, " ")
	var StudentInfo Student
	StudentInfo.Name = inputInfoSplited[0]
	StudentInfo.Age, _ = strconv.Atoi(inputInfoSplited[1])
	StudentInfo.Grade, _ = strconv.Atoi(inputInfoSplited[2])

	if ms.contains(StudentInfo.Name) {
		return false
	}
	ms.students[StudentInfo.Name] = &StudentInfo
	return true
}

func (ms *MemStorage) Get(name string) string {
	getStudent, ok := ms.students[name]
	if ok == true {
		return getStudent.Name + " " + strconv.Itoa(getStudent.Age) + " " + strconv.Itoa(getStudent.Grade)
	} else {
		return "Такого студента нет в коллекции"
	}
}

func (ms *MemStorage) Students() {
	for key, v := range ms.students {
		fmt.Printf("%s: %v %d %d\n", key, (v.Name), (v.Age), (v.Grade))
	}
}

func (ms *MemStorage) contains(stName string) bool {
	for key := range ms.students {
		if ms.students[key].Name == stName {
			return true
		}
	}
	return false
}
