package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Structs
type User struct {
	FirstName string
	LastName  string
	Email     string
}

type Car struct {
	Make  string
	Model string
	Year  int
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func (c Car) Info() string {
	return c.Make + " " + c.Model + " " + string(c.Year)
}
func (c Car) Drive() string {
	return "Driving"
}

func (c Car) Stop() string {
	return "Stopping"
}

func main() {
	fmt.Println("Hello, World!")

	//http server
	//	http.ListenAndServe(":3333", nil)

	// Structs
	u := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@gmail.com"}

	fmt.Println(u.FullName())

	c := Car{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2019}

	fmt.Println(c.Info())
	fmt.Println(c.Drive())
	fmt.Println(c.Stop())

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(u)
	})
	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(c)
	})

	http.ListenAndServe(":3333", nil)

}
