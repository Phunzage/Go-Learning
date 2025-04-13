//每隔两秒监视文件改动

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	files, err := os.Open("main.txt")
	if err != nil {
		log.Print(err)
	}

	defer files.Close()

	for {
		time.Sleep(2 * time.Second)
		content, err := os.ReadFile("main.txt")
		if err != nil {
			log.Print(err)
		}
		fmt.Println(string(content))
	}

}
