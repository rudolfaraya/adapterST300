package tests

import (
	"fmt"
	"io/ioutil" //Solo para obtener tramas crudas desde fichero
	"strings"

	"github.com/gobs/pretty"
	"github.com/rudolfaraya/adapter340/common"
	"github.com/rudolfaraya/adapter340/structs"
)

func ReadFromFile() {

	data, readError := ioutil.ReadFile("tests/RAWFRAMES/RAWFRAME4")
	if readError != nil {
		fmt.Printf("Lectura del archivo fallida\n")
		panic(readError)
	}

	//Separa por línea para obtener cada trama por separado
	rawFrames := strings.Split(string(data), "\r")

	//Slice para almacenar tramas verificadas
	framesVerified := []string{}

	//Slice de tramas verificadas y parseadeas
	framesParsed := []structs.Msg{}
	framesSTTParsed := []structs.STT{}
	framesEMGParsed := []structs.EMG{}
	framesEVTParsed := []structs.EVT{}
	framesALTParsed := []structs.ALT{}

	//Primera Verificación de las tramas: Está vacía?, Es comando ST300?, Es Versión 528?
	for i, frame := range rawFrames {
		if frameVerified, error := common.VerifyFrame(frame); error != nil {
			fmt.Printf("FRAME%d: Verification failed:\n%s", i+1, error)
		} else {
			framesVerified = append(framesVerified, frameVerified)
			fmt.Printf("FRAME%d: Verificación worked\n%s\n", i+1, frameVerified)
		}
		fmt.Println("\n")
	}
	//Segunda verificación de las tramas: El tipo de trama corresponde a su estructura, si es así,
	//se procede a parsear su contenido, almacenándolo en un slice de "msg"
	for i, frameVerified := range framesVerified {
		if tramaDePrueba, error := common.ParserCommon(frameVerified, i+1); error != nil {
			fmt.Printf("FRAME%d: Parser failed:\n%s", i+1, error)
		} else {
			framesParsed = append(framesParsed, tramaDePrueba)
			fmt.Printf("FRAME%d: Parser worked\n", i+1)
		}
		fmt.Println("\n")
	}

	//Imprimir por pantalla el slice de "msg" (tramas verificas y parseadas)
	for i, frameParsed := range framesParsed {
		//Guardar cada tipo trama en su respectiva cola de su tipo e imprimir por pantalla
		_, STTTest, EMGTest, EVTTest, ALTTest, _ := common.Mapper(frameParsed)
		if len(STTTest.Hdr) != 0 {
			pretty.PrettyPrint(STTTest)
			framesSTTParsed = append(framesSTTParsed, STTTest)
		} else if len(EMGTest.Hdr) != 0 {
			pretty.PrettyPrint(EMGTest)
			framesEMGParsed = append(framesEMGParsed, EMGTest)
		} else if len(EVTTest.Hdr) != 0 {
			pretty.PrettyPrint(EVTTest)
			framesEVTParsed = append(framesEVTParsed, EVTTest)
		} else if len(ALTTest.Hdr) != 0 {
			pretty.PrettyPrint(ALTTest)
			fmt.Println(common.MapperAlt(ALTTest.AltID))
			framesALTParsed = append(framesALTParsed, ALTTest)
		} else {
			pretty.PrettyPrint(framesParsed[i])
		}
	}
}
