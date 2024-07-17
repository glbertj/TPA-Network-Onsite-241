package services

import (
	"back-end/data/request"
	"back-end/data/response"
	"back-end/model"
	"back-end/repository"
	"back-end/utils"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserServiceImpl struct {
	UserRepository                repository.UserRepository
	NotificationSettingRepository repository.NotificationSettingRepository
	Validate                      *validator.Validate
}

func NewUserServiceImpl(UserRepository repository.UserRepository, NotificationSettingRepository repository.NotificationSettingRepository, Validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: UserRepository, NotificationSettingRepository: NotificationSettingRepository, Validate: Validate}

}

func (u UserServiceImpl) Authenticate(req request.AuthRequest) (res response.AuthResponse, err error) {
	user, err := u.UserRepository.FindByEmailAndVerified(req.Email, true)
	if err != nil {
		return response.AuthResponse{}, utils.UserNotFound
	}

	if err = bcrypt.CompareHashAndPassword(*user.Password, []byte(req.Password)); err != nil {
		return response.AuthResponse{}, utils.InvalidPassword
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return response.AuthResponse{}, err
	}

	authResponse := response.AuthResponse{
		UserId:   user.UserId,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}

	return authResponse, nil

}

func (u UserServiceImpl) Register(req request.UserRequest) (res response.RegisterResponse, err error) {
	err = u.Validate.Struct(req)

	if err != nil {

		return response.RegisterResponse{}, err
	}

	user, err := u.UserRepository.FindByEmailAndVerified(req.Email, true)
	if err == nil {
		return response.RegisterResponse{}, utils.UserAlreadyExist
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.RegisterResponse{}, err
	}

	user, err = u.UserRepository.FindByEmailAndVerified(req.Email, false)
	linkUUID := utils.GenerateUUID()

	if err != nil {
		notificationSetting := model.NotificationSetting{
			NotificationSettingId: utils.GenerateUUID(),
			EmailFollower:         false,
			EmailAlbum:            false,
			WebFollower:           false,
			WebAlbum:              false,
		}

		err = u.NotificationSettingRepository.Create(notificationSetting)
		if err != nil {
			return response.RegisterResponse{}, err
		}

		user = model.User{
			UserId:                utils.GenerateUUID(),
			Username:              req.Username,
			Password:              &hashedPassword,
			GoogleId:              nil,
			Role:                  "Listener",
			VerifiedAt:            nil,
			Email:                 req.Email,
			Gender:                nil,
			Country:               nil,
			Avatar:                nil,
			Dob:                   nil,
			VerifyLink:            &linkUUID,
			NotificationSettingId: notificationSetting.NotificationSettingId,
		}
		err = u.UserRepository.Save(user)
		if err != nil {
			return response.RegisterResponse{}, err
		}
	} else {
		user.Username = req.Username
		user.Password = &hashedPassword
		user.VerifyLink = &linkUUID
		err = u.UserRepository.UpdateRegister(user.UserId, linkUUID, req.Username, hashedPassword)
		if err != nil {
			return response.RegisterResponse{}, err
		}
	}

	return response.RegisterResponse{
		UserId:     user.UserId,
		Username:   user.Username,
		Email:      user.Email,
		Role:       user.Role,
		VerifyLink: linkUUID,
	}, nil
}

func (u UserServiceImpl) UpdateVerificationStatus(id string) error {
	user, err := u.UserRepository.GetUserByVerifyLink(id)
	if err != nil {
		return err
	}
	now := time.Now()
	user.VerifiedAt = &now
	err = u.UserRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserServiceImpl) GetCurrentUser(cookie string) (res response.UserResponse, err error) {
	userId, err := utils.GetJWTClaims(cookie)

	user, err := u.UserRepository.GetCurrentUser(userId)
	if err != nil {
		return response.UserResponse{}, utils.UserNotFound
	}
	return response.UserResponse{
		UserId:                user.UserId,
		Username:              user.Username,
		Email:                 user.Email,
		Role:                  user.Role,
		Avatar:                user.Avatar,
		Country:               user.Country,
		Gender:                user.Gender,
		Dob:                   user.Dob,
		NotificationSettingId: user.NotificationSettingId,
		NotificationSetting:   user.NotificationSetting,
	}, nil

}

