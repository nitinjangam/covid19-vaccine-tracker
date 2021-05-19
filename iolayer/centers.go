package iolayer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/nitinjangam/covid19-vaccine-tracker/models"
)

const (
	baseURL                = "https://cdn-api.co-vin.in/api"
	calendarByPinURLFormat = "/v2/appointment/sessions/calendarByPin?pincode=%s&date=%s"
)

var (
	errprt = color.New(color.FgRed).Add(color.Bold).Add(color.Italic)
	info   = color.New(color.FgGreen).Add(color.Bold)
)

func today() string {
	return time.Now().Format("02-01-2006")
}

//GetCalendarByPin function to fetch sessions based on pincode
func GetCalendarByPin(pin string) (models.AllCenters, error) {
	url := fmt.Sprintf(baseURL+calendarByPinURLFormat, pin, today())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.AllCenters{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "hi_IN")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36 Edg/90.0.818.51")

	info.Println(fmt.Sprintf("Querying endpoint: %s", url))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.AllCenters{}, err
	}
	defer resp.Body.Close()

	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.AllCenters{}, err
	}
	//log.Print("Response: ", string(bdy))

	if resp.StatusCode != http.StatusOK {
		// Sometimes the API returns "Unauthenticated access!", do not fail in that case
		if resp.StatusCode == http.StatusUnauthorized {
			return models.AllCenters{}, nil
		}
		return models.AllCenters{}, fmt.Errorf("Request failed with statusCode: %d", resp.StatusCode)
	}
	res, err := jsonToStruct(bdy)
	if err != nil {
		return models.AllCenters{}, err
	}
	return res, nil
}

func jsonToStruct(b []byte) (models.AllCenters, error) {
	res := models.AllCenters{}
	if err := json.Unmarshal(b, &res); err != nil {
		return models.AllCenters{}, fmt.Errorf("Error while Unmarshal: %v", err)
	}
	return res, nil
}
