package main

import (
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile(`C:\Users\Admin\Desktop\nihk.txt`,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, // 覆盖写入
		0644,
	)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	for {
		time.Sleep(1 * time.Second)
		data := []byte("锟斤拷?!!%#@@@")
		_, err := file.Write(data)
		if err != nil {
			log.Print(err)
		}
	}

}
