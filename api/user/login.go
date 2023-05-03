package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"github.com/AlexBowmanCoding/content-hub-back-end/mongoDB"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Mongo struct{
	Client *mongo.Client
	DB mongodb.MongoDB
}

type BodyResponse struct{
	Err string `json:"err"`
	OtherData any `json:"otherData"`
}

func HashPassword(password string) (string, error) {
	bytePass := []byte(password)
    bytes, err := bcrypt.GenerateFromPassword(bytePass, 14)
	if err != nil{
		return "", err
	}
	hashedPassword := string(bytes)
    return hashedPassword, err
}

func CheckHashedPassword(storedPassword string, enteredPassword string) (bool, error) {
	byteOne := []byte(storedPassword)
	byteTwo := []byte(enteredPassword)
	err := bcrypt.CompareHashAndPassword(byteOne, byteTwo)
	if err != nil{
		return false, err
	}
	return true, err
	
}
func CheckUserLoginInfo(userinDB User, userInJSON User) (int, error){
	passwordbool, err := CheckHashedPassword(userinDB.Password, userInJSON.Password)
	
	if err != nil {
		return http.StatusInternalServerError, err
	}
		
	if userinDB.Username == userInJSON.Username && passwordbool {
		return http.StatusOK, nil
	} 
	
	return http.StatusBadRequest, errors.New("incorrect username or password")
}

func (user Mongo) NewUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		errNewUser := errors.New("invalid body requirements")
		w.WriteHeader(http.StatusBadRequest)
		response:= BodyResponse{
			Err: errNewUser.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	password, err := HashPassword(newUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		
		response := BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	newUser.Password = password

	userCollection := user.Client.Database("ContentHub").Collection("Users")

	err = user.DB.Post(*userCollection, newUser)
	if err != nil{
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)
}

func (user Mongo) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userInJSON User

	vars := mux.Vars(r)
	userId := vars["id"]
	err := json.NewDecoder(r.Body).Decode(&userInJSON)
	if err != nil {
		err = errors.New("invalid body requirements")
		w.WriteHeader(http.StatusBadRequest)
		response:= BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	userCollection := user.Client.Database("ContentHub").Collection("Users")
	result, err := user.DB.Get(*userCollection, userId)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		response:= BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	var userInDB User 

	userInDB.ID = result.ID
	userInDB.Password = result.Password
	userInDB.Username = result.Username

	if userInDB.ID == "" {
		err = errors.New("user not found")
		w.WriteHeader(http.StatusNotFound)
		response:= BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	status, err := CheckUserLoginInfo(userInDB, userInJSON)
	w.WriteHeader(status)
	response:= BodyResponse{
		OtherData: userInDB,
	}
	if err != nil{
		response.Err = err.Error()
	}
	json.NewEncoder(w).Encode(response)
}

func (user Mongo) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		errNewUser := errors.New("invalid body requirements")
		w.WriteHeader(http.StatusBadRequest)
		response:= BodyResponse{
			Err: errNewUser.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	password, err := HashPassword(newUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		
		response := BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	newUser.Password = password

	userCollection := user.Client.Database("ContentHub").Collection("Users")

	var mongoUser mongodb.User

	mongoUser.ID = userId
	mongoUser.Username = newUser.Username
	mongoUser.Password = newUser.Password
	err = user.DB.Update(*userCollection, mongoUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		response := BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)
}

func (user Mongo) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	userCollection := user.Client.Database("ContentHub").Collection("Users")

	err := user.DB.Delete(*userCollection, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		response := BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
}
