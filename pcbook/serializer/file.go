package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

//Записываем протобаф файл на диск предварительно преобразовав его в JSON
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSONConvert(message)

	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %v", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write to JSON data to file: %v", err)
	}

	return nil
}

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write message: %v", err)
	}

	//fmt.Println("WriteProtobufToBinaryFile succeeded	")
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("failed to unmarshal message: %v", err)
	}

	return nil
}
