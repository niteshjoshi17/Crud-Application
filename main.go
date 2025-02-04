package main

import (
	"fmt"
)

func main() {
	CreateUser()
	ReadUsers()
	updateUser(&users)
	DeleteUser()
	fmt.Println("Welcome to Crud Application!")
}

type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt string
}

// Use Slice to store multiple values:
var users []User

func CreateUser() {
	var user User // Declare a new User struct

	// Take input from the user
	fmt.Print("Enter User ID: ")
	fmt.Scan(&user.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&user.Name)

	fmt.Print("Enter Email: ")
	fmt.Scan(&user.Email)

	fmt.Print("Enter Created At (YYYY-MM-DD): ")
	fmt.Scan(&user.CreatedAt)

	// Append the user to Users slice
	users = append(users, user)

	fmt.Println("User created successfully!")
	fmt.Println("User List:", users)
}

func ReadUsers() {
	if len(users) == 0 {
		fmt.Println("No users found")
		return // Exit if no users exist
	}

	fmt.Println("\n List of Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Created At: %s\n",
			user.ID, user.Name, user.Email, user.CreatedAt)
	}
}

func updateUser(users *[]User) {
	var id int
	fmt.Print("Enter User ID to Update: ")
	fmt.Scan(&id)

	// Find user by ID and update details
	for i := range *users {
		if (*users)[i].ID == id {
			fmt.Print("Enter New Name: ")
			fmt.Scan(&(*users)[i].Name)
			fmt.Print("Enter New Email: ")
			fmt.Scan(&(*users)[i].Email)

			fmt.Println("User Updated Successfully!")
			return
		}
	}
	fmt.Println("User Not Found.")
}

func DeleteUser() {
	var id int
	fmt.Print("Enter User ID to Delete: ")
	fmt.Scan(&id)

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			fmt.Println("User Deleted Successfully!")
			return
		}
	}

	fmt.Println("User Not Found.")
}
