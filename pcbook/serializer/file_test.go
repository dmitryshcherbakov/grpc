package serializer

import (
	"fmt"
	"testing"
	//"pcbook.pc/sample"
	"pcbook.pc/serializer"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	//laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile("", binaryFile)
	//require.NoError(t, err)

	//err := serializer.WriteProtobufToBinaryFile("STOPPPPPPP!!!!")
	fmt.Println("serializer.WriteProtobufToBinaryFile succeeded	")
}
