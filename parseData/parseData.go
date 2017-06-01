package parseData

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	errVersion = errors.New("Error-Parser: Version can't be parsed")
	errDevID   = errors.New("Error-Parser: DevID can't be parsed")
)

func ParseVersion(version string) (int64, error) {

	versionParsed, err := strconv.ParseInt(version, 10, 32)
	if err != nil || versionParsed <= 0 {
		fmt.Println("Hubo un error al parsear Version")
		return versionParsed, errVersion
	}
	return versionParsed, nil
}

func ParseDevID(devID string) (int64, error) {

	devIDParsed, err := strconv.ParseInt(devID, 10, 32)
	if err != nil || devIDParsed <= 0 {
		fmt.Println("Hubo un error al parsear DevID")
		return devIDParsed, errDevID
	}
	return devIDParsed, nil
}
