// 21. Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

// Пример: адаптируем велосипед под интерфейс
// машины (почему бы и нет?)

// Для простых случаев не нужно создавать дополнительную
// структуру, оборачивающую адаптируемую структуру,
// достаточно реализовать для нее методы интерфейса.

type Car interface {
	Go()
	Speed() float64 // km/h
	RegNumber() string
}

type Bycicle struct {
	Model       string
	WheelRPM    float64
	WheelRaduis float64
}

// Предположим, что прежде чем использовать велосипед
// в качестве машины, его надо зарегистировать.
type RegisteredBicycle struct {
	bicycle   *Bycicle
	regNumber string
}

// констуктор для регистрации велосипеда
func NewRegisteredBicycle(bycicle *Bycicle, regNumber string) *RegisteredBicycle {
	return &RegisteredBicycle{
		bicycle:   bycicle,
		regNumber: regNumber,
	}
}

// Реализуем методы интерфейса Car для RegisteredBicycle
func (b *RegisteredBicycle) Go() {
	fmt.Println("Bicycle going")
}

func (b *RegisteredBicycle) Speed() float64 {
	speed := b.bicycle.WheelRPM    // rots/min
	speed *= b.bicycle.WheelRaduis // m/min
	return speed * 1000 / 60       // km/h
}

func (b *RegisteredBicycle) RegNumber() string {
	return b.regNumber
}

func main() {
	// NINETY ONE Leopard 27.5T 21 Speed Hybrid Bike
	bicycle := &Bycicle{
		Model:       "Leopard",
		WheelRPM:    45,
		WheelRaduis: 0.4,
	}

	regNumber := "6921TF"
	regBike := NewRegisteredBicycle(bicycle, regNumber)
	doSomethingWithCar(regBike)
}

func doSomethingWithCar(car Car) {
	fmt.Println("Car reg number:", car.RegNumber())
	fmt.Println("Car speed:", car.Speed(), "km/h")
	fmt.Println("Car now goes away!")

	car.Go()
}

// Вывод:
// Car reg number: 6921TF
// Car speed: 300 km/h
// Car now goes away!
// Bicycle going
