package repository

import (
	"back-end/database"
	"back-end/model"
	"encoding/json"
	"errors"
)

type QueueRepositoryImpl struct {
	rdb *database.Redis
}

func NewQueueRepositoryImpl(rdb *database.Redis) *QueueRepositoryImpl {
	return &QueueRepositoryImpl{rdb: rdb}
}

func (q *QueueRepositoryImpl) ClearQueue(key string) error {
	return q.rdb.Del(key)
}

func (q *QueueRepositoryImpl) Enqueue(key string, song model.Song) error {
	songBytes, err := json.Marshal(song)
	if err != nil {
		return err
	}
	return q.rdb.RPush(key, songBytes)
}

func (q *QueueRepositoryImpl) Dequeue(key string) (model.Song, error) {
	songBytes, err := q.rdb.LPop(key)
	if err != nil {
		return model.Song{}, err
	}
	if songBytes == nil {
		return model.Song{}, errors.New("queue is empty")
	}

	var song model.Song
	if err := json.Unmarshal(songBytes, &song); err != nil {
		return model.Song{}, err
	}

	return song, nil
}

func (q *QueueRepositoryImpl) GetQueue(key string) (model.Song, error) {
	songBytes, err := q.rdb.LIndex(key, 0)
	if err != nil {
		return model.Song{}, err
	}
	if songBytes == nil {
		return model.Song{}, errors.New("queue is empty")
	}

	var song model.Song
	if err := json.Unmarshal(songBytes, &song); err != nil {
		return model.Song{}, err
	}

	return song, nil
}

func (q *QueueRepositoryImpl) GetAllQueue(key string) ([]model.Song, error) {
	songBytes, err := q.rdb.LRange(key, 0, -1)
	if err != nil {
		return nil, err
	}

	var songs []model.Song
	for _, songByte := range songBytes {
		var song model.Song
		if err := json.Unmarshal(songByte, &song); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (q *QueueRepositoryImpl) RemoveFromQueue(key string, index int) error {
	placeholder := "TO_BE_DELETED"

	err := q.rdb.LSet(key, int64(index), placeholder)
	if err != nil {
		return err
	}

	err = q.rdb.LRem(key, 1, placeholder)
	if err != nil {
		return err
	}

	return nil
}
