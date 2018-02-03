package controller

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/xizahk/gowebsite/app/model"
	"github.com/xizahk/gowebsite/app/database"
)

var usersWithImages []model.UserWithImages

// GetUsersWithImagesHandler writes as response a JSON of an array of UserWithImages
//   queried from pictures database.
func GetUsersWithImagesHandler(writer http.ResponseWriter, route *http.Request) {
	// Get users with images from database
	usersWithImages, err := database.GetUsersWithImages()
	// Convert usersWithImages variable to JSON
	usersWithImagesList, err := json.Marshal(usersWithImages)

	// Check for error
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write JSON of usersWithImages as response
	writer.Write(usersWithImagesList)
}