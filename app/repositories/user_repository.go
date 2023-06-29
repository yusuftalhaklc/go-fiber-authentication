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
	if !foundUser.Deleted_at.Equal(time.Time{}) {
		return errors.New("User not found")
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
func (r *UserRepositoryImpl) Logout(token *string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundUser *models.User

	err := r.collection.FindOne(ctx, bson.M{"token": *token}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return errors.New("invalid credentials")
	}
	if !foundUser.Deleted_at.Equal(time.Time{}) {
		return errors.New("User not found")
	}

	newToken, _ := utils.InvalidateToken(*token)

	filter := bson.M{"email": *foundUser.Email}

	var updateObj primitive.D
	*&foundUser.Logout_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateObj = append(updateObj, bson.E{Key: "logout_at", Value: *&foundUser.Logout_at})
	updateObj = append(updateObj, bson.E{Key: "token", Value: newToken})

	_, insertErr := r.collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}})
	if insertErr != nil {
		return errors.New("cannot logout")
	}
	defer cancel()

	return nil
}

func (r *UserRepositoryImpl) Delete(token *string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundUser *models.User

	err := r.collection.FindOne(ctx, bson.M{"token": *token}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return errors.New("invalid credentials")
	}

	filter := bson.M{"email": *foundUser.Email}

	if !foundUser.Deleted_at.Equal(time.Time{}) {
		return errors.New("user already deleted")
	}

	var updateObj primitive.D
	*&foundUser.Deleted_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateObj = append(updateObj, bson.E{Key: "deleted_at", Value: *&foundUser.Deleted_at})

	_, deleteErr := r.collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}})
	if deleteErr != nil {
		return errors.New("cannot deleted")
	}
	defer cancel()

	return nil
}

func (r *UserRepositoryImpl) GetUser(token *string) (*models.GetResponse, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundUser *models.User
	var responseUser *models.GetResponse

	err := r.collection.FindOne(ctx, bson.M{"token": *token}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return responseUser, errors.New("invalid credentials")
	}

	filter := bson.M{"token": *foundUser.Token}

	insertErr := r.collection.FindOne(ctx, filter).Decode(&responseUser)
	if insertErr != nil {
		return responseUser, errors.New("something went wrong")
	}
	defer cancel()

	return responseUser, nil
}

func (r *UserRepositoryImpl) Update(user *models.User, token *string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundUser *models.User

	err := r.collection.FindOne(ctx, bson.M{"token": *token}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return errors.New("invalid credentials")
	}

	filter := bson.M{"token": *foundUser.Token}
	var updateObj primitive.D

	if user.First_name != nil {
		updateObj = append(updateObj, bson.E{Key: "first_name", Value: *user.First_name})
	}
	if user.Last_name != nil {
		updateObj = append(updateObj, bson.E{Key: "last_name", Value: *user.Last_name})
	}
	if user.Password != nil {
		*user.Password = utils.HashPassword(*user.Password)
		updateObj = append(updateObj, bson.E{Key: "password", Value: *user.Password})
	}
	if user.Email != nil {
		err := r.collection.FindOne(ctx, bson.M{"email": *user.Email}).Decode(&foundUser)
		if err == nil {
			return errors.New("email already taken")
		} else {
			updateObj = append(updateObj, bson.E{Key: "email", Value: *user.Email})
		}

	}
	if user.Phone != nil {
		updateObj = append(updateObj, bson.E{Key: "phone", Value: *user.Phone})
	}
	if user.Avatar != nil {
		updateObj = append(updateObj, bson.E{Key: "avatar", Value: *user.Avatar})
	}

	_, updateErr := r.collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}})
	if updateErr != nil {
		return errors.New("cannot updated")
	}
	defer cancel()
	user = foundUser
	return nil
}
