package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"github.com/Garuda19/benchmarking/protobuf_files"
	sample "github.com/Garuda19/benchmarking/flatbuf_files/flat_benchmarking"

	"github.com/golang/protobuf/proto"
	"github.com/nasermirzaei89/chance"
	flatbuffers "github.com/google/flatbuffers/go"
)

type parsedDelta struct {
	SessionId            string
	Type                 int32
	Chips                int64
	Tracking             []string
	InstallOS            string
	Timestamp            int64
	ClearTimeStamp       bool
	IsTournament         bool
}

func main() {

	//test := &benchmarking.Test{
	//	Label: "hello",
	//	Type:  17,
	//	Reps:  []int64{1, 2, 3},
	//}
	//data, err := proto.Marshal(test)
	//if err != nil {
	//	log.Fatal("marshaling error: ", err)
	//}
	//newTest := &benchmarking.Test{}
	//err = proto.Unmarshal(data, newTest)
	//if err != nil {
	//	log.Fatal("unmarshaling error: ", err)
	//}
	//// Now test and newTest contain the same data.
	//if test.GetLabel() != newTest.GetLabel() {
	//	log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	//}
	//fmt.Print(test)

	totalTestSize := 100000
	protoArrays := make([][]byte, totalTestSize)
	var  serializationTime int64
	var  deserializationTime int64
	var dataSize float64
	serializationTime = 0
	deserializationTime = 0
	dataSize = 0

	fmt.Print("test size ", totalTestSize, "\n")
	for i := 0; i < totalTestSize; {
		startTime := makeTimestamp()
		data, err := generateDelta()
		if err != nil {
			log.Fatal("marshaling proto error: ", err)
		} else {
			protoArrays[i] = data
			i++
			serializationTime += (makeTimestamp()) - startTime
			dataSize += float64(len(data))/1024
		}
	}

	for i := 0; i < totalTestSize; i++ {
		startTime := makeTimestamp()
		err := parseDelta(protoArrays[i])
		if (err != nil) {
			log.Fatal("unmarshaling proto error: ", err)
		} else {
			deserializationTime += makeTimestamp() - startTime
		}
	}
	fmt.Print("protobuf ------ \n")
	fmt.Print("total serialization time ", serializationTime, "\n")
	fmt.Print("average serialization time ", (float64(serializationTime))/(float64(totalTestSize)), "\n")
	fmt.Print("total deserialization time ", deserializationTime, "\n")
	fmt.Print("average deserialization time ", (float64(deserializationTime))/(float64(totalTestSize)), "\n")
	fmt.Print("total data size ", dataSize, "\n")
	fmt.Print("average data size ", (float64(dataSize))/(float64(totalTestSize)), "\n")


	protoArrays = make([][]byte, totalTestSize)
	deserializationTime = 0
	serializationTime = 0
	dataSize = 0
	for i := 0; i < totalTestSize; {
		startTime := makeTimestamp()
		data, err := generateDeltaJSON()
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		} else {
			protoArrays[i] = data
			i++
			serializationTime += (makeTimestamp()) - startTime
			dataSize += float64(len(data))/1024
		}
	}
	for i := 0; i < totalTestSize; i++ {
		startTime := makeTimestamp()
		err := parseDeltaJSON(protoArrays[i])
		if (err != nil) {
			log.Fatal("unmarshaling json error: ", err)
		} else {
			deserializationTime += makeTimestamp() - startTime
		}
	}
	fmt.Print("JSON ----- \n")
	fmt.Print("total serialization time ", serializationTime, "\n")
	fmt.Print("average serialization time ", (float64(serializationTime))/(float64(totalTestSize)), "\n")
	fmt.Print("total deserialization time ", deserializationTime, "\n")
	fmt.Print("average deserialization time ", (float64(deserializationTime))/(float64(totalTestSize)), "\n")
	fmt.Print("total data size ", dataSize, "\n")
	fmt.Print("average data size ", (float64(dataSize))/(float64(totalTestSize)), "\n")

	protoArrays = make([][]byte, totalTestSize)
	deserializationTime = 0
	serializationTime = 0
	dataSize = 0
	for i := 0; i < totalTestSize; {
		startTime := makeTimestamp()
		data := generateFlatDelta()
		protoArrays[i] = data
		i++
		serializationTime += (makeTimestamp()) - startTime
		dataSize += float64(len(data))/1024
	}
	for i := 0; i < totalTestSize; i++ {
		startTime := makeTimestamp()
		parseFlatDelta(protoArrays[i])
		deserializationTime += makeTimestamp() - startTime
	}
	fmt.Print("FlatBuffer ----- \n")
	fmt.Print("total serialization time ", serializationTime, "\n")
	fmt.Print("average serialization time ", (float64(serializationTime))/(float64(totalTestSize)), "\n")
	fmt.Print("total deserialization time ", deserializationTime, "\n")
	fmt.Print("average deserialization time ", (float64(deserializationTime))/(float64(totalTestSize)), "\n")
	fmt.Print("total data size ", dataSize, "\n")
	fmt.Print("average data size ", (float64(dataSize))/(float64(totalTestSize)), "\n")

}

