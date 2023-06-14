package main

import (
	"fmt"

	"github.com/asadbek21coder/tasks/task1"
)

func main() {

	kabisa := task1.Kabisa(2024)
	if kabisa {
		fmt.Println("Bu yil kabisa yili, 366 kun bor")
	} else {
		fmt.Println("Bu yil kabisa yili emas, 365 kun bor")
	}
}
