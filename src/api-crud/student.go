package main

var students = []*student{}

type student struct {
	id    string
	name  string
	Grade int32
}

func GetStudents() []*student {
	return students
}

func SelectStudent(id string) *student {
	for _, each := range students {
		if each.id == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &student{id: "s01", name: "fadhilla", Grade: 2})
	students = append(students, &student{id: "s02", name: "hadiat", Grade: 2})
	students = append(students, &student{id: "s03", name: "nathania", Grade: 2})
}
