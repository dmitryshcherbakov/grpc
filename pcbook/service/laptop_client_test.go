package service_test

import (
	"testing"
	"context"
	"net"

	"github.com/stretchr/testify/require"
	"github.com/dmitryshcherbakov/grpc/pcbook/proto/pb"
	"github.com/dmitryshcherbakov/grpc/pcbook/service"
	"github.com/dmitryshcherbakov/grpc/pcbook/sample"
	"github.com/dmitryshcherbakov/grpc/pcbook/serializer"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAdress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAdress)

	laptop := sample.NewLaptop() // Создадим новый образец ноутбука
	expectedID := laptop.Id //Сохраним идентификатор ноутубка в отдельном перменной для сравнения
	req := &pb.CreateLaptopRequest{//Создадим новый объект зароса с ноутбуком
		Laptop: laptop, 
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req) // Используем объект Laptop Client для вызова функции CreatLaptop
	require.NoError(t, err)//Проверяем что не возвращается ошибка
	require.NotNil(t, res) //Проверяем что значение не null
	require.Equal(t, expectedID,res.Id) //Возвращаем идентификатор должен соответствовать ожидаемому

	other, err := laptopServer.Store.Find(res.Id)// Необходимо проверить что ноутбук действительно храниться на сервер, воспользуемся функцией для поиска ноутбука в Laptop_Store
	require.NoError(t, err) //Не должно быть ошибок
	require.NotNil(t, other) //Ноутбук должне быть найден

	//Теперь хотим проверить что сохранили именно тот ноутбук который отправили
	requireSameLaptop(t, laptop, other)

}

func startTestLaptopServer(t *testing.T) ( *service.LaptopServer, string ){
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAdress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAdress, grpc.WithInsecure())
	require.NoError(t, err) // Проверяем что нет ошибок
	return pb.NewLaptopServiceClient(conn) //Возвращаем новый клиента службы с соданным ноутбуком
}

//Для сравнени выполянем сериализацию объектов в JSON
func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSONConvert(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSONConvert(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}











