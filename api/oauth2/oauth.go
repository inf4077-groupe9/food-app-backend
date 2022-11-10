package oauth2

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateJWTToken(id primitive.ObjectID, username, email string) string {
	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// add data (preload)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["email"] = email
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// encoded token
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		panic("Error to encode token")
	}

	return t
}
