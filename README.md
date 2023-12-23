# Rest Platform BackEnd

## A Rest API for connecting to various third party API and User Authentication


| Endpoint                    	| Method 	| JSON Parameters                                            	| Description                          	       |
|-----------------------------	|--------	|------------------------------------------------------------	|----------------------------------------------|
| /weather                      | GET     | zipCode: int, tempType: str                                 | Returns weather data for the zipcode provided|
| /users/<int:id>               | GET   	|                                                             | Returns user data                          	 |
| /users/<int:id>               | PUT    	| id: int, username: str, password: str                       | Updates User data                            |
| /users/<int:id>               | DELETE 	|                                                             | Deletes a User                     	         |
| /users                        | POST    | id: int, username: str, password: str                       | Creates a new User                  	       |
| /users/login/<int:id>         | GET 	  | id: int, username: str, password: str              	        | User Authenitcation                 	       |

