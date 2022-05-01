package main

import "fmt"

func Run() error {
	return nil
}

func main() {
	if err := Run(); err != nil {

		fmt.Println("Error Setting Up our application")

		fmt.Println(err)
	}
}
