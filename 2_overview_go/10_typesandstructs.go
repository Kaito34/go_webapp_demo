package main

import (
	"log"
	"time"
)

// declare and initialze at the same time
var s = "seven"

// useful way to group data together
type User struct {
    // why capital? capital can be allowed from outside the package, small letter can be only accessed within the package
    FirstName string
    LastName string
    PhoneNumber string
    Age int
    BirthDate time.Time
}

var Special string


func main() {
    var s2 = "six"


    log.Println("s is", s)
    log.Println("s2 is", s2)

    saySomething("xxx")

    user := User {
        FirstName: "Trevour",
        LastName: "Savler",
        PhoneNumber: "1 555 555-1212",
    }

    log.Println(user.FirstName, user.LastName, user.BirthDate)
}

// hand back string
func saySomething(s3 string) (string, string) {
    log.Println("s from the saySomething func is", s)
    return s3, "World"
}