package report

import (
	"errors"
	"strings"
	"time"

	"github.com/rudolfaraya/adapter340/parseData"
	"github.com/rudolfaraya/adapter340/structs"
)

var (
	errUnderFields      = errors.New("Error: More fields are needed")
	errOverNumberFields = errors.New("Error: Too much fields")
	errWrongData        = errors.New("Error-Parser: Data can't be parsed")
	errFrameSpecial     = errors.New("Error-Parser: Frame with special orden of fields")
	errEmpty            = errors.New("Error-Verifier: Empty Frame")
	errCommand          = errors.New("Error-Verifier: Can't verify if the frame has a valid command")
	errVersion          = errors.New("Error-Verifier: Can't verify if frame version is 528")
)

//parserCommon Report se encarga de parsear los campos comunes de las tramas
//de tipo Reporte (STT,EMG,EVT,ALT). Este tipo de trama no posee el mismo orden
//en los campos siguientes a los comunes, porque depende del modelo del dispositivo
func ParserCommonReport(frame string, frameID int) (structs.Msg, error) {

	frameParsedCommonEmpty := structs.Msg{}

	//Verificar la cantidad correcta de ";"
	semiColonNumber := strings.Count(frame, ";")
	if semiColonNumber > 20 {
		return frameParsedCommonEmpty, errOverNumberFields
	} else if semiColonNumber < 19 {
		return frameParsedCommonEmpty, errUnderFields
	}
	//Separar los datos de cada campo dentro de un slice
	frameSlice := strings.Split(frame, ";")

	//VERIFICACION DE CADA CAMPO COMÚN
	//Parseo de DevID
	DevIDParsed, err := parseData.ParseDevID(frameSlice[1])
	if err != nil {
		return frameParsedCommonEmpty, errWrongData
	}
	//Parseo de Version
	VersionParsed, err := parseData.ParseVersion(frameSlice[3])
	if err != nil {
		return frameParsedCommonEmpty, errWrongData
	}

	frameParsedCommon := structs.Msg{
		ServerTimestamp: time.Now(),
		FrameID:         frameID,
		Hdr:             frameSlice[0],
		DevID:           DevIDParsed,
		Version:         VersionParsed}

	frameParsedCommonAndSpecial, err := ParserSpecialReport(frameSlice, frameParsedCommon)
	if err != nil {
		return frameParsedCommonEmpty, errWrongData
	}
	return frameParsedCommonAndSpecial, nil

}
