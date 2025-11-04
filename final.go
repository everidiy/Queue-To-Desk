package main

import "fmt"

func main() {
	queue := [5]string{"", "", "", "", ""}

	for {
		var input string
		fmt.Scan(&input)

		if input == "" {
			continue
		}

		if input == "очередь" {
			for i := 0; i < 5; i++ {
				if queue[i] == "" {
					fmt.Printf("%d. -\n", i+1)
				} else {
					fmt.Printf("%d. %s\n", i+1, queue[i])
				}
			}
			continue
		}

		if input == "количество" {
			free := 0
			occupied := 0
			for i := 0; i < 5; i++ {
				if queue[i] == "" {
					free++
				} else {
					occupied++
				}
			}
			fmt.Printf("Осталось свободных мест: %d\n", free)
			fmt.Printf("Всего человек в очереди: %d\n", occupied)
			continue
		}

		if input == "конец" {
			for i := 0; i < 5; i++ {
				if queue[i] == "" {
					fmt.Printf("%d. -\n", i+1)
				} else {
					fmt.Printf("%d. %s\n", i+1, queue[i])
				}
			}
			break
		}

		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Printf("Запись на место номер %d невозможна: \n", num)
			fmt.Printf("некорректный ввод")
			continue
		}

		if num < 1 || num > 5 {
			fmt.Printf("Запись на место номер %d невозможна: \n", num)
			fmt.Printf("некорректный ввод")
			continue
		}

		freeSlots := 0
		for i := 0; i < 5; i++ {
			if queue[i] == "" {
				freeSlots++
			}
		}
		if freeSlots == 0 {
			fmt.Printf("Запись на место номер %d невозможна: \n", num)
			fmt.Printf("очередь переполнена")
			continue
		}

		if queue[num-1] != "" {
			fmt.Printf("Запись на место номер %d невозможна: \n", num)
			fmt.Printf("место уже занято")
			continue
		}

		queue[num-1] = input

	}
}
