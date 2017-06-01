package report

import (
	"strings"
	"time"

	"github.com/rudolfaraya/adapter340/parseData"
	"github.com/rudolfaraya/adapter340/structs"
)

//ParserALV
func ParserALV(frame string, frameID int) (structs.Msg, error) {

	frameParsed := structs.Msg{}

	//Separar los datos de cada campo dentro de un slice
	frameSlice := strings.Split(frame, ";")

	//VERIFICACION DE CADA CAMPO COMÚN
	//Parseo de DevID
	DevIDParsed, err := parseData.ParseDevID(frameSlice[1])
	if err != nil {
		return frameParsed, errWrongData
	}

	frameParsedALV := structs.Msg{
		ServerTimestamp: time.Now(),
		FrameID:         frameID,
		Hdr:             frameSlice[0],
		DevID:           DevIDParsed}

	return frameParsedALV, nil

}
