package apis

import (
	"fmt"
	"sync"

	"projects.com/apps/twitter-app/data"
)

func Broadcast() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {

		defer wg.Done()
		for {
			followNotification := <-FollowNotifier
			fmt.Printf("NOTIFICATION001: User ID %d has followed %d\n", followNotification.CurrentUserId, followNotification.Followee)

			notification := data.FollowNotification{Action: "FollowFeed", Follower: followNotification.CurrentUserId, Followee: followNotification.Followee}

			for _, wsclients := range data.ActiveSessions {
				wsclients.Websocket_Session.WriteJSON(notification)
			}
		}

	}()

	wg.Add(1)

	go func() {

		defer wg.Done()
		for {
			postNotification := <-data.PostsNotifier
			fmt.Printf("NOTIFICATION002: User ID %s has posted %s\n", postNotification.User_Id, postNotification.Content)

			notification := data.PostedNotification{Action: "PostFeed", FolloweeUserID: postNotification.User_Id, ContentPost: postNotification.Content}

			for _, wsclients := range data.ActiveSessions {
				wsclients.Websocket_Session.WriteJSON(notification)
			}
		}

	}()

	wg.Add(1)

	go func() {

		defer wg.Done()
		for {

			sessionAdded := <-data.SessionNotifier
			fmt.Printf("NOTIFICATION003: New Session added with session-id %s and user-id %s\n", sessionAdded.Session_Id, sessionAdded.User_Id)

			for _, sessionDetails := range data.ActiveSessions {
				fmt.Printf("NOTIFICATION003: Existing Session added with session-id %s and user-id %s\n", sessionDetails.Session_Id, sessionDetails.User_Id)
			}

		}

	}()

	wg.Wait()
}
