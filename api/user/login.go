package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)



// BodyResponse  struct for holding the json repsonse's body.
type BodyResponse struct {
	Err       string `json:"err"`
	OtherData any    `json:"otherData"`
}

// HashPassword encrypts a password before storing it in the database.
func HashPassword(password string) (string, error) {
	bytePass := []byte(password)
	bytes, err := bcrypt.GenerateFromPassword(bytePass, 14)
	if err != nil {
		return "", err
	}
	hashedPassword := string(bytes)
	return hashedPassword, err
}

// CheckHashedPassword checks a hashed password against and unencrypted password to see if they are he same for login authentication.
func CheckHashedPassword(storedPassword string, enteredPassword string) (bool, error) {
	byteOne := []byte(storedPassword)
	byteTwo := []byte(enteredPassword)
	err := bcrypt.CompareHashAndPassword(byteOne, byteTwo)
	if err != nil {
		return false, err
	}
	return true, err

}

// CheckUserLoginInfo checks login info from the userInJSON struct against the userinDB struct to see if they are the same for login authentication.
func CheckUserLoginInfo(userinDB User, userInJSON User) (int, error) {
	passwordbool, err := CheckHashedPassword(userinDB.Password, userInJSON.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if userinDB.Username == userInJSON.Username && passwordbool {
		return http.StatusOK, nil
	}

	return http.StatusBadRequest, errors.New("incorrect username or password")
}

// NewUser Creates a new user in the mongoDB and encrypts the password given from the json.
func (user MongoUser) NewUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	
	//Grab user data from json.
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		errNewUser := errors.New("invalid body requirements")
		w.WriteHeader(http.StatusBadRequest)
		response := BodyResponse{
			Err: errNewUser.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	//Hash password from json.
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

	// Puts the User data in the mongoDB.
	err = user.Post(*userCollection, newUser)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)
}

// LoginUser checks login data against json user data to authenticate login 
func (user MongoUser) LoginUser(w http.ResponseWriter, r *http.Request) {
	var userInJSON User

	vars := mux.Vars(r)
	userId := vars["id"]

	//Grab user data from json.
	err := json.NewDecoder(r.Body).Decode(&userInJSON)
	if err != nil {
		err = errors.New("invalid body requirements")
		w.WriteHeader(http.StatusBadRequest)
		response := BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	//Get user data from the mongo database.
	userCollection := user.Client.Database("ContentHub").Collection("Users")
	result, err := user.Get(*userCollection, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := BodyResponse{
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
		response := BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Authenticate user login info
	status, err := CheckUserLoginInfo(userInDB, userInJSON)
	w.WriteHeader(status)
	response := BodyResponse{
		OtherData: userInDB,
	}
	if err != nil {
		response.Err = err.Error()
	}
	json.NewEncoder(w).Encode(response)
}


// UpdateUser takes user data from json and updates user data in the database.
func (user MongoUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var newUser User

	//Grab user data from json.
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		errNewUser := errors.New("invalid body requirements")
		w.WriteHeader(http.StatusBadRequest)
		response := BodyResponse{
			Err: errNewUser.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	//Hash password from json.
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

	var mongoUser User

	mongoUser.ID = userId
	mongoUser.Username = newUser.Username
	mongoUser.Password = newUser.Password

	//Update user in database.
	err = user.Update(*userCollection, mongoUser)
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


// DeleteUser Deletes a user from the database.
func (user MongoUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	userCollection := user.Client.Database("ContentHub").Collection("Users")

	err := user.Delete(*userCollection, userId)
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

// GetUser Gets user data from the database excluding the password.
func (user MongoUser) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["id"]


	//Get user data from the mongo database.
	userCollection := user.Client.Database("ContentHub").Collection("Users")
	result, err := user.Get(*userCollection, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	result.Password = "Password not included for security reasons."

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
