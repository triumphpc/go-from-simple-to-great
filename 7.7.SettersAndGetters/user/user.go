package user

import "errors"

type User struct {
	name       string
	secondName string
	age        int
}

// Set name for user struct
func (u *User) SetName(n string) error {
	if len(n) < 2 {
		return errors.New("invalid name")
	}
	u.name = n
	return nil
}

// Set Second name for user struct
func (u *User) SetSecondName(n string) error {
	if len(n) < 2 {
		return errors.New("invalid name")
	}
	u.secondName = n
	return nil
}

// Set age for user struct
func (u *User) SetAge(a int) error {
	if a < 1 || a > 110 {
		return errors.New("invalid age")
	}
	u.age = a
	return nil
}

// Get methods
func (u *User) Name() string {
	return u.name
}

func (u *User) SecondName() string {
	return u.secondName
}

func (u *User) Age() int {
	return u.age
}
