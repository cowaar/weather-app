package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/cowaar/weather-app/util"
)
var config = util.ReadConfig()

func GetWeather(locationId string) string {
    
    apiKey := config.ApiKey
    url := config.FiveDayForecastEndpoint

    request := fmt.Sprintf("%s/%s?apikey=%s", url, locationId, apiKey)

	response, err := http.Get(request)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    return string(responseData)

}

func GetPokemon() {
      response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))

}

