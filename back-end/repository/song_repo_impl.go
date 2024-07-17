package repository

import (
	"back-end/data/response"
	"back-end/database"
	"back-end/model"
	"back-end/utils"
	"encoding/json"
	"gorm.io/gorm"
)

type SongRepositoryImpl struct {
	DB  *gorm.DB
	rdb *database.Redis
}

func NewSongRepositoryImpl(DB *gorm.DB, rdb *database.Redis) *SongRepositoryImpl {
	return &SongRepositoryImpl{DB: DB, rdb: rdb}
}

func (s SongRepositoryImpl) GetAllSong() (res []model.Song, err error) {
	songs, err := s.rdb.Get(utils.SongKey)
	if err != nil {
		err = s.DB.Preload("Artist").Preload("Artist.User").Preload("Album").Find(&res).Error
		if err != nil {
			return res, err
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			return res, err
		}
		_ = s.rdb.Set(utils.SongKey, string(resJSON))

		return res, nil
	} else {
		if err := json.Unmarshal([]byte(songs), &res); err != nil {
			return res, err
		}
		return res, nil
	}
}

func (s SongRepositoryImpl) GetSongById(id string) (res model.Song, err error) {
	songs, err := s.rdb.Get(utils.SongKey + id)
	if err != nil {
		err = s.DB.Where("song_id = ?", id).Preload("Artist").Preload("Artist.User").Preload("Play").Preload("Album").Find(&res).Error
		if err != nil {
			return res, err
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			return res, err
		}
		_ = s.rdb.Set(utils.SongKey+id, string(resJSON))

		return res, nil
	} else {
		if err := json.Unmarshal([]byte(songs), &res); err != nil {
			return res, err
		}
		return res, nil
	}

}

func (s SongRepositoryImpl) FindSongByTitle(title string) (res []response.SongSearch, err error) {
	err = s.DB.Table("songs AS s").
		Select("s.song_id,s.title, COUNT(p.play_id) AS play_count").
		Joins("LEFT JOIN plays p ON s.song_id = p.song_id").
		Where("UPPER(s.title) LIKE ?", "%"+title+"%").
		Group("s.song_id,s.title").
		Order("play_count DESC").
		Limit(6).
		Scan(&res).Error
	return
}

func (s SongRepositoryImpl) GetTop5TrackFromAlbum(albumId string) (res []model.Song, err error) {
	err = s.DB.Table("songs AS s").
		Select("s.*").
		Joins("LEFT JOIN plays p ON s.song_id = p.song_id").
		Where("s.album_id = ?", albumId).
		Group("s.song_id,s.title,s.release_date,s.duration,s.file,s.album_id,s.artist_id").
		Order("COUNT(p.play_id) DESC").
		Limit(5).
		Preload("Artist").
		Preload("Artist.User").
		Preload("Album").
		Find(&res).Error
	return
}

func (s SongRepositoryImpl) GetTop5TrackFromArtist(artistId string) (res []model.Song, err error) {
	err = s.DB.Table("songs AS s").
		Select("s.*").
		Joins("LEFT JOIN plays p ON s.song_id = p.song_id").
		Where("s.artist_id = ?", artistId).
		Group("s.song_id,s.title,s.release_date,s.duration,s.file,s.album_id,s.artist_id").
		Order("COUNT(p.play_id) DESC").
		Limit(5).
		Preload("Artist").
		Preload("Artist.User").
		Preload("Album").
		Find(&res).Error
	return
}

func (s SongRepositoryImpl) GetSortedSong() (res []model.Song, err error) {
	err = s.DB.Table("songs AS s").
		Select("s.*").
		Joins("LEFT JOIN plays p ON s.song_id = p.song_id").
		Group("s.song_id,s.title,s.release_date,s.duration,s.file,s.album_id,s.artist_id").
		Order("play_count DESC").
		Preload("Artist").
		Preload("Artist.User").
		Preload("Album").
		Find(&res).Error
	return
}

func (s SongRepositoryImpl) GetSongByArtist(artistId string) (res []model.Song, err error) {
	songs, err := s.rdb.Get(utils.SongKey + artistId)
	if err != nil {
		err = s.DB.Where("artist_id = ?", artistId).Preload("Artist").Preload("Artist.User").Preload("Play").Preload("Album").Find(&res).Error
		if err != nil {
			return res, err
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			return res, err
		}
		_ = s.rdb.Set(utils.SongKey+artistId, string(resJSON))

		return res, nil
	} else {
		if err := json.Unmarshal([]byte(songs), &res); err != nil {
			return res, err
		}
		return res, nil
	}

}

func (s SongRepositoryImpl) GetSongByAlbum(albumId string) (res []model.Song, err error) {
	songs, err := s.rdb.Get(utils.SongKey + albumId)
	if err != nil {
		err = s.DB.Where("album_id = ?", albumId).Preload("Artist").Preload("Artist.User").Preload("Play").Preload("Album").Find(&res).Error

		if err != nil {
			return res, err
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			return res, err
		}
		_ = s.rdb.Set(utils.SongKey+albumId, string(resJSON))

		return res, nil
	} else {
		if err := json.Unmarshal([]byte(songs), &res); err != nil {
			return res, err
		}
		return res, nil
	}
}

func (s SongRepositoryImpl) CreateSong(song model.Song) error {
	err := s.rdb.Del(utils.SongKey)
	if err != nil {
		return err
	}
	err = s.DB.Create(&song).Error
	return err
}
