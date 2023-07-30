// 10. Дана последовательность температурных колебаний: -25.4, -27.0
// 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в
// группы с шагом в 10 градусов. Последовательность в подмножноствах
// не важна.

// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5},
// 20: {24.5}, etc.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TemperatureSet map[float32]struct{}

func main() {
	temperatures := []float32{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5,
	}

	groups := ConvertToGroups(temperatures)
	PrintlnGroups(groups)
}

func ConvertToGroups(temperatures []float32) map[int]TemperatureSet {
	groups := make(map[int]TemperatureSet)

	for _, t := range temperatures {
		groupKey := int(t/10) * 10
		if set, ok := groups[groupKey]; ok {
			set.Add(t)
		} else {
			set = NewTemperatureSet()
			set.Add(t)
			groups[groupKey] = set
		}
	}

	return groups
}

func PrintlnGroups(groups map[int]TemperatureSet) {
	notFirst := false
	for groupKey, set := range groups {
		if notFirst {
			fmt.Print(", ")
		}
		fmt.Printf("%d:%v", groupKey, set.String())
		notFirst = true
	}
	fmt.Println()
}

func NewTemperatureSet() TemperatureSet {
	return TemperatureSet(make(map[float32]struct{}))
}

func (ts TemperatureSet) Add(t float32) {
	ts[t] = struct{}{}
}

func (ts TemperatureSet) String() string {
	sb := strings.Builder{}
	sb.WriteByte('{')
	notFirst := false
	for t := range ts {
		if notFirst {
			sb.WriteString(", ")
		}
		sb.WriteString(strconv.FormatFloat(float64(t), 'f', -1, 32))
		notFirst = true
	}
	sb.WriteByte('}')
	return sb.String()
}
