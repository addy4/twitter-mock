package registeration

import "projects.com/apps/twitter-app/data"

func RegisterUser() {
	data.RegisteredUsers["aabhatia"] = data.SignedUser{User_Name: "aabhatia", User_Id: "b39503c6-6952-49e5-b0ea-efdbbcd112ef", Email_Id: "ab.com"}
	data.RegisteredUsers["user1"] = data.SignedUser{User_Name: "user1", User_Id: "727c360f-8674-4474-9da8-ec6bb1845bf3", Email_Id: "ab.com"}
	data.RegisteredUsers["user2"] = data.SignedUser{User_Name: "user2", User_Id: "59a9cf45-dcc7-4c43-88ee-d4a2d4292872", Email_Id: "ab.com"}
}
