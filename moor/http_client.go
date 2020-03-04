package moor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func SanitizeUrl(url string) string {
	if url[0:4] != "http" {
		url = "http://" + url
	}

	return url
}

func BlockerCharactersAmount() int {
	var err error
	var amount int
	amountStr := os.Getenv("MOOR_BLOCKER_CHARACTERS_AMOUNT")
	if len(amountStr) <= 0 {
		amount = BLOCKER_CHARACTERS_AMOUNT
	} else {
		amount, err = strconv.Atoi(amountStr)
		if err != nil {
			fmt.Println("[moor] MOOR_BLOCKER_CHARACTERS_AMOUNT could not be converted to int")
			fmt.Print("[moor]")
			fmt.Print(err)
			fmt.Printf("[moor] Defaulting to BLOCKER_CHARACTERS_AMOUNT=%v\n", BLOCKER_CHARACTERS_AMOUNT)
			amount = BLOCKER_CHARACTERS_AMOUNT
		}
	}
	return amount
}

func Get(url string, client *http.Client) string {
	if client == nil {
		client = &http.Client{}
	}
	url = SanitizeUrl(url)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("[moor] Error during GET")
		fmt.Printf("[moor]")
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)[BlockerCharactersAmount():]
}
