package main

import (
	"log"
)

type User struct {
    FirstName string
    LastName string
}

func mymap() {

	myMap := make(map[string]User)

    me := User {
        FirstName: "Trevor",
        LastName: "Sawler",
    }

    myMap["me"] = me

    log.Println(myMap["me"].FirstName)

    var myNewVar float32

    myNewVar = 11.1

    log.Println(myNewVar)        

}

func main() {
    numbers := []int{1,2,3,4,5,6,7,8,9,10}

    log.Println(numbers)
    log.Println(numbers[0:2])

    names := []string{"one", "seven", "fish", "cat"}

    log.Println(names)
}