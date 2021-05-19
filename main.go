package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/nitinjangam/covid19-vaccine-tracker/iolayer"
	"github.com/nitinjangam/covid19-vaccine-tracker/models"
)

var (
	errprt      = color.New(color.FgRed).Add(color.Bold).Add(color.Italic)
	info        = color.New(color.FgGreen).Add(color.Bold)
	email, pass string
)

func main() {
	//Fetch data from sheet
	allData, err := models.FetchData()
	if err != nil {
		errprt.Printf("error from FetchData: %v", err)
	}
	centers := models.AllCenters{}
	for k, v := range allData {
		errprt.Println(v)
		centers, err = iolayer.GetCalendarByPin(k)
		if err != nil {
			errprt.Printf("error from GetCalendarByPin: %v", err)
		}
		available := centers.ValidateCenters()
		if err := iolayer.PrepareAndSend(available, v, email, pass); err != nil {
			errprt.Printf("error from PrepareAndSend: %v", err)
		}
	}
}

func init() {
	email = os.Getenv("emailID")
	pass = os.Getenv("emailPassword")
}
