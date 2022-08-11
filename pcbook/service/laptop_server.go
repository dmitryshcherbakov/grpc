package service

import(
	"github.com/google/uuid"
	"context"
	"fmt"
	"time"
	"log"
	"errors"

	"github.com/dmitryshcherbakov/grpc/pcbook/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{store}
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,

) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("CreateLaptop %v", laptop.Id)

	if len(laptop.Id) > 0 {

		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "InvalidArgument %v", err)	
		}
	}else{
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to generate UUID %v", err)
		}
		laptop.Id = id.String()
	}

		//Блок для теста предположим что у нас долгое выполнение кода на этом участке
		// На стороне клиента получим ошибку с кодом DeadlineExceeded
		// Мы хотим чтобы сервер тоже прекратил свое выполнеие в этом случаи
		time.Sleep(time.Second * 6) 
		fmt.Println(ctx.Err());

		if ctx.Err() == context.Canceled {
			log.Printf("request is canceled")
			return nil, status.Errorf(codes.Canceled, "request is canceled")
		}

		if ctx.Err() == context.DeadlineExceeded {
			log.Printf("deadline is exceeded server")
			//Возвращаем клиенту код ошибки
			return nil, status.Errorf(codes.DeadlineExceeded, "deadline exceeded server")
		}


		//Save loptop to memory
		err := server.Store.Save(laptop)
		if err != nil {
			code := codes.Internal
			if errors.Is(err, ErrAlreadyExists) {
				code = codes.AlreadyExists
			}
			return nil, status.Errorf(code, "Failed to save lo %v", err)
		}

		log.Printf("Saving laptop with id %s", laptop.Id)

		res := &pb.CreateLaptopResponse{
			Id: laptop.Id,
		}
		return res, nil
}
