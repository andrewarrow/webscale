package pixabay

import (
	"os"
)

var apiKey string

func Init() {
	apiKey = os.Getenv("PIXABAY")
}
