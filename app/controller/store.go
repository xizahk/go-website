package controller

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/xizahk/gowebsite/app/model"
)

func getUsers(db *sql.DB) ([]model.User, error) {
	// Query all users from 'users' table.
	rows, err := db.Query("SELECT * FROM users")
	// Check for error
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a list of Users by iterating through rows from 'users' table
	var users = []model.User{}
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Userid, &user.Username, &user.Firstname, &user.Lastname)
		// Check for error
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func getImageMap(db *sql.DB) (map[string][]model.Image, error) {
	// Query all images from 'images' table.
	rows, err := db.Query("SELECT * FROM images")
	// Check for error
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a map of userid (string) to a list of images by iterating through rows from 'images' table
	var imageMap = make(map[string][]model.Image)
	// Iterate through each row and append a new Image object to the corresponding list of images using 
	//   the image's Userid as key
	for rows.Next() {
		var image model.Image
		var privateString string
		err = rows.Scan(&image.Imageid, &image.Imageurl, &image.Userid, &privateString)
		if err != nil {
			return nil, err
		}
		// Convert bit type from MySQL table to bool
		image.Private = privateString == "\x01"
		imageMap[image.Userid] = append(imageMap[image.Userid], image)
	}

	return imageMap, nil
}

// GetUsersWithImages returns list of UserWithImages from database
func GetUsersWithImages() ([]*model.UserWithImages, error) {
	// Open connection to MySQL database
	db, err := sql.Open("mysql", "root:secretpassword@tcp(54.175.23.88)/pictures")
	// Check for errors
	if err != nil {
		return nil, err
	}
	// Ping database
	err = db.Ping()
	// Check for errors
	if err != nil {
		return nil, err
	}
	defer db.Close()
	
	// Get a list of Users from database
	users, err := getUsers(db)
	if err != nil {
		return nil, err
	}

	// Get a userMap of (key: userID (string), value: model.Image) from database
	imageMap, err := getImageMap(db)
	if err != nil {
		return nil, err
	}
	
	// Create a list of UserWithImages by iterating through the list of Users and the imageMap
	usersWithImages := []*model.UserWithImages{}
	for _, user := range users {
		// Loop through each user and create a new UserWithImages object for each user
		userWithImages := &model.UserWithImages{}
		userWithImages.User = user
		userWithImages.Images = []model.Image{}
		for _, image := range imageMap[user.Userid] {
			// Append images that have the same userid as the current user.
			//  Ignore images that are set to private
			if image.Private == false {
				userWithImages.Images = append(userWithImages.Images, image)
			}
		}
		// Append the UserWithImages object to the list
		usersWithImages = append(usersWithImages, userWithImages)
	}

	return usersWithImages, nil
}