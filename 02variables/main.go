package main

import "fmt"

const LoginToken string = "bfdjkb873478fdsbkj" // Public


func main() {
	var username string = "Nishant"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallvalue int = 256
	fmt.Println(smallvalue)
	fmt.Printf("Variable is of type: %T \n", smallvalue)

	var smallFloat float64 = 256.5534343243434
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aiases
	var defaultValue int
	fmt.Println(defaultValue)
	fmt.Printf("Variable is of type: %T \n", defaultValue)

	// implicit type
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)
	// website = 3 // not allowed

	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
