package model

// Image struct for 'images' in 'Pictures' database
type Image struct {
	Imageid     int 	`json:"imageid"`
	Imageurl    string	`json:"imageurl"`
	Userid  	string	`json:"userid"`
	Private 	bool	`json:"private"`
}

// User struct for 'users' in 'Pictures' database
type User struct {
	Userid        string `json:"userid"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// UserWithImages struct for holding Users and Images
type UserWithImages struct {
	User User `json:"user"`
	Images []Image `json:"images"`
}
