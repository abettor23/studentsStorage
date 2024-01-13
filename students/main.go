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

// Набор методов Storage, содержащий логические операции программы.
type Storage interface {
	GetInfo() (string, error)
	StringValidate(string) (string, int, int, error)
	NewStudent(name string, age int, grade int) *Student
	PutStudent(st *Student)
	GetAllStudents() map[string]*Student
}

// Структура хранилища информации о студнетах.
type StudentsStorage struct {
	students map[string]*Student
}

// Метод хранилища, который считывает пользовательский ввод строчно.
func (s *StudentsStorage) GetInfo() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите текст: ")
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err
}

// Метод хранилища, который проверяет корректность ввода посредством образования слайса tempSlice.
// Если ввод не соответствует ожиданиям, выводится ошибка с подсказкой.
// Если ввод проходит проверку, метод возвращает данные о студенте (имя, возраст, оценку)
func (s *StudentsStorage) StringValidate(str string) (string, int, int, error) {
	tempSlice := strings.Split(str, " ")

	if len(tempSlice) != 3 {
		return "", 0, 0, errors.New("Неверный ввод. Ожидается имя, возраст, оценка: \"Петя 21 5\"")
	}

	tempName := tempSlice[0]

	tempAge, err := strconv.Atoi(tempSlice[1])
	if err != nil || tempAge < 16 || tempAge > 100 {
		return "", 0, 0, errors.New("Возраст указан неверно. Ожидается целое число от 16 до 100")
	}

	tempGrade, err := strconv.Atoi(tempSlice[2])
	if err != nil || tempGrade < 1 {
		return "", 0, 0, errors.New("Оценка указана неверно. Ожидается целое число не менее 1.")
	}

	return tempName, tempAge, tempGrade, err
}

// Метод хранилища, который заполняет структуру Student полученными данными и собственно создает студента.
func (s *StudentsStorage) NewStudent(name string, age int, grade int) *Student {
	newStudent := Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	return &newStudent
}

// Метод хранилища, который инициализирует это хранилище и служит для последующего вывода информации о нем.
func (s *StudentsStorage) GetAllStudents() map[string]*Student {
	return s.students
}

// Метод хранилища, который помещает структуру студента (по сути информацию о нем) в это хранилище.
func (s *StudentsStorage) PutStudent(st *Student) {
	s.students[st.Name] = st
}

// Представляет основную структуру приложения с репозиторием Storage.
// Отвечает за управление логикой приложения.
type App struct {
	repository Storage
}

// Запускает основной цикл приложения.
// В этом цикле бесконечно считывается пользовательский ввод, проводится валидация, создается новый студент
// и информация о нем помещается в хранилище.
// При вводе ctrl+d (EOF) программа завершает цикл и выводит информацию о всех студентах.
func (a *App) Run() {
	for {
		input, err := a.repository.GetInfo()
		if err == io.EOF {
			a.printInfo()
			return
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения: %v\n", err)
		} else {
			name, age, grade, err := a.repository.StringValidate(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			newStudent := a.repository.NewStudent(name, age, grade)
			a.repository.PutStudent(newStudent)
		}
	}
}

// Выводит информацию о всех студентах из хранилища.
// Если студентов нет, выводится соответствующее сообщение.
func (a *App) printInfo() {
	if len(a.repository.GetAllStudents()) > 0 {
		fmt.Println("Студенты из хранилища:")
		for _, student := range a.repository.GetAllStudents() {
			fmt.Println(student.Name, student.Age, student.Grade)
		}
	} else {
		fmt.Println("Данные о студентах отсутствуют. Выход.")
	}
}

// Создает экземпляр App с предоставленным хранилищем Storage.
func NewApp(repository Storage) *App {
	return &App{repository: repository}
}

func main() {
	repository := &StudentsStorage{students: make(map[string]*Student)}
	app := NewApp(repository)
	app.Run()
}
