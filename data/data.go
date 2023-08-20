package data

type FriendsMemory map[int]map[int]bool

var (
	Friends = make(FriendsMemory)
)
