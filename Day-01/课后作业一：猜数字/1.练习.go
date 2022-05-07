package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//随机数需要使用的随机数种子rand.Seed()，不然会生成相同的随机数
func main() {
	selectNumGame1()
}

func selectNumGame1() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	selectNum := rand.Intn(maxNum)
	// fmt.Printf("我们需要猜的数为%v\n", selectNum)

	fmt.Printf("请输入0-100的整数:")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		HandleErr("reader.ReadString", err)

		input = strings.TrimSuffix(input, "\r\n")

		guess, err := strconv.Atoi(input)
		HandleErr("strconv.Atoi", err)
		fmt.Println("你猜测的数字是:", guess)

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

func HandleErr(when string, err error) {
	if err != nil {
		fmt.Println(err, when)
		return
	}
}
