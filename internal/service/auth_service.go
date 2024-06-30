package service

import (
	"synapsis-backend-test/internal/domain"
	"synapsis-backend-test/config"
	"synapsis-backend-test/pkg/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Register(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func Login(username, password string) (string, error) {
	var user domain.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.GenerateToken(user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func DetailUser(username string) (*domain.User, error) {
	var user *domain.User
	
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}


// func JwtSelf(c *fiber.Ctx) *domain.UserJwt {

// 	jwt_user := c.Locals("user").(*jwt.Token)
// 	claims := jwt_user.Claims.(jwt.MapClaims)

// 	user := domain.UserJwt{
// 		username: claims["username"].(string),
// 		role: claims["role"].(string),
// 		exp: claims["exp"].(int),
// 	}

// 	return &user
// }