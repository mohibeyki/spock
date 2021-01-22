package service

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/sjwt"
	"github.com/mohibeyki/spock/model"
	"github.com/mohibeyki/spock/pkg/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Offset returns the starting number of result for pagination
func Offset(offset string) int {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	return offsetInt
}

// Limit returns the number of result for pagination
func Limit(limit string) int {
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 25
	}
	return limitInt
}

// SortOrder returns the string for sorting and orderin data
func SortOrder(table, sort, order string) string {
	return table + "." + ToSnakeCase(sort) + " " + ToSnakeCase(order)
}

// Search adds where to search keywords
func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("name LIKE ?", "%"+search+"%")
			db = db.Or("description LIKE ?", "%"+search+"%")
		}
		return db
	}
}

// ToSnakeCase changes string to database table
func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

// GenerateToken generates a HS256 jwt token
func GenerateToken(user *model.User) string {
	claims := sjwt.New()
	claims.SetTokenID()                                 // UUID generated
	claims.SetSubject(user.Email)                       // Subject of the token
	claims.SetIssuer("Biook.me")                        // Issuer of the token
	claims.SetAudience([]string{"Biook.me", "Spock"})   // Audience the toke is for
	claims.SetIssuedAt(time.Now())                      // IssuedAt in time, value is set in unix
	claims.SetNotBeforeAt(time.Now())                   // Token is valid
	claims.SetExpiresAt(time.Now().Add(time.Hour * 24)) // Token expires in 24 hours
	claims.Set("avatar", user.Avatar)
	claims.Set("role", user.Role)

	config := config.GetConfig()
	return claims.Generate([]byte(config.Auth.PrivateKey))
}

// GetAndValidateToken checks if a token is valid, returns the payload if so
func GetAndValidateToken(token string) (model.User, error) {
	config := config.GetConfig()
	isVerified := sjwt.Verify(token, []byte(config.Auth.PrivateKey))
	user := model.User{}

	if !isVerified {
		return user, errors.New("bad token")
	}

	claims, err := sjwt.Parse(token)
	if err != nil {
		return user, err
	}
	if err := claims.Validate(); err != nil {
		return user, err
	}

	user.Email, _ = claims.GetStr("sub")
	user.Avatar, _ = claims.GetStr("avatar")
	user.Role, _ = claims.GetStr("role")
	return user, nil
}

// HashAndSalt is our password hashing function
func HashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePasswords compares hashed password with a plain one
func ComparePasswords(hash string, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		log.Println(err)
		return false
	}
	return true
}
