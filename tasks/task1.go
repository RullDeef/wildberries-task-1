// 1. Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

package main

import (
	"fmt"
	"time"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
}

func (h Human) FullName() string {
	return fmt.Sprintf("%s %s", h.FirstName, h.LastName)
}

func (h Human) IsMature() bool {
	return h.Age >= 18
}

type Action struct {
	Name string
	Human
	StartedAt time.Time
	Duration  time.Duration
}

func (a Action) IsActiveAt(time time.Time) bool {
	return !a.StartedAt.After(time) &&
		a.StartedAt.Add(a.Duration).After(time)
}

func main() {
	a := Action{
		Name: "work",
		Human: Human{
			FirstName: "John",
			LastName:  "Doe",
			Age:       20,
			Gender:    "male",
		},
		StartedAt: time.Date(2023, 7, 24, 8, 0, 0, 0, time.Local),
		Duration:  time.Hour * 9,
	}

	// методы FullName и IsMature встроены в Action из Human
	fmt.Println("full name -", a.FullName())
	fmt.Println("is mature -", a.IsMature())

	// также можно обращаться через имя встроенной структуры
	if a.Human.Gender == "male" {
		fmt.Print("he ")
	} else {
		fmt.Print("she ")
	}

	fmt.Println("worked at 13:30 -",
		a.IsActiveAt(time.Date(2023, 7, 24, 13, 30, 0, 0, time.Local)))
}
