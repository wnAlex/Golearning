package student

type student struct {
	Name string
	Age  int
	sex  string
}

func GetStudent(name string, age int) student {
	return student{
		Name: name,
		Age:  age,
	}
}
func GetStudentPointer(name string, age int) *student {
	return &student{
		Name: name,
		Age:  age,
	}
}
func (s *student) GetSex() string {
	return s.sex
}
func (s *student) SetSex(sex string) {
	s.sex = sex
}
