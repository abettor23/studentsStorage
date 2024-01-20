package app

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"studentsApp/pkg/storage"
	"studentsApp/pkg/student"
)

// App управляет основными операциями приложения и взаимодействует с хранилищем студентов.
type App struct {
	repository storage.Storage
}

// Создает экземпляр App с предоставленным хранилищем Storage.
func NewApp(repository storage.Storage) *App {
	return &App{repository: repository}
}

// Запускает основной цикл приложения.
// В этом цикле бесконечно считывается пользовательский ввод, проводится валидация, создается новый студент
// и информация о нем помещается в хранилище.
// При вводе ctrl+d (EOF) программа завершает цикл и выводит информацию о всех студентах.
func (a *App) Run() {
	for {
		input, err := a.GetInfo()
		if err == io.EOF {
			a.PrintInfo()
			return
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения: %v\n", err)
		} else {
			name, age, grade, err := a.StringValidate(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			newStudent := student.NewStudent(name, age, grade)
			a.repository.Put(newStudent)
		}
	}
}

// GetInfo читает ввод пользователя строчно и возвращает введенную строку или ошибку.
func (a *App) GetInfo() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите текст: ")
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err
}

// StringValidate проверяет корректность ввода посредством образования слайса tempSlice.
// Если ввод не соответствует ожиданиям, выводится ошибка с подсказкой.
// Если ввод проходит проверку, метод возвращает данные о студенте (имя, возраст, оценку)
func (a *App) StringValidate(str string) (string, int, int, error) {
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

// printInfo выводит собранную информацию о студентах из хранилища.
func (a *App) PrintInfo() {
	info := a.repository.Get()
	fmt.Println(info)
}
