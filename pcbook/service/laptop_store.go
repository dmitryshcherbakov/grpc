package service

import(
	//"context"
	"errors"
	"fmt"
	//"log"
	"sync"


	"github.com/dmitryshcherbakov/grpc/pcbook/proto/pb"
	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("already exists record")

type LaptopStore interface {
	Save (laptop *pb.Laptop) error
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data map[string]*pb.Laptop

}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	//deep copier

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("copier copy laptop error: %v", err)
	}

	store.data[other.Id] = other
	return nil


}