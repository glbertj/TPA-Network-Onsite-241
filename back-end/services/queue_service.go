package services

import (
	"back-end/data/response"
	"back-end/model"
)

type QueueService interface {
	ClearQueue(key string) error
	Enqueue(key string, song model.Song) error
	Dequeue(key string) (response.SongResponse, error)
	GetQueue(key string) (response.SongResponse, error)
	GetAllQueue(key string) ([]response.SongResponse, error)
	RemoveFromQueue(key string, index int) error
}
