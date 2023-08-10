package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	createOption = "a"
	readOption   = "b"
	updateOption = "c"
	deleteOption = "d"
)

func readLine() string {
	var reader bufio.Reader
	reader = *bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')

	if err != nil {
		panic("It's not possible to get the value!")
	} else {
		return strings.TrimSpace(line)
	}
}
func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type User struct {
	id    int
	name  string
	email string
	age   int
}

var id int = 0
var users map[int]User

func createUser() {
	clearConsole()
	fmt.Print("Enter a name: ")
	name := readLine()
	fmt.Print("Enter an email: ")
	email := readLine()
	fmt.Print("Enter an age: ")
	age, err := strconv.Atoi(readLine())
	if err != nil {
		panic("It is not possible convert that value to int")
	}
	id++

	user := User{id, name, email, age}
	users[id] = user

	fmt.Println(">>> User created! ")
}
func readUser() {
	clearConsole()
	for id, user := range users {
		fmt.Printf("%d - %s \n", id, user.name)
	}
	fmt.Println("")
}
func updateUser() {
	clearConsole()
	fmt.Print("Enter user id to update: ")
	id, err := strconv.Atoi(readLine())
	if err != nil {
		panic("It is not possible convert that value to int")
	}
	if _, ok := users[id]; !ok {
		panic("User not found !")
	}

	var user User

	fmt.Print("Enter a name: ")
	name := readLine()
	user.name = name
	fmt.Print("Enter an email: ")
	email := readLine()
	user.email = email
	fmt.Print("Enter an age: ")
	age, err := strconv.Atoi(readLine())
	if err != nil {
		panic("It is not possible convert that value to int")
	}
	user.age = age
	users[id] = user

	fmt.Println(">>> updated")
}
func deleteUser() {
	clearConsole()
	fmt.Print("Enter user id to delete: ")
	id, err := strconv.Atoi(readLine())
	if err != nil {
		panic("It is not possible convert that value to int")
	}
	if _, ok := users[id]; ok {
		delete(users, id)
	}
	fmt.Println(">>> deleted")
}
func main() {
	users = make(map[int]User)
	for {
		fmt.Println("___________")
		fmt.Println("a) create")
		fmt.Println("b) read")
		fmt.Println("c) update")
		fmt.Println("d) delete")
		fmt.Println("")
		fmt.Println("q) quit")
		fmt.Println("___________")
		fmt.Print("Select a option: ")

		option := readLine()

		if option == "q" {
			break
		}

		switch option {
		case createOption:
			createUser()
		case readOption:
			readUser()
		case updateOption:
			updateUser()
		case deleteOption:
			deleteUser()
		default:
			clearConsole()
			fmt.Println("noup")
		}
	}
	fmt.Println("Bye Bye.")
}
