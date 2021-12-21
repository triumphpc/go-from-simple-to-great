//
// Go (Golang) From simple to great. The Complete Developer's Guide.
// Lesson 7.6: Method-values and method-expressions
//
// @author Alex Versus 2021
// www.alexversus.com

package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  uint
	sex  string
}

type Letter struct {
	title string
	user  User
}

type Cat struct {
	name string
}

// Entrypoint
func main() {
	user1 := User{"Alex", 25, "male"}
	user2 := User{"Kate", 18, "female"}

	// value method from user1
	isMaleMethod := user1.isMale
	fmt.Println(isMaleMethod()) // true

	user1.sex = "female"
	fmt.Println(isMaleMethod()) // false

	addPrefixMethod := user2.addPrefix
	fmt.Println(addPrefixMethod("Prefix ")) // &{PrefixKate 18 female}
	fmt.Println(user2)                      // {PrefixKate 18 female}

	// example call method with timeout
	l := new(Letter)
	l.title = "New letter"
	l.user = user1

	// set timeout
	time.AfterFunc(5*time.Second, l.sendToUser)

	// method-expression example
	u := &User{}
	u.name = "Igor"
	u.sex = "male"

	//equal calls
	fmt.Println(u.NameAndSex("Hi "))
	// expression
	exp := User.NameAndSex
	fmt.Println(exp(*u, "Hi "))
	// same
	fmt.Println((User).NameAndSex(*u, "Hi "))

	// by pointer
	fmt.Println(u.NameAndSexSecond("Hi "))
	// expression
	exp2 := (*User).NameAndSexSecond
	fmt.Println(exp2(u, "Hi "))
	fmt.Printf("%T", exp2)
	// same
	fmt.Println((*User).NameAndSexSecond(u, "Hi "))

	// Example use method-expressions
	cat := Cat{"Masha"}
	var act func(*Cat, int)

	if true {
		act = (*Cat).SayMiau
	} else {
		act = (*Cat).Sleep
	}

	// run
	action(act, &cat, 3)

}

// Check on male
func (u *User) isMale() bool {
	if u.sex == "male" {
		return true
	}
	return false
}

func (u *User) addPrefix(p string) *User {
	u.name = p + u.name
	return u
}

// Send letter to user
func (l *Letter) sendToUser() {
	fmt.Println("Send letter ", l.title, "to user ", l.user.name)
}

// Method-expression example
func (u User) NameAndSex(p string) string {
	return p + u.sex + " " + u.name
}

// Method-expression example by point
func (u *User) NameAndSexSecond(p string) string {
	return p + u.sex + " " + u.name + " second"
}

// one method
func (c *Cat) SayMiau(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Println(c.name, "say miau")
	}
}

// second method
func (c *Cat) Sleep(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Println(c.name, "sleep")
	}

}

// Action func
func action(f func(*Cat, int), c *Cat, cnt int) {
	f(c, cnt)
}
