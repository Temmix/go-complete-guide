package main

import (
	"fmt"

	"example.com/structx/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	var userData *user.User
	userData, err := user.New(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	// callling the receiver function
	userData.GetUserDetails()
	userData.ClearUserDetails()
	userData.GetUserDetails()

	fmt.Println("*********************************************")
	adminUser, err := user.NewAdmin("Admin@email.com", "Admin**Password")
	if err != nil {
		fmt.Println(err)
		return
	}
	adminUser.GetUserDetails()
	adminUser.ClearUserDetails()
	adminUser.GetUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
