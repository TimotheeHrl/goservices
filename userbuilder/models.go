package userbuilder

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/grokify/goauth/scim"
	"github.com/grokify/mogo/type/stringsutil"
)

type User struct {
	Gender      string      `json:"gender,omitempty"`
	Name        Name        `json:"name,omitempty"`
	Location    Location    `json:"location,omitempty"`
	Email       string      `json:"email,omitempty"`
	Login       Login       `json:"login,omitempty"`
	DateOfBirth DateOfBirth `json:"dob,omitempty"`
	Registered  Registered  `json:"registered,omitempty"`
	Phone       string      `json:"phone,omitempty"`
	Cell        string      `json:"cell,omitempty"`
	ID          ID          `json:"id,omitempty"`
	Picture     Picture     `json:"picture,omitempty"`
	Nationality string      `json:"nat,omitempty"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

var rxAZ = regexp.MustCompile(`^[A-Za-z]+$`)

func (name *Name) IsAZSimple() bool {
	name.First = strings.TrimSpace(name.First)
	name.Last = strings.TrimSpace(name.Last)
	if rxAZ.MatchString(name.First) && rxAZ.MatchString(name.Last) {
		return true
	}
	return false
}

func (user *User) PhoneFormatted() string {
	raw := strings.TrimSpace(user.Phone)
	if user.Nationality == "US" {
		try := stringsutil.DigitsOnly(raw)
		if len(try) == 10 {
			return "+1" + try
		}
	}
	return raw
}

func (user *User) CellFormatted() string {
	raw := strings.TrimSpace(user.Cell)
	if user.Nationality == "US" {
		try := stringsutil.DigitsOnly(raw)
		if len(try) == 10 {
			return "+1" + try
		}
	}
	return raw
}

func (user *User) Scim() scim.User {
	return UserToScim(*user)
}

type Location struct {
	Street      Street      `json:"street,omitempty"`
	City        string      `json:"city,omitempty"`
	State       string      `json:"state,omitempty"`
	Postcode    int         `json:"postcode,omitempty"`
	Coordinates Coordinates `json:"coordinates,omitempty"`
	Timezone    Timezone    `json:"timezone,omitempty"`
}

type Street struct {
	Number int    `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
}

func (street *Street) String() string {
	return fmt.Sprintf("%v %s", street.Number, street.Name)
}

type Coordinates struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

type Timezone struct {
	Offset      string `json:"offset,omitempty"`
	Description string `json:"description,omitempty"`
}

type Login struct {
	UUID     string `json:"uuid,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Salt     string `json:"salt,omitempty"`
	MD5      string `json:"md5,omitempty"`
	SHA1     string `json:"sha1,omitempty"`
	SHA256   string `json:"sha256,omitempty"`
}

type DateOfBirth struct {
	Date time.Time `json:"date,omitempty"`
	Age  int       `json:"age,omitempty"`
}

type Registered struct {
	Date time.Time `json:"date,omitempty"`
	Age  int       `json:"age,omitempty"`
}

type ID struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Picture struct {
	Large     string `json:"large,omitempty"`
	Medium    string `json:"medium,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}
