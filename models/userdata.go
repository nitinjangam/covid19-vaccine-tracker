package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Entries structure to get all formdata
type Entries struct {
	Details []struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		PinCode   string `json:"pincode"`
		GetUpdate string `json:"getupdate"`
	} `json:"details"`
}

//FetchData function to fetch data from Forms. In our case it's from JSON file for now...
func FetchData() (map[string][]string, error) {
	resp := Entries{}
	pinMap := make(map[string][]string)
	jsonFile, err := os.Open("formdata.json")
	if err != nil {
		return map[string][]string{}, err
	}
	bytes, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(bytes, &resp); err != nil {
		return map[string][]string{}, err
	}
	for _, v := range resp.Details {
		if v.GetUpdate == "yes" {
			pinMap[v.PinCode] = append(pinMap[v.PinCode], v.Email)
		}
	}
	return pinMap, nil
}
