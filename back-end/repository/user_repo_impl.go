package repository

import (
	"back-end/database"
	"back-end/model"
	"back-end/utils"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type UserRepositoryImpl struct {
	DB  *gorm.DB
	rdb *database.Redis
}

func NewUserRepositoryImpl(DB *gorm.DB, rdb *database.Redis) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: DB, rdb: rdb}
}

func (u UserRepositoryImpl) Save(user model.User) error {
	err := u.DB.Create(&user).Error
	return err
}

func (u UserRepositoryImpl) FindAll() (users []model.User, err error) {

	redisUser, err := u.rdb.Get("users")
	if err != nil {
		err = u.DB.Where("verified_at IS NOT NULL").Preload("NotificationSetting").Find(&users).Error
		if err != nil {
			return
		}
		userJSON, err := json.Marshal(users)
		if err != nil {
			return users, err
		}
		_ = u.rdb.Set("users", string(userJSON))

		return users, nil
	} else {
		if err := json.Unmarshal([]byte(redisUser), &users); err != nil {
			return users, err
		}
		return users, nil
	}

}

func (u UserRepositoryImpl) FindUserByID(id string) (user model.User, err error) {
	err = u.DB.Where("user_id = ?", id).Preload("NotificationSetting").First(&user).Error
	return
}

func (u UserRepositoryImpl) FindByEmailAndVerified(id string, verified bool) (user model.User, err error) {
	if verified {
		err = u.DB.Where("email = ? AND verified_at IS NOT NULL", id).Preload("NotificationSetting").First(&user).Error
	} else {
		err = u.DB.Where("email = ? AND verified_at IS NULL", id).Preload("NotificationSetting").First(&user).Error
	}
	return
}

func (u UserRepositoryImpl) FindByEmail(email string) (user model.User, err error) {
	err = u.DB.Where("email = ?", email).Preload("NotificationSetting").First(&user).Error
	return
}

func (u UserRepositoryImpl) Update(user model.User) (err error) {
	err = u.rdb.Del(utils.CurrentUserKey + user.UserId)
	if err != nil {
		return err
	}
	err = u.rdb.Del(utils.PlaylistKey + user.UserId)
	if err != nil {
		return err
	}
	err = u.rdb.Del(utils.PlaylistDetailKey + user.UserId)
	if err != nil {
		return err
	}
	err = u.rdb.Del("users")
	if err != nil {
		return err
	}
	err = u.DB.Updates(&user).Error
	return
}

func (u UserRepositoryImpl) Delete(user model.User) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImpl) UpdateRole(userId string) (err error) {
	err = u.rdb.Del(utils.CurrentUserKey + userId)
	if err != nil {
		return err
	}
	err = u.rdb.Del(utils.PlaylistKey + userId)
	if err != nil {
		return err
	}
	err = u.rdb.Del(utils.PlaylistDetailKey + userId)
	if err != nil {
		return err
	}
	err = u.rdb.Del("users")
	if err != nil {
		return err
	}
	err = u.DB.Model(&model.User{}).Where("user_id = ?", userId).Update("role", "Artist").Error
	return
}

func (u UserRepositoryImpl) UpdateGoogleId(userId string, email string, googleId string) (err error) {
	err = u.rdb.Del(utils.CurrentUserKey + userId)
	if err != nil {
		return err
	}
	err = u.rdb.Del(utils.PlaylistKey + userId)
	if err != nil {
		return err
	}
	err = u.rdb.Del("users")
	if err != nil {
		return err
	}
	err = u.DB.Model(&model.User{}).Where("email = ?", email).Update("google_id", googleId).Error
	return
}

func (u UserRepositoryImpl) GetCurrentUser(userId string) (user model.User, err error) {
	redisUser, err := u.rdb.Get(utils.CurrentUserKey + userId)
	if err != nil {
		err = u.DB.Where("user_id = ?", userId).Preload("NotificationSetting").First(&user).Error

		if err != nil {
			return
		}
		userJSON, err := json.Marshal(user)
		if err != nil {
			return user, err
		}
		_ = u.rdb.Set(utils.CurrentUserKey+userId, string(userJSON))

		return user, nil
	} else {
		if err := json.Unmarshal([]byte(redisUser), &user); err != nil {
			return user, err
		}
		return user, nil
	}
}

func (u UserRepositoryImpl) GetUserByVerifyLink(verifyLink string) (user model.User, err error) {
	err = u.DB.Where("verify_link = ?", verifyLink).First(&user).Error
	return
}

func (u UserRepositoryImpl) UpdateVerifyLink(userId string, verifyLink string) (err error) {
	err = u.DB.Model(&model.User{}).Where("user_id = ?", userId).Update("verify_link", verifyLink).Error
	return
}

func (u UserRepositoryImpl) UpdateRegister(userId string, verifyLink string, username string, hash []byte) (err error) {
	err = u.DB.Model(&model.User{}).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"verify_link": verifyLink,
		"username":    username,
		"password":    hash,
	}).Error
	err = u.rdb.Del(utils.PlaylistKey + userId)
	if err != nil {
		return err
	}

	err = u.rdb.Del(utils.PlaylistDetailKey + userId)
	if err != nil {
		return err
	}

	err = u.rdb.Del("users")
	if err != nil {
		return err
	}
	return
}

func (u UserRepositoryImpl) ChangePassword(password []byte, userId string) error {
	err := u.DB.Model(&model.User{}).Where("user_id = ?", userId).Update("password", password).Error
	return err
}

func (u UserRepositoryImpl) Logout(userId string) error {
	err := u.rdb.Del(utils.CurrentUserKey + userId)
	return err
}

func (u UserRepositoryImpl) UpdateProfilePicture(userId string, avatar string) error {
	err := u.rdb.Del(utils.CurrentUserKey + userId)
	if err != nil {
		return err
	}
	err = u.rdb.Del(utils.PlaylistKey + userId)
	if err != nil {
		return err
	}

	err = u.rdb.Del(utils.PlaylistDetailKey + userId)
	if err != nil {
		return err
	}

	err = u.rdb.Del("users")
	if err != nil {
		return err
	}
	err = u.DB.Model(&model.User{}).Where("user_id = ?", userId).Update("avatar", avatar).Error
	return err
}

func (u UserRepositoryImpl) UpdateProfile(userId string, dob time.Time, country string, gender string) error {
	err := u.rdb.Del(utils.CurrentUserKey + userId)
	if err != nil {
		return err
	}

	err = u.rdb.Del(utils.PlaylistKey + userId)
	if err != nil {
		return err
	}

	err = u.rdb.Del(utils.PlaylistDetailKey + userId)
	if err != nil {
		return err
	}

	err = u.rdb.Del("users")
	if err != nil {
		return err
	}

	err = u.DB.Model(&model.User{}).Where("user_id = ?", userId).Updates(map[string]interface{}{
		"dob":     dob,
		"country": country,
		"gender":  gender}).Error
	return err
}
