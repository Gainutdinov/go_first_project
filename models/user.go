package models

type User struct {
    ID        int
    FirstName string
    LastName  string
}

var (
    // slice which will hold pointers to User struct variables
    users []*User
    nextID = 1
)
