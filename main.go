package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// for {
	// 	var year int
	// 	fmt.Scan(&year)
	// 	if year == -1 {
	// 		break
	// 	}
	// 	kabisa := task1.Kabisa(year)
	// 	if kabisa {
	// 		fmt.Println("Bu yil kabisa yili, 366 kun bor")
	// 	} else {
	// 		fmt.Println("Bu yil kabisa yili emas, 365 kun bor")
	// 	}
	// }

	// fmt.Println(task1.Swapper(8, 8))
	// for {
	// 	var day, month int
	// 	fmt.Println("kunni kiriting: ")
	// 	fmt.Scan(&day)
	// 	fmt.Println("oyni kiriting: ")
	// 	fmt.Scan(&month)
	// 	fmt.Print("Keyingi sana: ")
	// 	fmt.Println(task1.NextDate(day, month))
	// 	fmt.Println()
	// 	if day == 0 {
	// 		break
	// 	}
	// }

	// fmt.Println(task1.PerfectNumber(1000))

	// fmt.Println(task1.IsPrime(11))

	

	timeElapsed := time.Since(start)
	fmt.Printf("it took %s\n", timeElapsed)
}
