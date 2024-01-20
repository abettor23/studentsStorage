package student

// Student представляет данные о студенте.
type Student struct {
	Name  string
	Age   int
	Grade int
}

// NewStudent создает и возвращает новый экземпляр Student.
func NewStudent(name string, age int, grade int) *Student {
	newStudent := Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	return &newStudent
}
