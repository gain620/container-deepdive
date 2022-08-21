package cmd

import (
	"fmt"
	"github.com/gain620/weatherctl/config"
	"github.com/gain620/weatherctl/model"
	"github.com/imroc/req/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

var weatherCmd = &cobra.Command{
	Use:     "weather",
	Short:   "Get current weather info according to your current location",
	Example: "conctl weather [options] [args]",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Requesting the current temperature from the weather API server ... \n")

		city, err := cmd.Flags().GetString("city")
		if err != nil {
			log.Fatal(err)
		}

		temp, err := cmd.Flags().GetString("temp")
		if err != nil {
			log.Fatal(err)
		}

		aqi, err := cmd.Flags().GetString("aqi")
		if err != nil {
			log.Fatal(err)
		}

		weatherInfo := &model.WeatherInfo{}

		resp, err := req.C().
			SetUserAgent(config.UserAgent).
			SetTimeout(5 * time.Second).R().
			SetPathParams(map[string]string{
				"api_key": config.WeatherAPIKey,
				"city":    city,
			}).
			Get(fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key={api_key}&q={city}&aqi=%s", aqi))

		if err != nil {
			log.Fatal(fmt.Errorf("conctl weather - couldn't get weather info from the api server : %w", err))
		}

		if resp.IsSuccess() {
			err = resp.Unmarshal(weatherInfo)
			if err != nil {
				log.Fatal(err)
			}

			var tempType string
			var currTemp float64
			switch temp {
			case "celsius":
				tempType = "\u2103"
				currTemp = weatherInfo.Current.TempC
			case "fahrenheit":
				tempType = "\u2109"
				currTemp = weatherInfo.Current.TempF
			default:
				log.Fatal("Please type the appropriate temperature type, celsius or fahrenheit")
			}

			fmt.Printf("[%s]'s current temperature is : %v%v", weatherInfo.Location.Name, currTemp, tempType)
		} else {
			log.Fatal("bad response:", resp)
		}

	},
}
