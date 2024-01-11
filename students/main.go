package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func main() {

	//ОБъявление  пустого хранилища
	studentStorage := make(map[string]*Student)

	// Основная логика  бесконечного запроса строки у пользователя
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите текст: ")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Проверка на ctrl+d с учетом есть ли что нибудь в хранилище или сразу был такой ввод (eof)
		if err == io.EOF {
			if len(studentStorage) > 0 {
				fmt.Println("Студенты из хранилища:")
				for _, student := range studentStorage {
					fmt.Println(student.Name, student.Age, student.Grade)
				}
				break
			} else {
				fmt.Println("Хранилище студентов пусто.")
				break
			}
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения: %v\n", err)
			continue
		}

		// Если все ок - формирование  студента.
		// Если возникет ошибка в функции newStudent - вывод ошибки и сброс на начало
		tempStudent, err := newStudent(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Если до этого все ок - добавление студента в хранилище
		studentStorage[tempStudent.Name] = tempStudent
	}

}

// Функцию создающая студента и отдающая структуру
func newStudent(s string) (*Student, error) {
	tempSlice := strings.Split(s, " ")

	// Проверки на длину, на числа в возрасте и оценке
	if len(tempSlice) != 3 {
		return nil, errors.New("возможно опечатка, заново")
	}

	tempAge, err := strconv.Atoi(tempSlice[1])
	if err != nil || tempAge < 0 || tempAge > 100 {
		return nil, errors.New("не указаны числа, указаны неверно")
	}

	tempGrade, err := strconv.Atoi(tempSlice[2])
	if err != nil || tempGrade < 0 {
		return nil, errors.New("не указаны числа, указаны неверно")
	}

	// Если все ок - формируем структуру студента
	newStudent := Student{
		Name:  tempSlice[0],
		Age:   tempAge,
		Grade: tempGrade,
	}

	return &newStudent, nil
}
