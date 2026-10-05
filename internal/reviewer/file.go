package reviewer

import "os"

func ReadFile(name string) ([]byte, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return b, nil
}
