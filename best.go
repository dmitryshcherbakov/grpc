package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// to consume messages
	topic := "quickstart"
	//partition := 0
	connect := "broker:29092"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{connect},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		MaxWait:   100 * time.Millisecond,
	})
	defer reader.Close()
	fmt.Println("go func 1")

	go func() {
		fmt.Println("go func 2")
		for {
			m, err := reader.ReadMessage(context.Background())

			if err != nil {
				fmt.Println("error while receiving message: %s", err.Error())
				continue
			}

			fmt.Println("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))
		}
		//return
	}()

	fmt.Scanln() // ждем ввода пользователя
	fmt.Println("The End")

	//msg := <-messages
	//fmt.Println(msg)

	/*for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}*/

	/*if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}*/
}
