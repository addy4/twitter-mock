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
			followNotification := <-data.FollowNotifier
			fmt.Printf("NOTIFICATION001: User ID %s has followed %s\n", followNotification.CurrentUser, followNotification.Followee)

			for _, wsclients := range data.ActiveSessions {
				wsclients.Websocket_Session.WriteJSON(followNotification)
			}
		}

	}()

	wg.Add(1)

	go func() {

		defer wg.Done()
		for {
			postNotification := <-data.PostsNotifier
			fmt.Printf("NOTIFICATION002: User ID %s has posted %s\n", postNotification.CurrentUser, postNotification.ContentPost)

			for _, wsclients := range data.ActiveSessions {
				wsclients.Websocket_Session.WriteJSON(postNotification)
			}
		}

	}()

	wg.Add(1)

	go func() {

		defer wg.Done()
		for {

			sessionAdded := <-data.SessionNotifier
			fmt.Printf("NOTIFICATION003: New Session added with session-id %s and user-id %s\n", sessionAdded.Session_Id, sessionAdded.User_Id)

			/*
				for _, sessionDetails := range data.ActiveSessions {
					fmt.Printf("NOTIFICATION003: Existing Session added with session-id %s and user-id %s\n", sessionDetails.Session_Id, sessionDetails.User_Id)
				}
			*/

		}

	}()

	wg.Wait()
}
