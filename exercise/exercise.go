package main

import "fmt"

type Product struct {
	id    string
	price int
}

func main() {
	hobbiesArray := [3]string{"Programming", "Reading", "Watching cat videos"}

	fmt.Println("Your hobbies are", hobbiesArray)

	firstHobby := hobbiesArray[0]
	hobbiesSlice := hobbiesArray[1:]

	fmt.Println("The first hobby is", firstHobby)
	fmt.Println("The second and third hobbies are", hobbiesSlice)

	hobbiesSlice = hobbiesArray[:2]
	fmt.Println("The first and second hobbies are", hobbiesSlice)

	hobbiesSlice = hobbiesSlice[1:3]
	fmt.Println("The second and last hobbies are", hobbiesSlice)

	goalsDynamicArray := []string{}

	goalsDynamicArray = append(goalsDynamicArray, "Master the basics of Go")
	goalsDynamicArray = append(goalsDynamicArray, "Create a dynamic linker in Go by the end of the course")
	fmt.Println("Your goals are", goalsDynamicArray)

	goalsDynamicArray[1] = "Port a project from C to Go"
	goalsDynamicArray = append(goalsDynamicArray, "Create a backend in Go")

	fmt.Println("Your updated goals are", goalsDynamicArray)

	products := []Product{{"laptop", 1300}, {"phone", 1000}}
	fmt.Println("The products are", products)

	products = append(products, Product{"pc", 2000})
	fmt.Println("The updated products are", products)
}