func (u UserServiceImpl) LoginWithGoogle(req request.GoogleRequest) (res response.AuthResponse, err error) {
	now := time.Now()
	user := model.User{
		UserId:     utils.GenerateUUID(),
		Username:   req.Username,
		Password:   nil,
		GoogleId:   &req.GoogleId,
		Role:       "Listener",
		VerifiedAt: &now,
		Email:      req.Email,
		Gender:     nil,
		Country:    nil,
		Avatar:     nil,
		Dob:        nil,
	}

	user2, err := u.UserRepository.FindByEmail(req.Email)
	if err == nil {
		user2.GoogleId = &req.GoogleId
		user2.Password = nil
		err = u.UserRepository.UpdateGoogleId(user2.UserId, req.Email, req.GoogleId)
		if err != nil {
			return response.AuthResponse{}, err
		}
		user = user2
	} else {
		notificationSetting := model.NotificationSetting{
			NotificationSettingId: utils.GenerateUUID(),
			EmailFollower:         false,
			EmailAlbum:            false,
			WebFollower:           false,
			WebAlbum:              false,
		}

		err = u.NotificationSettingRepository.Create(notificationSetting)
		if err != nil {
			return response.AuthResponse{}, err
		}
		user.NotificationSettingId = notificationSetting.NotificationSettingId
		err = u.UserRepository.Save(user)
		if err != nil {
			return response.AuthResponse{}, err
		}
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return response.AuthResponse{}, err
	}
	return response.AuthResponse{
		UserId:   user.UserId,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}, nil
}

func (u UserServiceImpl) GetUserById(id string) (res response.UserResponse, err error) {
	user, err := u.UserRepository.FindUserByID(id)
	if err != nil {
		return response.UserResponse{}, utils.UserNotFound
	}
	return response.UserResponse{
		UserId:                user.UserId,
		Username:              user.Username,
		Email:                 user.Email,
		Role:                  user.Role,
		Avatar:                user.Avatar,
		Country:               user.Country,
		Gender:                user.Gender,
		Dob:                   user.Dob,
		NotificationSettingId: user.NotificationSettingId,
		NotificationSetting:   user.NotificationSetting,
	}, nil
}

func (u UserServiceImpl) UpdateUserProfile(req request.UserUpdateRequest) (res response.UserResponse, err error) {
	err = u.Validate.Struct(req)
	if err != nil {
		return response.UserResponse{}, err
	}

	err = u.UserRepository.UpdateProfile(req.UserId, req.Dob, req.Country, req.Gender)
	if err != nil {
		return response.UserResponse{}, err
	}

	res = response.UserResponse{}

	return

}

func (u UserServiceImpl) GetAllUser() (res []response.UserResponse, err error) {
	users, err := u.UserRepository.FindAll()
	if err != nil {
		return
	}
	for _, user := range users {
		res = append(res, response.UserResponse{
			UserId:                user.UserId,
			Username:              user.Username,
			Email:                 user.Email,
			Role:                  user.Role,
			Avatar:                user.Avatar,
			Country:               user.Country,
			Gender:                user.Gender,
			Dob:                   user.Dob,
			NotificationSettingId: user.NotificationSettingId,
			NotificationSetting:   user.NotificationSetting,
		})
	}
	return
}

func (u UserServiceImpl) ForgotPassword(email string) (res response.RegisterResponse, err error) {
	user, err := u.UserRepository.FindByEmailAndVerified(email, true)
	if err != nil {
		return res, utils.UserNotFound
	}

	linkUUID := utils.GenerateUUID()
	user.VerifyLink = &linkUUID
	err = u.UserRepository.UpdateVerifyLink(user.UserId, linkUUID)
	if err != nil {
		return res, utils.UserNotFound
	}

	return response.RegisterResponse{
		UserId:     user.UserId,
		Username:   user.Username,
		Email:      user.Email,
		Role:       user.Role,
		VerifyLink: linkUUID,
	}, nil
}

func (u UserServiceImpl) GetUserByVerifyLink(id string) (res response.UserResponse, err error) {
	user, err := u.UserRepository.GetUserByVerifyLink(id)
	if err != nil {
		return res, err
	}
	return response.UserResponse{
		UserId:                user.UserId,
		Username:              user.Username,
		Email:                 user.Email,
		Role:                  user.Role,
		Avatar:                user.Avatar,
		Country:               user.Country,
		Gender:                user.Gender,
		Dob:                   user.Dob,
		NotificationSettingId: user.NotificationSettingId,
		NotificationSetting:   user.NotificationSetting,
	}, nil

}

func (u UserServiceImpl) ChangePassword(password string, userId string) error {
	user, err := u.UserRepository.FindUserByID(userId)
	if err != nil {
		return err
	}

	if user.Password != nil {
		if err = bcrypt.CompareHashAndPassword(*user.Password, []byte(password)); err == nil {
			return utils.SamePassword
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = u.UserRepository.ChangePassword(hashedPassword, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u UserServiceImpl) Logout(userId string) error {
	err := u.UserRepository.Logout(userId)
	if err != nil {
		return err
	}
	return nil
}

func (u UserServiceImpl) UpdateProfilePicture(userId string, avatar string) error {
	err := u.UserRepository.UpdateProfilePicture(userId, avatar)
	if err != nil {
		return err
	}
	return nil
}
