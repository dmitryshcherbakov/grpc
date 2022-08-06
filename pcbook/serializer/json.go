package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	/*marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent: " ",
		OrigName: true,
	}	

	return marshaler.MarshalToString(message)*/

	marshaler := protojson.MarshalOptions{
        Indent:          "  ",
        UseProtoNames:   true,
        EmitUnpopulated: true,
    }

	return marshaler.Marshal(message)
}