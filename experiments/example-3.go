package experiments

import "fmt"

func GetUserV2(id string) string {
	return fmt.Sprintf(
		"SELECT * FROM users WHERE id = '%s'",
		id,
	)
}
