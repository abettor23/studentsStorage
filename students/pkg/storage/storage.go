package storage

import (
	"fmt"
	"studentsApp/pkg/student"
)

// Storage определяет интерфейс для управления данными студентов.
type Storage interface {
	Get() string
	Put(st *student.Student)
	StudentsInfo() map[string]*student.Student
}

// StudentsStorage реализует Storage, храня информацию о студентах.
type StudentsStorage struct {
	Students map[string]*student.Student
}

// StudentsInfo возвращает текущую информацию о всех студентах.
func (s *StudentsStorage) StudentsInfo() map[string]*student.Student {
	return s.Students
}

// Get формирует и возвращает отформатированную строку с информацией о всех студентах.
func (s *StudentsStorage) Get() string {
	var info string
	if len(s.StudentsInfo()) > 0 {
		info = fmt.Sprintf("Студенты из хранилища:\n\n")
		for _, student := range s.StudentsInfo() {
			info = fmt.Sprintf("%s%s %d %d\n", info, student.Name, student.Age, student.Grade)
		}
		return info
	} else {
		info = "Данные о студентах отсутствуют."
		return info
	}
}

// Put добавляет или обновляет информацию о студенте в хранилище.
func (s *StudentsStorage) Put(st *student.Student) {
	s.Students[st.Name] = st
}
