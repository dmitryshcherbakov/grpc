package main

import (
	"flag"
	"fmt"
	"github.com/dmitryshcherbakov/grpc/pcbook/proto/pb"
	"github.com/dmitryshcherbakov/grpc/pcbook/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 0, "the port to listen on") //Используем флаг для получения номера порта из командной строки
	flag.Parse() // Распарсим флаг
	log.Printf("Starting serverAdress port %d", *port) //Печатаем простой лог

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore()) //Создаем новый объект сервера ноутбука
	grpcServer := grpc.NewServer() //Создадим сервер
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer) //Зарегистрируем сервер ноутбуков на сервер GRPC

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address) // Прслушиваем tcp соединение на полученном порту
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}