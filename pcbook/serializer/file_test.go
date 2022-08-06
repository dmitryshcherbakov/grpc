package serializer_test
import (
	"fmt"
	"testing"
	//"pcbook.pc/sample"
	//"pcbook.pc/serializer"
	"github.com/dmitryshcherbakov/grpc/pcbook/proto/pb"
	"github.com/dmitryshcherbakov/grpc/pcbook/sample"
	"github.com/dmitryshcherbakov/grpc/pcbook/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)
func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	//err := serializer.WriteProtobufToBinaryFile("STOPPPPPPP!!!!")
	fmt.Println("serializer.WriteProtobufToBinaryFile succeeded	")

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
