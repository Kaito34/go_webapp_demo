package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
    changeUsingPointer(&myString)
    log.Println("afer func call myString is set to", myString)

}


// pointer string
func changeUsingPointer(s *string) {
    log.Println("s is set to", s)
    newValue := "Red"
    *s = newValue
}