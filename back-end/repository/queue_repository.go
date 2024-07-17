package repository

import "back-end/model"

type QueueRepository interface {
	ClearQueue(key string) error
	Enqueue(key string, song model.Song) error
	Dequeue(key string) (model.Song, error)
	GetQueue(key string) (model.Song, error)
	GetAllQueue(key string) ([]model.Song, error)
	RemoveFromQueue(key string, index int) error
}
