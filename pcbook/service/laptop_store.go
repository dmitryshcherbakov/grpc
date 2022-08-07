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
	// Интерфейс для поиск ноутбука принимает id ноутбука возвращает уже объект
	Find(id string) (*pb.Laptop, error)
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

//Ищем ноутбук и возрвращаем обхект с ошибкой
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock() //Вызывам Мутекс для получаения блокировки чтения
	defer store.mutex.RUnlock() //Отключаем блокировка по завершению выполнянеия функции
	
	laptop := store.data[id]//Теперь ищшем ноутбук в карте store.data
	if laptop == nil { //Если ничего не найдено просто возвращаем nil
		return nil, nil
	}
	
	other := &pb.Laptop{}//Если ноутбук найден выполянем глубокое копирования объекта
	err := copier.Copy(other, laptop) // Копируем сам объект
	if err != nil { // Если возникла ошика возвращем ее
		return nil, fmt.Errorf("error creating temporary laptop data: %v", err)
	}

	return other, nil

}