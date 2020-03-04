package moor

import "os"

/*
This one may be overwritten with MOOR_BLOCKER_CHARACTERS_AMOUNT environment variable
(see http_client.go#BlockerCharactersAmount() function)
@TODO: add request parameter that would allow override
*/
const BLOCKER_CHARACTERS_AMOUNT = 16

/*
@TODO: Add overwrites
*/
var IGNORE_ENDPOINTS = []string{"favicon.ico", "favicon"}

var AUTH_TOKEN = os.Getenv("TOKEN")
