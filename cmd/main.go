package main

import (
	"fmt"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/persistence"
)

func main() {
	userRepo := persistence.NewUserMemoryRepository()
	userService := services.NewUserService(userRepo)

	// Create a new user
	newUser, err := userService.Create("John Doe", "johnDoe@doe.com", "password", "john", "doe")

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

	categoryRepo := persistence.NewCategoryMemoryRepo()
	categoryService := services.NewCategoryService(categoryRepo)

	cat, err := categoryService.Create("Category 1", "Description 1", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Category created successfully")

	// Find the category by ID

	foundCat, err := categoryService.FindByID(cat.ID.String())

	if err != nil {
		panic(err)
	}

	fmt.Println("Category found by ID:", foundCat)

	// Add a subcategory

	subCat, err := categoryService.Create("Subcategory 1", "Subcategory description 1", nil)

	if err != nil {
		panic(err)
	}

	err = categoryService.AddSubCategory(cat.ID.String(), subCat.ID.String())

	if err != nil {
		panic(err)
	}

	fmt.Println("Subcategory added successfully")

	// Find the category by ID

	foundCat, err = categoryService.FindByID(subCat.ID.String())

	if err != nil {
		panic(err)
	}

	fmt.Println("Category found by ID:", foundCat)

}
