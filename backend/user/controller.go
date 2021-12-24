package user

import (
	"backend/config"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	_ "github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// _jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}


func (uc *UserController) CheckStatus(c *fiber.Ctx) error {
	
	return c.SendString("Pong")
}

// Sign-Up
func (uc *UserController) UserSignup(c *fiber.Ctx) error{

	um := UserModel{}

	var request UserSignup
	
	if err := c.BodyParser(&request); err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	usr, err := um.getUserbyEmail(request.Email)

	if err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting user email.",
		})
	}

	if (usr != UserInfo{}) {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already taken",
		})
	}

	request.CreateDate = time.Now()

 	err = um.userSignup(request)

	 if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while sign up. %s", err),
		})
	}
	
	return nil
}


// Login
func (uc *UserController) Login(c *fiber.Ctx) error {

	var request UserLogin
	if err := c.BodyParser(&request); err != nil {
		log.Error(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while binding to model.",
		})
	}

	user, err := uc.userLogin(request)

	if err != nil {
		log.Error(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error while login. %s", err),
		})
	}

	return c.JSON(user)
}


func (uc *UserController) userLogin(body UserLogin) (UserRequest, error) {


	um := UserModel{}
	userInfo, err := um.getUserbyEmail(body.Email)


	var userRequest UserRequest

	if err != nil {
		return userRequest, err
	}

	if (userInfo == UserInfo{}) {
		return userRequest, errors.New("user does not exist.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(body.Password))

	if err != nil {
		return userRequest, errors.New("incorrect password")
	}

	userRequest = UserRequest {
		UserId: userInfo.UserId,
		Email: userInfo.Email,
	}


	token, err := uc.generateJWT(userRequest.UserId)

	userRequest.AccessToken = token.AccessToken

	return userRequest, nil

}

//Token
func (UserController) generateJWT(userId primitive.ObjectID ) (TokenDetails, error) {

	config := config.GetConfig()

	var tokenDetails TokenDetails

	tokenDetails.AccessTokenExpires = time.Now().Add(time.Hour * 24 * 30).Unix()
	tokenDetails.AccessTokenUUID = uuid.NewV4().String()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["uuid"] = tokenDetails.AccessTokenUUID
	claims["user_id"] = userId
	claims["exp"] = tokenDetails.AccessTokenExpires
	
	tokenString, err := token.SignedString([]byte(config.GetString("auth.signinkey")))

	if err != nil {
		return tokenDetails, err
	}

	tokenDetails.AccessToken = tokenString

	return tokenDetails,err

}





