package services

import (
	"back-end/data/response"
	"back-end/model"
	"back-end/repository"
	"github.com/go-playground/validator/v10"
)

type QueueServiceImpl struct {
	QueueRepository repository.QueueRepository
	Validate        *validator.Validate
}

func NewQueueServiceImpl(QueueRepository repository.QueueRepository, Validate *validator.Validate) *QueueServiceImpl {
	return &QueueServiceImpl{QueueRepository: QueueRepository, Validate: Validate}

}

func (q QueueServiceImpl) ClearQueue(key string) error {
	err := q.QueueRepository.ClearQueue(key)
	if err != nil {
		return err
	}
	return nil
}

func (q QueueServiceImpl) Enqueue(key string, song model.Song) error {
	err := q.QueueRepository.Enqueue(key, song)
	if err != nil {
		return err
	}
	return nil
}

func (q QueueServiceImpl) Dequeue(key string) (response.SongResponse, error) {
	song, err := q.QueueRepository.Dequeue(key)
	if err != nil {
		return response.SongResponse{}, err
	}
	res := response.SongResponse{
		SongId:      song.SongId,
		Title:       song.Title,
		ArtistId:    song.ArtistId,
		AlbumId:     song.AlbumId,
		ReleaseDate: song.ReleaseDate,
		Duration:    song.Duration,
		File:        song.File,
		Album:       song.Album,
		Play:        song.Play,
		Artist:      song.Artist,
	}

	return res, nil
}

func (q QueueServiceImpl) GetQueue(key string) (response.SongResponse, error) {
	song, err := q.QueueRepository.GetQueue(key)
	if err != nil {
		return response.SongResponse{}, err
	}
	res := response.SongResponse{
		SongId:      song.SongId,
		Title:       song.Title,
		ArtistId:    song.ArtistId,
		AlbumId:     song.AlbumId,
		ReleaseDate: song.ReleaseDate,
		Duration:    song.Duration,
		File:        song.File,
		Album:       song.Album,
		Play:        song.Play,
		Artist:      song.Artist,
	}

	return res, nil
}

func (q QueueServiceImpl) GetAllQueue(key string) ([]response.SongResponse, error) {
	songs, err := q.QueueRepository.GetAllQueue(key)
	if err != nil {
		return nil, err
	}
	var res []response.SongResponse
	for _, song := range songs {
		res = append(res, response.SongResponse{
			SongId:      song.SongId,
			Title:       song.Title,
			ArtistId:    song.ArtistId,
			AlbumId:     song.AlbumId,
			ReleaseDate: song.ReleaseDate,
			Duration:    song.Duration,
			File:        song.File,
			Album:       song.Album,
			Play:        song.Play,
			Artist:      song.Artist,
		})
	}

	return res, nil
}

func (q QueueServiceImpl) RemoveFromQueue(key string, index int) error {
	err := q.QueueRepository.RemoveFromQueue(key, index)
	if err != nil {
		return err
	}
	return nil
}
