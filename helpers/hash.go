// package helpers

// import "golang.org/x/crypto/bcrypt"

// func HashPassword(password string) (string, error) {
//     hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//     if err != nil {
//         return "", err
//     }
//     return string(hashedPassword), nil
// }

// 	func VerifyPassword(hashedPassword, password string) error {
// 	    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// 	}

// helpers/helpers.go
package helpers

import (
	"github.com/rachma1201/Task-5-pbi-btpns-Rahmaliyah-Kadir/database"
	"github.com/rachma1201/Task-5-pbi-btpns-Rahmaliyah-Kadir/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(user *models.User) error {
	db := database.GetDB()
	return db.Create(user).Error
}

func UpdateUser(userID string, user *models.User) error {
	db := database.GetDB()
	return db.Model(&models.User{}).Where("id = ?", userID).Updates(user).Error
}

func DeleteUser(userID string) error {
	db := database.GetDB()
	return db.Where("id = ?", userID).Delete(&models.User{}).Error
}

func AuthenticateUser(email, password string) (*models.User, error) {
	db := database.GetDB()
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}

func GenerateToken(user *models.User) (string, error) {
	// Implement token generation (e.g., JWT) here
	return "", nil
}
