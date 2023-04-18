package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta struct {
		Visits int
	}
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "Richard Liu",
		Age:  24,
		Meta: UserMeta{
			Visits: 4,
		},
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
