package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"strings"
)

func parseQuote(quoteMessage *slack.MessageEvent) string {

	var quote QuoteRequest
	if len(quoteMessage.Attachments) == 1 {
		attachment := quoteMessage.Attachments[0]
		for _, v := range attachment.Fields {
			switch v.Title {
			case "Services you are interested in":
				quote.Services = v.Value
			case "Street Address":
				quote.Address = v.Value
			case "Apartment, suite, etc":
				quote.Unit = v.Value
			case "City":
				quote.City = v.Value
			case "State/Province":
				quote.State = v.Value
			case "ZIP / Postal Code":
				quote.Zipcode = v.Value
			case "Prefix":
				quote.Prefix = v.Value
			case "First Name":
				quote.FirstName = v.Value
			case "Last Name":
				quote.LastName = v.Value
			case "Email Address":
				email := strings.Split(v.Value, "|")[1]
				email = strings.ReplaceAll(email, ">", "")
				quote.Email = email
			case "Phone Number":
				quote.Phone = v.Value
			case "Additional notes":
				quote.Notes = v.Value
			}
		}
	}
	message := fmt.Sprintf("Hey Ben, this quote just came in: %s. %s %s - %s %s, %s %s - %s %s - Services Interested in: %s Notes: %s", quote.Prefix, quote.FirstName, quote.LastName, quote.Address, quote.City, quote.State, quote.Zipcode, quote.Email, quote.Phone, quote.Services, quote.Notes)
	return message
}
