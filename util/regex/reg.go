package regex

import (
	"regexp"
	"log"
)
// This need thinking, how to handle response from container
func RegReplace(s string) string{
	reg, err := regexp.Compile("[^a-zA-Z0-9_ .,]+")
    if err != nil {
        log.Fatal(err)
    }
	processedString := reg.ReplaceAllString(s, "")
	return processedString
}