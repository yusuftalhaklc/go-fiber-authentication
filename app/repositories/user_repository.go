package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/yusuftalhaklc/go-fiber-authentication/app/config"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/models"
	"github.com/yusuftalhaklc/go-fiber-authentication/app/utils"
	"github.com/yusuftalhaklc/go-fiber-authentication/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepositoryImpl {
	database.Connect()
	DB := database.Db

	collection := DB.Collection(config.Config("DB_COLLECTION"))

	return &UserRepositoryImpl{
		collection: collection,
	}
}
func (r *UserRepositoryImpl) Create(user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	count, err := r.collection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		return errors.New("error occured while checking for the email")
	}

	if count > 0 {
		return errors.New("this email or phone number already exsits")
	}
	password := utils.HashPassword(*user.Password)
	user.Password = &password

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()

	token, err := utils.CreateToken(user)
	if err != nil {
		return errors.New("token was not created")
	}
	user.Token = &token

	_, err = r.collection.InsertOne(ctx, user)
	if err != nil {
		return errors.New("User was not created")
	}

	return nil
}

func (r *UserRepositoryImpl) Login(user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundUser models.User

	err := r.collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return errors.New("user not found, login seems to be incorrect")
	}
	if !(utils.VerifyPassword(*user.Password, *foundUser.Password)) {
		return errors.New("email or password incorrect")
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		return errors.New("token was not created")
	}

	user.Token = &token
	user.Last_login_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	filter := bson.M{"email": user.Email}

	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{Key: "token", Value: user.Token})
	updateObj = append(updateObj, bson.E{Key: "last_login_at", Value: user.Last_login_at})

	_, insertErr := r.collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}})
	if insertErr != nil {
		return errors.New("login error")
	}
	defer cancel()

	return nil
}
