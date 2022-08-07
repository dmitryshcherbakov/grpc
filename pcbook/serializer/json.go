package serializer


import ( 
	//"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
)
//import (
	//"github.com/golang/protobuf/jsonpb"
//	"github.com/golang/protobuf/proto"
//	"google.golang.org/protobuf/encoding/protojson"
//)

/*func ProtobufToJSONConvert(message proto.Message) (string, error) {
	/*marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent: " ",
		OrigName: true,
	}	
	return marshaler.MarshalToString(message)*/

/*	marshaler := protojson.MarshalOptions{
        Indent:          "  ",
        UseProtoNames:   true,
        EmitUnpopulated: true,
    }

	data, err := marshaler.Marshal(message)

    return string(data), err

	//return marshaler.Marshal(message)
}*/

func ProtobufToJSONConvert(message proto.Message) (string, error) {
	//data, err := ProtobufToJSONConvertMessage(message)
	/*data, err := JsonPlus(message)
	
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %v", err)
	}

	err = ioutil.WriteFile(filename, []byte(string(data)), 0644)

	if err != nil {
		return fmt.Errorf("cannot write to JSON data to file: %v", err)
	}
	return nil*/

	marshaler := protojson.MarshalOptions{
        Indent:          "  ",
        UseProtoNames:   true,
        EmitUnpopulated: true,
    }

	data, err := marshaler.Marshal(message)
    return string(data), err

	/*err = ioutil.WriteFile(filename, []byte(string(data)), 0644)
	if err != nil {
		return fmt.Errorf("cannot write to JSON data to file: %v", err)
	}

	return nil*/
}