package main

import (
	"errors"
	"fmt"
)

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := CreateOrg()
	fmt.Println(err)
}
