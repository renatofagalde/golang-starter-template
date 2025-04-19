package env

import "os"

func GetNewsTokenAPI() string {
	return os.Getenv("NEW_API_KEY")
}
