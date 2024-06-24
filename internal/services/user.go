package services

import (
	"blog/internal/config"
	"blog/internal/constants"
	"blog/internal/dto"
	models "blog/internal/models"
	repositories "blog/internal/repositories"
	errors "errors"
	"time"

	"gorm.io/gorm"

	"crypto/rand"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Register(userParams *dto.CreateUserDTO) (*models.User, error) {

	findUser, err := s.userRepository.FindUserByUserName(userParams.UserName)
	if findUser != nil {
		return nil, errors.New("this username has already been registered")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	salt := generateSalt()
	hashedPassword := hashPassword(userParams.Password, salt)

	user := &models.User{
		UserName:  userParams.UserName,
		Password:  hashedPassword,
		Salt:      salt,
		ID:        uuid.New().String(),
		Role:      userParams.Role,
		Email:     userParams.Email,
		Phone:     userParams.FullName,
		Avatar:    userParams.Avatar,
		Gender:    userParams.Gender,
		FullName:  userParams.FullName,
		Status:    constants.Active,
		Birthday:  userParams.Birthday,
		Address:   userParams.Address,
		CreatedAt: time.Now(),
	}

	return user, s.userRepository.Create(user)
}

func (s *UserService) Authenticate(userName string, password string) (*models.User, string, error) {
	user, err := s.userRepository.FindUserByUserName(userName)
	if user == nil {
		return nil, "", errors.New("the current username is not registered")
	}
	if err != nil {
		return nil, "", err
	}
	if !verifyPassword(password, user.Password, user.Salt) {
		return nil, "", errors.New("Password error")
	}

	token, err := generateToken(user.UserName, user.Password)

	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func generateSalt() string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(salt)
}

func hashPassword(password, salt string) string {
	combined := []byte(password + salt)
	hashedPassword, _ := bcrypt.GenerateFromPassword(combined, bcrypt.DefaultCost)
	return string(hashedPassword)
}

func verifyPassword(inputPassword, hashedPassword, salt string) bool {
	combined := []byte(inputPassword + salt)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), combined)
	return err == nil
}

func generateToken(userName string, hashedPassword string) (string, error) {
	// 设置 token 的过期时间为 7 天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   userName + hashedPassword,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
	return token.SignedString([]byte(cfg.SecretKey))
}
