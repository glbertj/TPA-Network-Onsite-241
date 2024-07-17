package sse

import (
	"back-end/model"
	"back-end/repository"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type NotificationSSE struct {
	UserRepo            repository.UserRepository
	NotificationChannel map[string]chan model.Notification
}

func NewNotificationSSE(userRepo repository.UserRepository) *NotificationSSE {
	return &NotificationSSE{
		UserRepo:            userRepo,
		NotificationChannel: make(map[string]chan model.Notification),
	}
}

//func (n NotificationSSE) StreamNotification(ctx *gin.Context) {
//	ctx.Writer.Header().Set("Cache-Control", "no-cache")
//	ctx.Writer.Header().Set("Connection", "keep-alive")
//	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
//	ctx.Writer.Header().Set("Transfer-Encoding", "chunked")
//
//	token, err := ctx.Cookie("jwt")
//	if err != nil {
//		return
//	}
//	user, err := n.userService.GetCurrentUser(token)
//	if err != nil {
//		return
//	}
//
//	if _, exists := n.NotificationChannel[user.UserId]; !exists {
//		n.NotificationChannel[user.UserId] = make(chan model.Notification)
//	}
//	notify := n.NotificationChannel[user.UserId]
//	ctx.Stream(func(w io.Writer) bool {
//		select {
//		case notification := <-notify:
//			notificationData, err := json.Marshal(notification)
//			if err != nil {
//				fmt.Println("Error marshalling notification:", err)
//				return true
//			}
//
//			event := fmt.Sprintf("event: notif-updated\ndata: %s\n\n", notificationData)
//			_, err = fmt.Fprint(w, event)
//			if err != nil {
//				fmt.Println("Error writing event:", err)
//				return false
//			}
//		case <-ctx.Request.Context().Done():
//			return false
//		}
//		return true
//	})
//
//}

func (n NotificationSSE) StreamNotification(ctx *gin.Context) {
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	id := ctx.Query("id")

	fmt.Println("ID: ", id)

	if _, exists := n.NotificationChannel[id]; !exists {
		n.NotificationChannel[id] = make(chan model.Notification)
	}

	ctx.Stream(func(w io.Writer) bool {
		initialEvent := "event: initial\ndata: Welcome\n\n"
		_, _ = fmt.Fprint(w, initialEvent)
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		for notification := range n.NotificationChannel[id] {
			fmt.Println("hai")
			notificationData, err := json.Marshal(notification)
			if err != nil {
				fmt.Println("Error marshalling notification:", err)
				continue
			}

			event := fmt.Sprintf("event: notif-updated\ndata: %s\n\n", notificationData)
			_, _ = fmt.Fprint(w, event)
			if f, ok := w.(http.Flusher); ok {
				fmt.Println("Flush")
				f.Flush()
			} else {
				fmt.Println("Not Flush")
			}

			if ctx.Writer.Status() != http.StatusOK {
				fmt.Println("Status not OK")
				return false
			}
		}
		return true
	})

}
