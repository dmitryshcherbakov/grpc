package main

import (
	"time"
	"context"
	"flag"
	"github.com/dmitryshcherbakov/grpc/pcbook/proto/pb"
	"github.com/dmitryshcherbakov/grpc/pcbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	// This example demonstrates the use of a Go program that can
	serverAddress := flag.String("address", "", "Addresses to connect to the server") //Получаем адрес сервера из командной строки
	flag.Parse()
	log.Printf("Dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure()) //Создаем небезопасно соединение с сервером
	if err != nil {
		log.Fatal("dials serverAddress not connections: ", err) //Пишем лог если ошибка
	}

	laptopClient := pb.NewLaptopServiceClient(conn)//Создаем новый клиентский объект ноутбка с соединением

	laptop := sample.NewLaptop()  //Создаем новый ноутбук

	laptop.Id = ""

	req := &pb.CreateLaptopRequest{ //Создаем новый объект запроса
		Laptop: laptop,
	}
	// Установим таймаут для выполнения при создании ноутбука
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req) //Вызываем функцию создания ноутбука, передаем фоновый контекст и запрос
	if err != nil {
		st, ok := status.FromError(err) // Преобразуем ошибку в объект состояния для проверки возвращенного кода состояния
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal (" cannot create new laptop:", err)
		}
		return
	}

	log.Printf("creating new laptop with id %d", res.Id)
}