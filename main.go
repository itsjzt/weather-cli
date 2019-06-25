package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	resp, err := http.Get("https://fcc-weather-api.glitch.me/api/current?lat=28&lon=72")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		type FCCData struct {
			Weather []struct {
				ID          int
				Description string
			}
			Main struct {
				Temp      float64
				Humidity  int
			}
		}

		var jsonData FCCData;
		json.Unmarshal([]byte(body), &jsonData)
		
		fmt.Printf("Temperature: %.2f deg Celsius \n", jsonData.Main.Temp)
		fmt.Printf("Description: %s \n", jsonData.Weather[0].Description)
		fmt.Printf("Humidity: %d \n", jsonData.Main.Humidity)
	}