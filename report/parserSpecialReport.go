package report

import (
	"fmt"
	"strconv"

	"github.com/rudolfaraya/adapter340/structs"
)

//parserSpecialReport trata el resto de los campos comunes en las tramas de tipo Reporte
func ParserSpecialReport(frameSlice []string, frameParsedCommon structs.Msg) (structs.Msg, error) {

	frameParsedSpecialEmpty := structs.Msg{}

	//VERIFICACION DE CADA CAMPO ESPECIAL
	//Comprobación Model
	ModelParsed, err := strconv.ParseInt(frameSlice[2], 10, 32)
	if err != nil || ModelParsed <= 0 {
		fmt.Println("Hubo un error al parsear Model")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Parseo de Date
	DateParsed, err := strconv.ParseInt(frameSlice[4], 10, 32)
	if err != nil {
		fmt.Println("Hubo un error al parsear Date")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Comprobación Time
	if len(frameSlice[5]) == 0 {
		fmt.Println("Time no ingresado")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Comprobación Cell
	if len(frameSlice[6]) == 0 {
		fmt.Println("Cell no ingresado")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Parseo de Latitude
	LatitudeParsed, err := strconv.ParseFloat(frameSlice[7], 64)
	if err != nil {
		fmt.Println("Hubo un error al parsear Latitude")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Parseo de Longitude
	LongitudeParsed, err := strconv.ParseFloat(frameSlice[8], 64)
	if err != nil {
		fmt.Println("Hubo un error al parsear Longitude")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Parseo de Speed
	SpeedParsed, err := strconv.ParseFloat(frameSlice[9], 64)
	if err != nil {
		fmt.Println("Hubo un error al parsear Speed")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Parseo de Heading
	HeadingParsed, err := strconv.ParseFloat(frameSlice[10], 64)
	if err != nil {
		fmt.Println("Hubo un error al parsear Heading")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Parseo de Satellites
	SatellitesParsed, err := strconv.ParseInt(frameSlice[11], 10, 32)
	if err != nil {
		fmt.Println("Hubo un error al parsear Satellites")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Parseo de GPSFixed
	GPSFixedParsed, err := strconv.ParseBool(frameSlice[12])
	if err != nil {
		fmt.Println("Hubo un error al parsear GPSFixed")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Parseo de Distance
	DistanceParsed, err := strconv.ParseInt(frameSlice[13], 10, 32)
	if err != nil {
		fmt.Println("Hubo un error al parsear Distance")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Parseo de PowerVolt
	PowerVoltParsed, err := strconv.ParseFloat(frameSlice[14], 64)
	if err != nil {
		fmt.Println("Hubo un error al parsear PowerVolt")
		return frameParsedSpecialEmpty, errWrongData
	}

	//Comprobación InOut
	if len(frameSlice[15]) == 0 {
		fmt.Println("InOut no ingresado")
		return frameParsedSpecialEmpty, errWrongData
	}
	//INT FIELDS
	frameParsedCommon.ContentsInt = append(frameParsedCommon.ContentsInt, ModelParsed)
	frameParsedCommon.ContentsInt = append(frameParsedCommon.ContentsInt, DateParsed)
	frameParsedCommon.ContentsInt = append(frameParsedCommon.ContentsInt, SatellitesParsed)
	frameParsedCommon.ContentsInt = append(frameParsedCommon.ContentsInt, DistanceParsed)
	//STRING FIELDS
	frameParsedCommon.ContentsString = append(frameParsedCommon.ContentsString, frameSlice[5])  //Time
	frameParsedCommon.ContentsString = append(frameParsedCommon.ContentsString, frameSlice[6])  //Cell
	frameParsedCommon.ContentsString = append(frameParsedCommon.ContentsString, frameSlice[15]) //InOut
	//FLOAT FIELDS
	frameParsedCommon.ContentsFloat = append(frameParsedCommon.ContentsFloat, LatitudeParsed)
	frameParsedCommon.ContentsFloat = append(frameParsedCommon.ContentsFloat, LongitudeParsed)
	frameParsedCommon.ContentsFloat = append(frameParsedCommon.ContentsFloat, SpeedParsed)
	frameParsedCommon.ContentsFloat = append(frameParsedCommon.ContentsFloat, HeadingParsed)
	frameParsedCommon.ContentsFloat = append(frameParsedCommon.ContentsFloat, PowerVoltParsed)
	//BOOL FIELDS
	frameParsedCommon.ContentsBool = append(frameParsedCommon.ContentsBool, GPSFixedParsed)

	//TRATAR ERRORES
	if frameParsedCommon.Hdr == "ST300STT" {
		frameParsedCommon, err := ParserSTT(frameSlice, frameParsedCommon)
		if err != nil {
			return frameParsedSpecialEmpty, errWrongData
		}
		return frameParsedCommon, nil
	} else if frameParsedCommon.Hdr == "ST300EMG" || frameParsedCommon.Hdr == "ST300EVT" || frameParsedCommon.Hdr == "ST300ALT" {
		frameParsedCommon, err := ParserEmgEvtAlt(frameSlice, frameParsedCommon)
		if err != nil {
			return frameParsedSpecialEmpty, errWrongData
		}
		return frameParsedCommon, nil
	}

	return frameParsedCommon, nil
}
