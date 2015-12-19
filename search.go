package forecast

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Search(key string, lat string, lon string) (*Forecast, error) {
	resp, err := GetCurrentForecast(key, lat, lon)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var f Forecast
	err = json.Unmarshal(body, &f)

	if err != nil {
		return nil, err
	}

	return &f, nil
}

func GetCurrentForecast(key string, lat string, lon string) (*http.Response, error) {
	coordinate := lat + "," + lon

	url := URL_FORCAST_API + "/" + key + "/" + coordinate

	res, err := http.Get(url)

	if err != nil {
		return res, err
	}
	return res, nil
}
