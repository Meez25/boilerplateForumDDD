package main

import (
	"flag"
	"fmt"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/http"
	"github.com/meez25/boilerplateForumDDD/infrastructure/persistence"
)

func main() {
	command := flag.String("command", "startserver", "command to control the application")

	db := persistence.NewSQLConnection()
	defer db.Close()
	userRepository := persistence.NewUserSQLRepository(db)
	userGroupRepository := persistence.NewUserGroupMemoryRepo()
	userService := services.NewUserService(userRepository, userGroupRepository)

	if *command == "startserver" {
		fmt.Println("Starting server on port 3000")
		http.StartServer()
	}

	var (
		username        string
		email           string
		password        string
		confirmPassword string
		firstName       string
		lastName        string
	)

	flag.StringVar(&username, "name", "admin", "Your username")
	flag.StringVar(&email, "email", "admin@admin.com", "Your email")
	flag.StringVar(&password, "password", "", "Your password")
	flag.StringVar(&confirmPassword, "confirmPassword", "", "Your password")
	flag.StringVar(&firstName, "firstName", "first name", "Your first name")
	flag.StringVar(&lastName, "lastName", "Last name", "Your last name")

	flag.Parse()

	userService.CreateAdmin(username, email, password, confirmPassword, firstName, lastName, true)
}
