package main

import (
	"fmt"
)

//1. User defined types

type User struct {
	name string
	email string
	ext        int
	privileged bool
}

// Admin Declaring fields based on other struct types
// Admin represents an admin user with privileges
type Admin struct {
	User
	level string
}

//METHODS
// Notify implements a method with a value receiver
func (u User) notify()  {
	fmt.Printf("Sending User Email to:  %s<%s>\n", u.name,u.email)
}
//Change email
func (u *User) changeEmail(email string)  {
	u.email = email
}

//REFERENCE TYPES
/*
type IP []byte

func (ip IP) MarshalText() ([]byte, error) {
	if len(ip) == 0{
		return []byte(""), nil
	}
	if len(ip) != net.IPv4len && len(ip) != net.IPv6len{
		return nil, errors.New("invalid IP address")
	}
	return []byte(ip.String()), nil
}



func ipEmptyString(ip IP) string {
	if len(ip) ==0{
		return ""
	}

	return string(ip.String())
}
*/








func main()  {
	// DECLARATION OF A VARIABLE OF THE STRUCT TYPE USING A STRUCT LITERAL
	//lisa := User{"Lisa",
	//	"lisa@email.com",
	//	123,
	//	true}

	bill := User{"Bill",
		"bill@gmail.com",
		123,true}
	bill.notify()
	bill.changeEmail("billy@gmail.com")
	bill.notify()
	//FOR ADMIN
	//fred := Admin{User{"Fred",
	//	"fred@gmail.com",
	//	123,
	//	true},
	//	"super"}
}