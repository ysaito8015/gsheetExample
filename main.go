package main

import (
	"context"
	//"errors"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"

	"google.golang.org/api/sheets/v4"
)

var spreadsheetID = os.Getenv("SPREADSHEET_ID")

func main() {
	credential := option.WithCredentialsFile("credentials/secret.json")

	srv, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		log.Fatal(err)
	}

	readRange := "5!A1:C3"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalln(err)
	}
	if len(resp.Values) == 0 {
		log.Fatalln("data not found")
	}
	for _, row := range resp.Values {
		fmt.Printf("%s, %s\n", row[0], row[1])
	}

	writeRange := "5!A4:C5"
	vr := &sheets.ValueRange{
		Values: [][]interface{}{
			{"xxx", "ttt", "yyy"},
			{"222", 111, "333"},
		},
	}
	_, err = srv.Spreadsheets.Values.Update(spreadsheetID, writeRange, vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalln(err)
	}
}
