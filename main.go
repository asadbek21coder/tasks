package main

import (
	"fmt"
	"time"

	"github.com/asadbek21coder/tasks/task1"
)

func main() {
	start := time.Now()

	// var limit int
	// var array []int
	// fmt.Println("Elementlar sonini kiritng: ")
	// fmt.Scan(&limit)

	// for i := 0; i < limit; i++ {
	// 	var num int
	// 	fmt.Scan(&num)
	// 	array = append(array, num)
	// }
	// fmt.Println(task1.OddIncEvenDec(array))

	// var array []int = []int{2,5,3,67,45,32,76,368,79,756,4353,42}
	// / for {
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

	// res := task1.Sign(-2)
	// fmt.Println(task1.Sign(9845))
	// res := task1.Descriminant(2, 14	, 5)
	// fmt.Println(res)
	// fmt.Println(task1.SumArray(3, 4, 15))

	// dif := task1.DifferenceDiskSphere(8.0, 5.0)
	// fmt.Println(dif)
	var array []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(task1.ReverseArray(array))

	tayyorMassiv, soni := task1.EvenIndexDesc(array)
	fmt.Println(tayyorMassiv, soni)
	timeElapsed := time.Since(start)
	fmt.Printf("it took %s\n", timeElapsed)
}

//  A*x^2 + Bx + C
//  D = B^2 - 4*A*C >0 2  =0 1 <0 0 ta yechim
//
