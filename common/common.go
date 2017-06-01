package common

import (
	"errors"
	"strings"
	"time"

	"github.com/rudolfaraya/adapter340/parseData"
	"github.com/rudolfaraya/adapter340/report"
	"github.com/rudolfaraya/adapter340/structs"
	"github.com/twinj/uuid"
)

var (
	errFrameSpecial = errors.New("Error-Parser: Frame with special orden of fields")
	errWrongData    = errors.New("Error-Parser: Data can't be parsed")
	errEmpty        = errors.New("Error-Verifier: Empty Frame")
	errCommand      = errors.New("Error-Verifier: Can't verify if the frame has a valid command")
	errVersion      = errors.New("Error-Verifier: Can't verify if frame version is 528")
)

// UniqueID returns a random generated id. The id is a string of length 32, corresponding to an UUID v4.
func UniqueID() string {
	return uuid.Formatter(uuid.NewV4(), uuid.Clean)
}

//parserCommon se encarga del parseo de los campos DevID y Version
//Guarda los datos Hdr, DevID y Version en el tipo de dato msg (trama parseada)
func ParserCommon(frame string, frameID int) (structs.Msg, error) {

	//frameEmpty := msg{}
	frameParsedCommon := structs.Msg{}

	if strings.Index(frame, "ST300LTM") != -1 || strings.Index(frame, "ST300STT") != -1 || strings.Index(frame, "ST300EMG") != -1 || strings.Index(frame, "ST300EVT") != -1 || strings.Index(frame, "ST300ALT") != -1 || strings.Index(frame, "ST300ALV") != -1 || strings.Index(frame, "ST300UEX") != -1 {
		//CÓDIGO PARA TRATAR TRAMAS CON ORDEN DISTINTO DE LOS CAMPOS Y QUE POSEEN EL CAMPO "MODEL" DEL DISPOSITIVO
		if strings.Index(frame, "ST300ALV") != -1 {
			//CÓDIGO PARA TRATA KEEPALIVE(NO POSEE VERSION,NI MODELO)
			frameParsedCommon, err := report.ParserALV(frame, frameID)
			if err != nil {
				return frameParsedCommon, errFrameSpecial
			}
			return frameParsedCommon, nil
		}
		frameParsedCommon, err := report.ParserCommonReport(frame, frameID)
		if err != nil {
			return frameParsedCommon, errWrongData
		}
		return frameParsedCommon, nil
	} else {
		//Separar los datos de cada campo dentro de un slice
		frameSlice := strings.Split(frame, ";")

		//VERIFICACION DE CADA CAMPO COMÚN
		//Parseo de DevID
		DevIDParsed, err := parseData.ParseDevID(frameSlice[2])
		if err != nil {
			return frameParsedCommon, errWrongData
		}
		//Parseo de Version
		VersionParsed, err := parseData.ParseVersion(frameSlice[3])
		if err != nil {
			return frameParsedCommon, errWrongData
		}

		frameParsedCommon := structs.Msg{
			ServerTimestamp: time.Now(),
			FrameID:         frameID,
			Hdr:             frameSlice[0],
			DevID:           DevIDParsed,
			Version:         VersionParsed}

		sliceData := frameSlice[4:]
		Data := strings.Join(sliceData, ";")
		frameParsedCommon.ContentsString = append(frameParsedCommon.ContentsString, Data)

		return frameParsedCommon, nil
	}
	return frameParsedCommon, errWrongData
}

//verifyFrame verifica si la trama está vacía, si posee comando válido y si ĺa versión correcta
func VerifyFrame(frame string) (string, error) {

	//Verificar si la trama está vacía
	if frame == "" {
		return frame, errEmpty
	}
	//Verificar si la trama posee un comando válido
	if strings.Index(frame, "ST300") == -1 {
		return frame, errCommand
	}
	//Verificar si la versión de la trama es 528
	//El comando ALV no muestra la versión en su trama (KeepAlive)
	if strings.Index(frame, "528") == -1 && strings.Index(frame, "ST300ALV") == -1 {
		return frame, errVersion
	}
	//Si no existen errores en la primera verificación, se retorna la trama cruda
	return frame, nil
}
