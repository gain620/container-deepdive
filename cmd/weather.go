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
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Getting weather from the API server ... %v \n", args[0:])

		//var cityArg string
		var weatherInfo model.WeatherInfo
		var celsius = "\u2103"
		//var fahrenheit = "\u2109"
		//conctl weather --aqi --city=seoul

		resp, err := req.C().
			SetUserAgent("Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1").
			SetTimeout(5*time.Second).R().
			SetHeader("Accept", "application/vnd.github.v3+json").
			SetPathParams(map[string]string{
				"api_key": config.WeatherAPIKey,
				"city":    "seoul",
			}).
			Get("https://api.weatherapi.com/v1/current.json?key={api_key}&q={city}&aqi=yes")

		if err != nil {
			log.Warn(fmt.Errorf("conctl weather - couldn't get weather info from the api server : %w", err))
		}

		if resp.IsSuccess() {
			err = resp.Unmarshal(weatherInfo)
			if err != nil {
				log.Warn(err)
			}

			fmt.Printf("[%s]'s current temperature is : %v%v", weatherInfo.Location.Name, weatherInfo.Current.TempC, celsius)

		} else {
			log.Warn("bad response:", resp)
		}

	},
}