func generateDelta() ([]byte, error){
	ch:= chance.New(chance.SetSeed(time.Now().UnixNano()))
	delta := &benchmarking.Delta{}
	delta.SessionId = ch.String()
	delta.Type = 3
	delta.Chips = ch.Int64()
	delta.ClearTimeStamp = ch.Bool()
	delta.Tracking = []string{ch.String(), ch.String()}
	delta.InstallOS = "gplay"
	delta.Timestamp = ch.Int64()
	delta.IsTournament = ch.Bool()
	return proto.Marshal(delta)

}

func generateDeltaJSON() ([]byte, error){

	ch:= chance.New(chance.SetSeed(time.Now().UnixNano()))
	delta := &parsedDelta {
		SessionId: ch.String(),
		Type:3,
		Chips:ch.Int64(),
		ClearTimeStamp:ch.Bool(),
		Tracking:[]string{chance.String(), ch.String()},
		InstallOS:"gplay",
		Timestamp:ch.Int64(),
		IsTournament:ch.Bool(),
	}
	return json.Marshal(delta)

}

func generateFlatDelta() []byte{
	ch:= chance.New(chance.SetSeed(time.Now().UnixNano()))
	builder := flatbuffers.NewBuilder(1024)


	track1 := builder.CreateString(chance.String())
	track2 := builder.CreateString(chance.String())
	sample.FlatDeltaStartTrackingVector(builder, 2)
	builder.PrependUOffsetT(track1)
	builder.PrependUOffsetT(track2)
	track := builder.EndVector(2)

	sessId := builder.CreateString(ch.String())
	os := builder.CreateString("gplay")

	sample.FlatDeltaStart(builder)
	sample.FlatDeltaAddSessionId(builder, sessId)
	sample.FlatDeltaAddType(builder, 3)
	sample.FlatDeltaAddChips(builder, ch.Int64())
	sample.FlatDeltaAddClearTimeStamp(builder, ch.Bool())
	sample.FlatDeltaAddTracking(builder, track)
	sample.FlatDeltaAddInstallOS(builder, os)
	sample.FlatDeltaAddTimestamp(builder, ch.Int64())
	sample.FlatDeltaAddIsTournament(builder, ch.Bool())
	flatDelta := sample.FlatDeltaEnd(builder)
	builder.Finish(flatDelta)
	return builder.FinishedBytes()
}

func parseFlatDelta(data []byte) *sample.FlatDelta{
	flatDelta := sample.GetRootAsFlatDelta(data, 0)
	flatDelta.SessionId()
	flatDelta.Type()
	flatDelta.Timestamp()
	flatDelta.Chips()
	flatDelta.ClearTimeStamp()
	flatDelta.Tracking(0)
	flatDelta.Tracking(1)
	flatDelta.InstallOS()
	flatDelta.Timestamp()
	flatDelta.IsTournament()

	return flatDelta
}

func parseDelta(data []byte,) error {
	newDelta := &benchmarking.Delta{}
	return proto.Unmarshal(data, newDelta)
}

func parseDeltaJSON(data []byte) error {
	newDelta := &parsedDelta{}
	return json.Unmarshal(data, newDelta)
}


func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}