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
func (r *UserRepositoryImpl) Create(user *models.User) (*models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	count, err := r.collection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		return nil, errors.New("error occured while checking for the email")
	}

	if count > 0 {
		return nil, errors.New("this email already exsits")
	}

	if models.IsRoleValid(user.UserRole) == false {
		return nil, errors.New("invalid role")
	}

	password := utils.HashPassword(*user.Password)
	user.Password = &password

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()

	_, err = r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.New("user could not be created")
	}

	return user, nil
}

func (r *UserRepositoryImpl) Login(user *models.User) (*models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundUser models.User

	err := r.collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return nil, errors.New("user not found, login seems to be incorrect")
	}
	if !(utils.VerifyPassword(*user.Password, *foundUser.Password)) {
		return nil, errors.New("email or password incorrect")
	}
	if !foundUser.Deleted_at.Equal(time.Time{}) {
		return nil, errors.New("User not found")
	}

	filter := bson.M{"email": user.Email}
	user.Last_login_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{Key: "last_login_at", Value: user.Last_login_at})

	_, insertErr := r.collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}})
	if insertErr != nil {
		return nil, errors.New("Login failed")
	}
	defer cancel()

	return user, nil
}
func (r *UserRepositoryImpl) Logout(email string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundUser *models.User

	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return errors.New("User not found")
	}
	if !foundUser.Deleted_at.Equal(time.Time{}) {
		return errors.New("User not found")
	}

	filter := bson.M{"email": *foundUser.Email}

	var updateObj primitive.D
	*&foundUser.Logout_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateObj = append(updateObj, bson.E{Key: "logout_at", Value: *&foundUser.Logout_at})

	_, logoutErr := r.collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}})
	if logoutErr != nil {
		return errors.New("Failed to log out")
	}
	defer cancel()

	return nil
}

func (r *UserRepositoryImpl) Delete(email string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundUser *models.User

	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return errors.New("User not found")
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
		return errors.New("user could not be deleted")
	}
	defer cancel()

	return nil
}

func (r *UserRepositoryImpl) GetUser(email string) (*models.GetResponse, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var responseUser *models.GetResponse

	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&responseUser)
	defer cancel()
	if err != nil {
		return responseUser, errors.New("User not found")
	}

	return responseUser, nil
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foundUser *models.User

	err := r.collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return errors.New("User not found")
	}

	filter := bson.M{"email": user.Email}
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
			return errors.New("this email is already exist")
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
		return errors.New("Failed to update user information")
	}
	defer cancel()
	user = foundUser
	return nil
}
