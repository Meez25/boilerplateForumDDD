package main

import (
	"fmt"

	"github.com/meez25/boilerplateForumDDD/application/services"
	"github.com/meez25/boilerplateForumDDD/infrastructure/persistence"
)

func main() {
	userRepo := persistence.NewUserMemoryRepository()
	userGroupRepo := persistence.NewUserGroupMemoryRepo()
	userService := services.NewUserService(userRepo, userGroupRepo)

	// Create a new user
	newUser, err := userService.Create("John Doe", "johnDoe@doe.com", "password", "john", "doe", "profile.jpg")

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

	// ----------------------------
	// Create a topic

	topicRepo := persistence.NewTopicMemoryRepo()
	topicService := services.NewTopicService(topicRepo)

	topic, err := topicService.CreateTopic("Topic 1", "Rich content 1", newUser.ID)

	if err != nil {
		panic(err)
	}

	fmt.Println("Topic created successfully")

	// Add a message to the topic

	err = topicService.AddMessage(topic.ID.String(), "Message 1", newUser.ID)

	if err != nil {
		panic(err)
	}

	fmt.Println("Message added successfully")

	// Find the topic by ID

	foundTopic, err := topicService.GetTopicByID(topic.ID.String())

	if err != nil {
		panic(err)
	}

	fmt.Println("Topic found by ID:", foundTopic)

	// ----------------------------

	// Create a user group

	userGroup, err := userService.CreateGroup("Group 1", "Group description 1", newUser)

	if err != nil {
		panic(err)
	}

	fmt.Println("User group created successfully")

	// Find the user group by ID

	foundUserGroup, err := userService.FindGroupByID(userGroup.ID.String())

	if err != nil {
		panic(err)
	}

	fmt.Println("User group found by ID:", foundUserGroup)

}
