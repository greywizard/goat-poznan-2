package main

import (
	"database/sql"
	"fmt"
)

var dbConnection *sql.DB

func appendExample() {
	objects := []string{}

	for {
		data := getSomeData()
		for _, v := range data {
			newObject := parseValue(v)
			objects = append(objects, newObject)
		}

		// (..)

		for _, v := range objects {
			doSomeStuff(v)
		}
	}
}

//AllegroCase is case
func AllegroCase() {
	x := []int{0}
	x = append(x, 1)
	x = append(x, 2)

	y := append(x, 3)
	z := append(x, 4)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func getActiveUsers() ([]string, error) {
	users := []string{}
	rows, err := dbConnection.Query("SELECT name FROM users WHERE active = ?", 1)
	defer rows.Close()

	if err != nil {
		return users, err
	}

	defer rows.Close()

	//(...)

	return users, nil
}

func main() {
	AllegroCase()
}

func parseValue(v string) string {
	return ""
}

func doSomeStuff(v string) {
}

func getSomeData() []string {
	return []string{}
}
