package main

import "fmt"

func main() {
	swtch := make([]bool, 100)
	// fmt.Println(swtch)
	var count int32

	for i := 1; i <= 100; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= len(swtch); k++ {
				if k%j == 0 {
					// fmt.Println(k)
					if swtch[k-1] == false {
						swtch[k-1] = true
						count++
					} else {
						swtch[k-1] = false
						count--
					}
				}
			}
			// for k, v := range swtch {
			// 	if k%j == 0 {
			// 		fmt.Println(k)
			// 		if v == false {
			// 			swtch[k] = true
			// 			count++
			// 		} else {
			// 			swtch[k] = false
			// 			count--
			// 		}
			// 	}
			// }
			// fmt.Println(swtch)
		}
		i += 3
	}
	// fmt.Println(swtch)

	fmt.Println(count)
}
