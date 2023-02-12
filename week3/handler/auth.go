package Handler

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/database"
	"github.com/siddhuR/CipherSchools-course/tree/main/Week2/models"
)

func (h *Handler) SignIn(in *gin.Context) {
	user := models.Users{}
	err := in.BindJSON(&user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}

	err = database.SignIn(h.DB, &user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	log.Println("generating token")
	token, err := generateToken(user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	log.Println("generating token done")
	in.JSON(http.StatusOK, token)

}

func generateToken(user models.Users) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["emailid"] = user.EmailId
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString := token.SignedString("rahul")

	if err != nil {
		log.Println(err)
		return "", err
	}
	return token, nil
}
