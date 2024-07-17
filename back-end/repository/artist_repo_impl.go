package repository

import (
	"back-end/data/response"
	"back-end/database"
	"back-end/model"
	"back-end/utils"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ArtistRepositoryImpl struct {
	DB  *gorm.DB
	rdb *database.Redis
}

func NewArtistRepositoryImpl(DB *gorm.DB, rdb *database.Redis) *ArtistRepositoryImpl {
	return &ArtistRepositoryImpl{DB: DB, rdb: rdb}
}

func (a ArtistRepositoryImpl) GetArtistByUserId(userId string, verified bool) (res model.Artist, err error) {
	if verified {
		verifiedArtist, err := a.rdb.Get(utils.VerificationRequestKey + userId)
		if err != nil {
			err = a.DB.Where("user_id = ? AND verified_at IS NOT NULL", userId).Preload("User").First(&res).Error
			if err != nil {
				return res, err
			}
			resJSON, err := json.Marshal(res)
			if err != nil {
				return res, err
			}
			_ = a.rdb.Set(utils.VerificationRequestKey+userId, string(resJSON))

			return res, nil
		} else {
			if err := json.Unmarshal([]byte(verifiedArtist), &res); err != nil {
				return res, err
			}
			return res, nil
		}
	} else {
		unverifiedArtist, err := a.rdb.Get(utils.UnVerificationRequestKey + userId)
		if err != nil {
			err = a.DB.Where("user_id = ? AND verified_at IS NULL", userId).Preload("User").First(&res).Error
			if err != nil {
				return res, err
			}
			resJSON, err := json.Marshal(res)
			if err != nil {
				return res, err
			}
			_ = a.rdb.Set(utils.UnVerificationRequestKey+userId, string(resJSON))

			return res, nil
		} else {
			if err := json.Unmarshal([]byte(unverifiedArtist), &res); err != nil {
				return res, err
			}
			return res, nil
		}
	}
}

func (a ArtistRepositoryImpl) GetArtistByArtistId(artistId string, verified bool) (res model.Artist, err error) {
	if verified {
		verifiedArtist, err := a.rdb.Get(utils.VerificationRequestKey + artistId)
		if err != nil {
			err = a.DB.Where("artist_id = ? AND verified_at IS NOT NULL", artistId).Preload("User").First(&res).Error
			if err != nil {
				return res, err
			}
			resJSON, err := json.Marshal(res)
			if err != nil {
				return res, err
			}
			_ = a.rdb.Set(utils.VerificationRequestKey+artistId, string(resJSON))

			return res, nil
		} else {
			if err := json.Unmarshal([]byte(verifiedArtist), &res); err != nil {
				return res, err
			}
			return res, nil
		}
	} else {
		unverifiedArtist, err := a.rdb.Get(utils.UnVerificationRequestKey + artistId)
		if err != nil {
			err = a.DB.Where("artist_id = ? AND verified_at IS NULL", artistId).Preload("User").First(&res).Error
			if err != nil {
				return res, err
			}
			resJSON, err := json.Marshal(res)
			if err != nil {
				return res, err
			}
			_ = a.rdb.Set(utils.UnVerificationRequestKey+artistId, string(resJSON))

			return res, nil
		} else {
			if err := json.Unmarshal([]byte(unverifiedArtist), &res); err != nil {
				return res, err
			}
			return res, nil
		}
	}
}

func (a ArtistRepositoryImpl) GetUnverifiedArtist() (res []model.Artist, err error) {
	unverifiedArtist, err := a.rdb.Get(utils.UnVerificationRequestKey)
	if err != nil {
		err = a.DB.Where("verified_at IS NULL").Preload("User").Find(&res).Error
		if err != nil {
			return res, err
		}
		resJSON, err := json.Marshal(res)
		if err != nil {
			return res, err
		}
		_ = a.rdb.Set(utils.UnVerificationRequestKey, string(resJSON))

		return res, nil
	} else {
		if err := json.Unmarshal([]byte(unverifiedArtist), &res); err != nil {
			return res, err
		}
		return res, nil
	}
}

func (a ArtistRepositoryImpl) CreateArtist(artist model.Artist) error {
	err := a.rdb.Del(utils.UnVerificationRequestKey)
	if err != nil {
		return err
	}
	err = a.DB.Create(&artist).Error
	return err
}

func (a ArtistRepositoryImpl) UpdateVerifyArtist(userId string, artistId string, verifiedAt time.Time) error {
	err := a.rdb.Del(utils.UnVerificationRequestKey)
	if err != nil {
		return err
	}
	err = a.rdb.Del(utils.UnVerificationRequestKey + artistId)
	if err != nil {
		return err
	}
	err = a.rdb.Del(utils.UnVerificationRequestKey + userId)
	if err != nil {
		return err
	}
	err = a.DB.Model(&model.Artist{}).Where("artist_id = ?", artistId).Update("verified_at", verifiedAt).Error
	return err
}

func (a ArtistRepositoryImpl) DeleteArtist(userId string, artistId string) error {
	err := a.rdb.Del(utils.UnVerificationRequestKey)
	if err != nil {
		return err
	}
	err = a.rdb.Del(utils.UnVerificationRequestKey + artistId)
	if err != nil {
		return err
	}
	err = a.rdb.Del(utils.UnVerificationRequestKey + userId)
	if err != nil {
		return err
	}
	err = a.DB.Where("artist_id = ?", artistId).Delete(&model.Artist{}).Error
	return err
}

func (a ArtistRepositoryImpl) GetArtistByName(name string) (res []response.ArtistSearch, err error) {

	err = a.DB.Table("artists AS a").
		Select("a.artist_id,a.user_id,u.username, COUNT(f.following_id) AS follow_count").
		Joins("LEFT JOIN users u ON a.user_id = u.user_id").
		Joins("LEFT JOIN follows f ON f.following_id = a.user_id").
		Where("UPPER(u.username) LIKE ?", "%"+name+"%").
		Group("a.user_id, u.username,a.artist_id").
		Order("follow_count DESC").
		Limit(6).
		Scan(&res).Error
	return
}
