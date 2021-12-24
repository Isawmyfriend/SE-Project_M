package user

import (
	"backend/database"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct{}

func (um UserModel) userSignup(body UserSignup) (err error) {

	db,ctx := database.DB()

	err = db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)
	

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	body.Password = string(encryptedPassword)

	seDatabase := db.Database("se_project")
	userCollection := seDatabase.Collection("users")

	_ ,err = userCollection.InsertOne(ctx,body)

	if err != nil {
		log.Fatal(err)
	}



	return nil

}

func (um UserModel) getUserbyEmail(email string) (UserInfo, error){

	var user UserInfo

	db,ctx := database.DB()

	err := db.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Disconnect(ctx)

	seDatabase := db.Database("se_project")
	userCollection := seDatabase.Collection("users")

	userWithEmail := userCollection.FindOne(ctx,bson.M{"email": email}).Decode(&user);

	if userWithEmail == nil {
		return user, err
	}

	return user,nil


}






