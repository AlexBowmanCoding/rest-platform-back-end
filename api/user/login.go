package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"main/mongoDB"
	"github.com/gorilla/mux"
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

func ReadJsonDB() (*[]User, error) {
	var users []User
	var user User
	usersBytes, err := ioutil.ReadFile("users.json")
	if err != nil {
		return nil, err
	}
	log.Print(usersBytes)
	err = json.Unmarshal(usersBytes, &user)
	if err != nil{
		return nil, err
	}
	users = append(users, user)
	return &users, nil
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

func LoginUser(w http.ResponseWriter, r *http.Request) {
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

	//TEMP DB
	users, err := ReadJsonDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response:= BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	for _, user := range *users {
		
		if user.ID == userId {
			passwordbool, err := CheckHashedPassword(user.Password, userInJSON.Password)
			if err != nil {
				response:= BodyResponse{
					Err: err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response)
				return
			}
			
			if user.Username == userInJSON.Username && passwordbool {
				w.WriteHeader(http.StatusOK)
				return
			} else {
				err = errors.New("incorrect username or password")
				w.WriteHeader(http.StatusBadRequest)
				response:= BodyResponse{
					Err: err.Error(),
				}
				json.NewEncoder(w).Encode(response)
				return
			}
		}
	}

	err = errors.New("user not found")
	w.WriteHeader(http.StatusNotFound)
	response:= BodyResponse{
		Err: err.Error(),
	}
	json.NewEncoder(w).Encode(response)
	
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := errors.New("delete user not implemented yet")
	w.WriteHeader(http.StatusInternalServerError)
	response:= BodyResponse{
		Err: err.Error(),
	}
	json.NewEncoder(w).Encode(response)
}
