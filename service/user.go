package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohibeyki/spock/model"
	"gorm.io/gorm"
)

// GetUser searches the DB for a specific user
func GetUser(db *gorm.DB, id string) (*model.User, error) {
	var err error
	user := new(model.User)

	if err := db.Where("id = ? ", id).First(&user).Error; err != nil {
		log.Println(err)

		return nil, err
	}

	return user, err
}

// GetUserByEmail searches the DB for a specific user
func GetUserByEmail(db *gorm.DB, email string) (*model.User, error) {
	var err error
	user := new(model.User)

	if err := db.Where("email = ? ", email).First(&user).Error; err != nil {
		log.Println(err)

		return nil, err
	}

	return user, err
}

// GetUsers returns a slice of users
func GetUsers(c *gin.Context, db *gorm.DB, args model.Args) ([]model.User, int64, int64, error) {
	users := []model.User{}
	var totalData int64

	table := "users"
	query := db.Select(table + ".*")
	query = query.Offset(Offset(args.Offset))
	query = query.Limit(Limit(args.Limit))
	query = query.Order(SortOrder(table, args.Sort, args.Order))
	query = query.Scopes(Search(args.Search))

	res := query.Find(&users)
	if err := res.Error; err != nil {
		log.Println(err)
		return users, 0, 0, err
	}

	db.Table(table).Count(&totalData)

	return users, res.RowsAffected, totalData, nil
}

// CreateUser creates a new user
func CreateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	if err := db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUser updates user
func UpdateUser(db *gorm.DB, user *model.User) (*model.User, error) {
	if err := db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// DeleteUser deletes a user
func DeleteUser(db *gorm.DB, id string) error {
	user := new(model.User)
	if err := db.Where("id = ? ", id).Delete(&user).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DeleteUserByEmail deletes a user by their email
func DeleteUserByEmail(db *gorm.DB, email string) error {
	user := new(model.User)
	if err := db.Where("email = ? ", email).Delete(&user).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
