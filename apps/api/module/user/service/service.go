package service

import (
	"errors"
	"github.com/S1nceU/CRMS/apps/api/domain"
	"github.com/S1nceU/CRMS/apps/api/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const TokenExpireDuration = time.Hour * 2

type UserService struct {
	repo domain.UserRepository
}

type jwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) ListUser() ([]*model.User, error) {
	return u.repo.ListUser()
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	return u.repo.CreateUser(user)
}

func (u *UserService) UpdateUser(user *model.User) (*model.User, error) {
	return u.repo.UpdateUser(user)
}

func (u *UserService) DeleteUser(userId uuid.UUID) error {
	var err error
	newUser := &model.User{
		Id: userId,
	}
	err = u.repo.DeleteUser(newUser)
	return err
}

func (u *UserService) GetUserByUserId(userId uuid.UUID) (*model.User, error) {
	var err error
	newUser := &model.User{
		Id: userId,
	}
	newUser, err = u.repo.GetUserByUserId(newUser)
	return newUser, err
}

func (u *UserService) GetUserByUsername(username string) (*model.User, error) {
	var err error
	newUser := &model.User{
		Username: username,
	}
	newUser, err = u.repo.GetUserByUsername(newUser)
	return newUser, err
}

func (u *UserService) Login(username, password string) (string, error) {
	var err error
	newUser := &model.User{
		Username: username,
	}
	newUser, err = u.repo.GetUserByUsername(newUser)

	if err != nil {
		return "", err
	} else {
		if newUser.Password == "" {
			return "", errors.New("user not found")
		}
		if newUser.Password != password {
			return "", errors.New("password is incorrect")
		}
	}

	claim := jwtClaims{
		Username: newUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "S1nceU",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte("secret"))
}

func (u *UserService) Authentication(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	}, jwt.WithValidMethods([]string{"HS256"}))

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", errors.New("invalid token")
}
