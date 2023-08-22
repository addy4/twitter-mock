package data

import (
	"fmt"
)

type FriendsMemory map[int]map[int]bool

var (
	Friends = make(FriendsMemory)
)

func GetFollowers(friendList FriendsMemory, userID int) {

	fmt.Println("For user ID: ", userID)
	for follower := range friendList[userID] {
		fmt.Printf("... above user follows %d with status\n", follower)
	}
}
