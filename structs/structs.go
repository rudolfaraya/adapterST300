package structs

import (
	"errors"
	"time"
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

type Msg struct {
	ServerTimestamp time.Time //Server Time Stamp
	FrameID         int       //Identificador Único
	Hdr             string    //[0] Comando (Común)
	DevID           int64     //[1] Identificador Dispositivo (Común)
	Version         int64     //[2] Versión (Común)
	ContentsInt     []int64   //[3] Contenidos de la trama de tipo entero (N° de campos variables)
	ContentsFloat   []float64 //[4] Contenidos de la trama de tipo flotante (N° de campos variables)
	ContentsString  []string  //[5] Contenidos de la trama de tipo string (N° de campos variables)
	ContentsBool    []bool    //[6] Contenidos de la trama de tipo booleano (N° de campos variables)
}

type STT struct {
	ServerTimestamp time.Time //Server Time Stamp
	FrameID         int       //Identificador Único
	Hdr             string    //[0]
	DevID           int64     //[1]
	Model           int64     //[2]
	Version         int64     //[3]
	Date            int64     //[4]
	Time            string    //[5]
	Cell            string    //[6]
	Latitude        float64   //[7]
	Longitude       float64   //[8]
	Speed           float64   //[9]
	Heading         float64   //[10]
	Satellites      int64     //[11]
	GPSFixed        bool      //[12]
	Distance        int64     //[13]
	PowerVolt       float64   //[14]
	InOut           string    //[15]
	Mode            int64     //[16]-STT
	MsgNum          int64     //[17]-STT
	HourMeter       int64     //[18]-STT
	BackupVolt      float64   //[19]-STT
	MsgType         bool      //[20]-STT
}

type EMG struct {
	ServerTimestamp time.Time //Server Time Stamp
	FrameID         int       //Identificador Único
	Hdr             string    //[0]
	DevID           int64     //[1]
	Model           int64     //[2]
	Version         int64     //[3]
	Date            int64     //[4]
	Time            string    //[5]
	Cell            string    //[6]
	Latitude        float64   //[7]
	Longitude       float64   //[8]
	Speed           float64   //[9]
	Heading         float64   //[10]
	Satellites      int64     //[11]
	GPSFixed        bool      //[12]
	Distance        int64     //[13]
	PowerVolt       float64   //[14]
	InOut           string    //[15]
	EmgID           int64     //[16]
	HourMeter       int64     //[17]
	BackupVolt      float64   //[18]
	MsgType         bool      //[19]
}

type EVT struct {
	ServerTimestamp time.Time //Server Time Stamp
	FrameID         int       //Identificador Único
	Hdr             string    //[0]
	DevID           int64     //[1]
	Model           int64     //[2]
	Version         int64     //[3]
	Date            int64     //[4]
	Time            string    //[5]
	Cell            string    //[6]
	Latitude        float64   //[7]
	Longitude       float64   //[8]
	Speed           float64   //[9]
	Heading         float64   //[10]
	Satellites      int64     //[11]
	GPSFixed        bool      //[12]
	Distance        int64     //[13]
	PowerVolt       float64   //[14]
	InOut           string    //[15]
	EvtID           int64     //[16]
	HourMeter       int64     //[17]
	BackupVolt      float64   //[18]
	MsgType         bool      //[19]
}

type ALT struct {
	ServerTimestamp time.Time //Server Time Stamp
	FrameID         int       //Identificador Único
	Hdr             string    //[0]
	DevID           int64     //[1]
	Model           int64     //[2]
	Version         int64     //[3]
	Date            int64     //[4]
	Time            string    //[5]
	Cell            string    //[6]
	Latitude        float64   //[7]
	Longitude       float64   //[8]
	Speed           float64   //[9]
	Heading         float64   //[10]
	Satellites      int64     //[11]
	GPSFixed        bool      //[12]
	Distance        int64     //[13]
	PowerVolt       float64   //[14]
	InOut           string    //[15]
	AltID           int64     //[16]
	HourMeter       int64     //[17]
	BackupVolt      float64   //[18]
	MsgType         bool      //[19]
}
