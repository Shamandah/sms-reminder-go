package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	reminders := []string{
		"Reminder: Please take your insulin at 8:00 AM.",
		"Reminder: Please take your insulin at 2:00 PM.",
		"Reminder: Please take your insulin at 8:00 PM.",
	}

	for _, reminder := range reminders {
		params := &openapi.CreateMessageParams{}
		params.SetTo("+254790504058")  // Patient number
		params.SetFrom("+17726778710") // Twilio number
		params.SetBody(reminder)

		resp, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println("Error sending SMS:", err)
		} else {
			fmt.Println("Message sent successfully! SID:", *resp.Sid)
		}
	}
}
