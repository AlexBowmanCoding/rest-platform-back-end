package weather

import(
	owm "github.com/briandowns/openweathermap"
	"net/http"
	"errors"
	"encoding/json"
	"os"
)

type BodyResponse struct {
	Err       string `json:"err"`
	OtherData any    `json:"otherData"`
}

type WeatherAPI struct{
	apiKey string
}

type ForecastRequest struct{
	ZipCode int     `json:"zipCode"`
	TempType string `json:"tempType"`
}

func NewWeatherAPI() WeatherAPI {
	var api WeatherAPI
	api.apiKey = os.Getenv("OWM_API_KEY")
	return api
}

func (api WeatherAPI)GetForecastbyZip(r ForecastRequest) ( *owm.CurrentWeatherData, error){
	w, err := owm.NewCurrent("F","EN", api.apiKey)
	if err != nil {
		
		return nil, err
	}
	err = w.CurrentByZip(r.ZipCode, "US")
	if err != nil {
		print("err")
		return w, err
	}
	return w, nil
}

func (api WeatherAPI)GetWeather(w http.ResponseWriter, r *http.Request){
	api = NewWeatherAPI()

	var request ForecastRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		err = errors.New("invalid body requirements")
		w.WriteHeader(http.StatusBadRequest)
		response := BodyResponse{
			Err: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	weatherData, err := api.GetForecastbyZip(request)
	
	if err != nil {
		
		w.WriteHeader(http.StatusInternalServerError)
		response := BodyResponse{
			Err: err.Error(),
			OtherData: weatherData,
			
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weatherData)
}