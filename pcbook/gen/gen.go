package main

import (
	"fmt"

	//"google.golang.org/genproto/googleapis/maps/playablelocations/v3/sample"

	"github.com/golang/protobuf/ptypes"
	//"pcbook.pc/proto/pb"
	//"github.com/dmitryshcherbakov/grpc/tree/main/pcbook/proto/pb"
	//"github.com/dmitryshcherbakov/grpc/tree/main/pcbook/proto/pb"
	"github.com/dmitryshcherbakov/grpc/tree/main/pcbook/proto/pb"
	"gitlab.com/techschool/pcbook/sample"
	"pcbook.pc/sample"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  sample.RandomKeyboardLayout(),
		Backlit: sample.RandomBool(),
	}

	return keyboard
}

func NewCPU() *pb.CPU {
	brand := sample.RandomCPUBrand()
	name := sample.RandomCPUName(brand)

	numberCores := sample.RandomInt(2, 8)
	numberThreads := sample.RandomInt(numberCores, 12)

	minGhz := sample.RandomFloat64(2.0, 3.5)
	maxGhz := sample.RandomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

func NewGPU() *pb.GPU {
	brand := sample.RandomGPUBrand()
	name := sample.RandomGPUName(brand)

	minGhz := sample.RandomFloat64(1.0, 1.5)
	maxGhz := sample.RandomFloat64(minGhz, 2.0)

	memory := &pb.Memory{
		Value: uint64(sample.RandomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return gpu
}

func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(sample.RandomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return ram
}

func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(sample.RandomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(sample.RandomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   sample.RandomFloat32(13, 17),
		Resolution: sample.RandomScreenResolution(),
		Panel:      sample.RandomScreenPanel(),
		Multitouch: sample.RandomBool(),
	}

	return screen
}

func NewLaptop() *pb.Laptop {
	brand := sample.RandomLaptopBrand()
	name := sample.RandomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       sample.RandomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storages{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: sample.RandomFloat64(1.0, 3.0),
		},
		PriceUsd:    sample.RandomFloat64(1500, 3000),
		ReleaseYear: uint32(sample.RandomInt(2015, 2022)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
}

func main() {
	fmt.Println("GOGOGOGOGOGO!")

	list := NewKeyboard()

	fmt.Println("Layout ", list.Layout)
	fmt.Println("Backlit ", list.Backlit)

}
