package main

import (
	"bufio"
	"encoding/json"
	//"errors"
	"fmt"
	//"io/ioutil" 		//Solo para obtener tramas crudas desde fichero
	//"log"
	"net"
	"os"
	//"reflect"
	//"os/signal"
	//"strconv"
	"strings"
	"time"

	//"github.com/Shopify/sarama"
	"github.com/gobs/pretty"
	"github.com/golang/protobuf/proto"
	"github.com/rudolfaraya/adapter340/common"
	"github.com/rudolfaraya/adapter340/protoinfo"
	"github.com/rudolfaraya/adapter340/structs"
)

func errorEscrituraFichero(e error) {
	if e != nil {
		panic(e)
	}
}

func errorMarshall(e error) error {
	if e != nil {
		return fmt.Errorf("error marshalling FrameInfo to proto: %s", e)
	}
	return nil
}

func main() {

	//Slice de tramas verificadas y parseadeas
	//framesFormated := []protoinfo.FrameInfo{}
	framesParsed := []structs.Msg{}
	framesSTTParsed := []structs.STT{}
	framesEMGParsed := []structs.EMG{}
	framesEVTParsed := []structs.EVT{}
	framesALTParsed := []structs.ALT{}

	//----------------------------------------- INICIO PRUEBA DE PROTO----------------------------------
	fid := common.UniqueID()
	now := time.Now().UTC()

	rawFrame := []byte("ST300STT;205167653;03;528;20170425;20:48:17;741e41;+19.445874;-099.126892;000.000;000.00;0;0;0;12.39;100000;2;0107;000108;0.0;0")

	fi := &structs.FrameInfo{
		FrameID:         fid,
		ServerTimestamp: now,
		Adapter:         "ST340",
		Received:        true,
		Data:            rawFrame,
		DataFormat:      "",
		DevID:           "123456789012345",
	}

	pb := protoinfo.ToProto(fi)
	b, marshalErr := proto.Marshal(pb)
	errorMarshall(marshalErr)

	//fmt.Println(rawFrame)
	//fmt.Println(fi)
	//fmt.Println(pb)
	fmt.Println(b)

	//----------------------------------------- FIN PRUEBA DE PROTO----------------------------------
	/*
		//----------------------------------------INICIO PRUEBA KAFKA------------------------------------
		// Setup configuration
		config := sarama.NewConfig()
		// Return specifies what channels will be populated.
		// If they are set to true, you must read from
		// config.Producer.Return.Successes = true
		// The total number of times to retry sending a message (default 3).
		config.Producer.Retry.Max = 5
		// The level of acknowledgement reliability needed from the broker.
		config.Producer.RequiredAcks = sarama.WaitForAll
		brokers := []string{"localhost:9092"}
		producer, err := sarama.NewAsyncProducer(brokers, config)
		if err != nil {
			// Should not reach here
			panic(err)
		}

		defer func() {
			if err := producer.Close(); err != nil {
				// Should not reach here
				panic(err)
			}
		}()

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)

		var enqueued, errors int
		doneCh := make(chan struct{})

		var contador int64 = 0
		go func() {
			for {

				time.Sleep(500 * time.Millisecond)

				//strTime := strconv.Itoa(int(time.Now().Unix()))
				strContador := strconv.FormatInt(contador, 10)
				msg := &sarama.ProducerMessage{
					Topic: "important",
					Key:   sarama.StringEncoder(strContador),
					Value: sarama.StringEncoder(sarama.ByteEncoder(b)),
				}
				contador++
				select {
				case producer.Input() <- msg:
					enqueued++
					fmt.Println("Produce message", msg.Key, msg.Value)
				case err := <-producer.Errors():
					errors++
					fmt.Println("Failed to produce message:", err)
				case <-signals:
					doneCh <- struct{}{}
				}

			}
		}()

		<-doneCh
		log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
	*/
	//------------------------------------------FIN PRUEBA KAFKA------------------------------------

	//Creción del fichero donde se guardarán las tramas procesadas
	file, err := os.Create("framesProcessed/framesProcessed17")
	errorEscrituraFichero(err)

	defer file.Close()

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// accept connection on port
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err.Error())
		os.Exit(1)
	}

	contador2 := 0
	// Loop infinito en espera de tramas
	for {
		rawFrameTCP, err := bufio.NewReader(conn).ReadString('\r')

		fmt.Println("-----------------------------------------------------")

		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}
		fmt.Println(conn.RemoteAddr())
		rawFrameTCP = strings.TrimSuffix(rawFrameTCP, "\r")

		if frameVerified, error := common.VerifyFrame(rawFrameTCP); error != nil {
			fmt.Printf("\nFRAME%d: Verification failed:\n%s", contador2+1, error)
		} else {
			fmt.Printf("\nFRAME%d: Verification worked\n%s\n", contador2+1, frameVerified)
			if testFrame, error := common.ParserCommon(frameVerified, contador2+1); error != nil {
				fmt.Printf("\nFRAME%d: Parser failed:\n%s", contador2+1, error)
			} else {
				testFrameFormated := common.ToFormatPlatform(testFrame, rawFrameTCP)
				//framesFormated = append(framesFormated, testFrameFormated)
				framesParsed = append(framesParsed, testFrame)
				fmt.Printf("\nFRAME%d: Parser worked\n", contador2+1)
				//b, marshalErr := proto.Marshal(testFrameFormated)
				errorMarshall(marshalErr)

				pretty.PrettyPrint(testFrameFormated)
			}

			//Guardar cada tipo trama en su respectiva cola de su tipo e imprimir por pantalla
			_, STTTest, EMGTest, EVTTest, ALTTest, _ := common.Mapper(framesParsed[contador2])
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
				pretty.PrettyPrint(framesParsed[contador2])
			}

			outJson, err := json.Marshal(framesParsed[contador2])
			if err != nil {
				panic(err)
			}
			escritura, err := file.Write(outJson)
			errorEscrituraFichero(err)
			fmt.Printf("wrote %d bytes\n", escritura)

			file.Sync()

			contador2++
		}
		fmt.Println("-----------------------------------------------------")
		fmt.Println("PRUEBA GITHUB")
	}

}
