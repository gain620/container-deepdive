package cmd

import (
	"fmt"
	"github.com/gain620/container-deepdive/config"
	"github.com/gain620/container-deepdive/model"
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
		//if len(args) != 1 {
		//	return errors.New("enter the URL")
		//}
		//
		//_, err := url.ParseRequestURI(args[0])
		//if err != nil {
		//	return errors.New("invalid URL")
		//}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Getting the current temperature from the weather API server ... %v \n", args[0:])

		city, err := cmd.Flags().GetString("city")
		if err != nil {
			log.Fatal(err)
		}
		weatherInfo := &model.WeatherInfo{}
		var celsius = "\u2103"
		//var fahrenheit = "\u2109"
		var aqi string = "yes"

		resp, err := req.C().
			SetUserAgent("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1").
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

			fmt.Printf("[%s]'s current temperature is : %v%v", weatherInfo.Location.Name, weatherInfo.Current.TempC, celsius)

		} else {
			log.Fatal("bad response:", resp)
		}

	},
}
