package main

import (
	"fmt"

	"github.com/meez25/boilerplateForumDDD/infrastructure/http"
)

func main() {
	fmt.Println("Starting server on port 3000")
	http.StartServer()
}
