package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	var (
		year       int
		month      int
		dayInMonth int
	)

	for i := 0; i < 10; i++ {
		year = rng.Intn(3000) + 1
		month = rng.Intn(12) + 1
		dayInMonth = 31

		//检测是否是闰年
		switch month {
		case 2:
			if yeartest(year) {
				dayInMonth = 29
			} else {
				dayInMonth = 28
			}
		case 4, 6, 9, 11:
			dayInMonth = 30
		}

		day := rng.Intn(dayInMonth) + 1

		fmt.Println(era, day, month, year)
	}

}

func yeartest(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
