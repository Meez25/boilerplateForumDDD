package main

import (
	"fmt"

	// "github.com/meez25/boilerplateForumDDD/category/infrastructure"
	// "github.com/meez25/boilerplateForumDDD/internal/category"
	"github.com/meez25/boilerplateForumDDD/internal/user"
	"github.com/meez25/boilerplateForumDDD/user/infrastructure"
)

func main() {
	// This is the main function

	userRepo := infrastructure.NewUserMemoryRepository()
	userService := user.NewUserService(userRepo)

	// Create a new user
	newUser, err := user.NewUser("meez25", "yann@memofamille.com", "password", "Yann", "Meez")

	// Save the user
	err = userService.Register(newUser)

	if err != nil {
		panic(err)
	}

	fmt.Println("User created successfully")

	// Find the user by ID
	foundUser, err := userService.FindByID(newUser.ID.String())

	if err != nil {
		panic(err)
	}

	fmt.Println("User found by ID:", foundUser)

	// ----------------------------
	// Create a category

	// categoryRepo := infrastructure.NewCategoryMemoryRepo()
	// categoryService := category.NewCategoryService(categoryRepo)
	//
	// cat, err := categoryService.Create("Category 1", "Description 1", nil)
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println("Category created successfully")
	//
	// // Find the category by ID
	//
	// foundCat, err := categoryService.FindByID(cat.ID.String())
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println("Category found by ID:", foundCat)

}
