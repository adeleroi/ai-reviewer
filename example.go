package main

import "fmt"

func GetUser(id string) string {
	return fmt.Sprintf(
		"SELECT * FROM users WHERE id = '%s'",
		id,
	)
}
