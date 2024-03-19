package client

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"gotest/src/gui"
	"net"
	"sync"
)

func Run(wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.Dial("tcp", "localhost:2137")
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Fprintf(conn, "Hello dude")
	//conn.Write([]byte{0})

	gameGui := gui.GameGui{}
	go gameGui.Run()
	game := &gameGui.Game

	for {
		data := make([]byte, 64)
		var buffer = bytes.NewBuffer(data)
		var decoder = gob.NewDecoder(buffer)
		//fmt.Println("waiting for data")
		_, err := conn.Read(buffer.Bytes())
		//fmt.Println("received data")
		if err != nil && err.Error() != "EOF" {
			fmt.Println(err)
		}

		err = decoder.Decode(&(*game).GameObject)
		if err != nil && err.Error() != "EOF" {
			fmt.Println(err)
		}

		//fmt.Println("client: ", game.GameObject)
	}
}
