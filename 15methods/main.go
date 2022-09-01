package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent
	pranav := User{"Pranav", "pranavkonde2020@gmail.com", true, 21}
	fmt.Println(pranav)
	fmt.Printf("Pranav detail are: %+v\n", pranav)
	fmt.Printf("Name is %v and email is %v\n", pranav.Name, pranav.Email)
	pranav.GetStatus()
	pranav.NewMail()
	fmt.Printf("Name is %v and email is %v\n", pranav.Name, pranav.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is:", u.Email)
}
