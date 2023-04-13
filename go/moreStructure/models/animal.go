package models

type Animal struct {
	Name string
	Age  int32
}

func ToString(animal Animal) string {
	return animal.Name + " is " + string(animal.Age)
}
