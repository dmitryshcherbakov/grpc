package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/dmitryshcherbakov/grpc/pcbook/proto/pb"
)

/*Указываем rand использовать текущие UNIXNano как начальные значения*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY

	}
}

func RandomCPUBrand() string {
	return RandomStringFromSet("Intel", "AMD")
}

func RandomCPUName(brand string) string {
	if brand == "Intel" {
		return RandomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9998HK",
			"Core i9-9998JF",
			"Core i7-9398LD",
			"Core i3-864Gi",
		)
	}

	return RandomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3800U",
		"Ryzen 3 PRO 3200GE",
	)
}

func RandomGPUBrand() string {
	return RandomStringFromSet("NVIDIA", "AMD")
}

func RandomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return RandomStringFromSet(
			"RTX 2060",
			"RTX 590",
			"GTX 1660-i",
			"GTX 590-i",
		)
	}

	return RandomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 590-iT",
		"RX Vega-56",
	)
}

func RandomLaptopBrand() string {
	return RandomStringFromSet("Apple", "Dell", "Lenovo")
}

func RandomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return RandomStringFromSet("MacBoo Air", "Macbook Pro")
	case "Dell":
		return RandomStringFromSet("Latitude", "Vostro", "XPS", "Alie")
	default:
		return RandomStringFromSet("X12 Lenovo", "P1", "P2 53 Lenovo")
	}
}

func RandomScreenResolution() *pb.Screen_Resolution {
	height := RandomInt(1080, 4324)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}

	return resolution
}

func RandomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}

	return pb.Screen_OLED
}

func RandomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func RandomID() string {
	return uuid.New().String()
}
