package iolayer

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/nitinjangam/covid19-vaccine-tracker/models"
)

var (
	addr = os.Getenv("mailServer")
	host = os.Getenv("mailHost")
)

//PrepareAndSend function
func PrepareAndSend(c models.AllCenters, mails []string, email string, pass string) error {
	msg := "***************Vaccination slot available***************\n"
	msgInit := "***************Vaccination slot available***************\n"
	for _, v := range c.Centers {
		msg += fmt.Sprintf("Center_ID: %v\n", v.CenterID)
		msg += fmt.Sprintf("Center_Name: %v\n", v.Name)
		msg += fmt.Sprintf("State: %v\n", v.StateName)
		msg += fmt.Sprintf("District: %v\n", v.DistrictName)
		msg += fmt.Sprintf("Pincode: %v\n", v.Pincode)
		msg += fmt.Sprintf("Fee: %v\n", v.FeeType)
		msg += fmt.Sprintln("Sessions:")
		for _, v1 := range v.Sessions {
			msg += fmt.Sprintf("	Date: %v\n", v1.Date)
			msg += fmt.Sprintf("	Available_Capacity: %v\n", v1.AvailableCapacity)
			msg += fmt.Sprintf("	Min_Age_Limit: %v\n", v1.MinAgeLimit)
			msg += fmt.Sprintln("	Available_Slots:")
			for _, v2 := range v1.Slots {
				msg += fmt.Sprintf("		%v\n", v2)
			}
		}
		msg += fmt.Sprintln("--------------------------------------------------------")
	}
	if msg != msgInit {
		if err := sendMail(msg, mails, email, pass); err != nil {
			return err
		}
	}
	return nil
}

func sendMail(msg string, mails []string, email string, pass string) error {
	log.Printf("sending mail to %v", mails)
	log.Printf("Mail: %v", msg)
	auth := smtp.PlainAuth("", email, pass, host)
	message := "Subject: Vaccination slot available\n\n" + msg
	if err := smtp.SendMail(addr, auth, email, mails, []byte(message)); err != nil {
		return fmt.Errorf("error while sending mail: %v", err)
	}
	return nil
}
