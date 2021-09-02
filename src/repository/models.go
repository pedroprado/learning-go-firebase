package repository

type User struct {
	ID          string
	Name        string
	DateOfBirth string
	Version     int
}

type Address struct {
	Street string
	Number string
	City   string
	State  string
}

type Job struct {
	Title       string
	Description string
}
