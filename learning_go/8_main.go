package main // convention

import "fmt"

// need at least one func, also need main neccessarily
func main() { // no argument, no return
    fmt.Println("Hello, world.") // short for format
    // print line

    // declare variable. go is strict
    var whatToSay string 
    var i int
    // can't have unused variable inside function

    whatToSay = "Goodbye, cruel world"

    fmt.Println(whatToSay)
    
    i = 7

    fmt.Println("i is set to", i)

    // declare and put into variable
    whatWasSaid, theOtherThingThatWasSaid := saySomething()

    fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

// hand back string
func saySomething() (string, string) {
    return "something", "else"
}