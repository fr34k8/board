package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Validate and parse a card name string into its key and number components.
// For convenience, the key gets converted to uppercase.
func ParseCardName(name string) (string, uint32, error) {

	// Compile regular expression
	re, err := regexp.Compile(`^([a-zA-Z]{1,10})-([0-9]+)$`)
	if err != nil {
		return "", 0, err
	}

	// Match name against regexp
	result := re.FindAllStringSubmatch(name, -1)
	if len(result) != 1 || len(result[0]) != 3 {
		return "", 0, errors.New("Invalid card name")
	}

	// Convert key to uppercase
	key := strings.ToUpper(result[0][1])

	// Parse the number
	number, err := strconv.ParseUint(result[0][2], 10, 32)
	if err != nil {
		return "", 0, err
	}

	return key, uint32(number), nil

}
