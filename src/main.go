package main

import (
	"fmt"
	"gotest/src/client"
	"gotest/src/server"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	if len(os.Args[1:]) > 0 {
		fmt.Println("Im a Server")
		go server.Run(&wg)
	} else {
		fmt.Println("Im a Client")
		go client.Run(&wg)
	}

	wg.Wait()

}
