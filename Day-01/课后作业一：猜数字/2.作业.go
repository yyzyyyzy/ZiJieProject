package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectNumGame2() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	selectNum := rand.Intn(maxNum)
	// fmt.Printf("我们需要猜的数为%v\n", selectNum)

	var guess int
	for {
		fmt.Printf("请输入0-100的整数:")
		fmt.Scan(&guess)

		if guess > selectNum {
			fmt.Println("你猜测的数字过大")
		} else if guess < selectNum {
			fmt.Println("你猜测的数字过小")
		} else {
			fmt.Println("猜对了！")
			break
		}
	}
}

func main() {
	selectNumGame2()
}
