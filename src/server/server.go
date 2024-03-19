package server

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"gotest/src/game"
	"gotest/src/gui"
	"math"
	"net"
	"sync"
	"time"
)

func server(listeners *[]*net.Conn, game *game.Game) {
	var i int = 0
	for {
		i++
		fmt.Println("server tick ", i)
		//fmt.Println(listeners)
		time.Sleep(time.Millisecond * 50)
		game.GameObject.X = math.Sin((float64)(i) / 10)
		game.GameObject.Y = math.Cos((float64)(i) / 10)
		fmt.Println("server: ", game.GameObject)

		var buffer = new(bytes.Buffer)
		var encoder = gob.NewEncoder(buffer)
		encoder.Encode(game.GameObject)

		for _, l := range *listeners {
			(*l).Write(buffer.Bytes())
		}
	}
}

func Run(wg *sync.WaitGroup) {
	defer wg.Done()

	ln, err := net.Listen("tcp", ":2137")
	if err != nil {
		fmt.Println(err)
	}

	listeners := make([]*net.Conn, 0)

	gameGui := gui.GameGui{}
	go server(&listeners, &gameGui.Game)
	go gameGui.Run()

	for {
		conn, err := ln.Accept()
		fmt.Println("klient sie odezwal")
		if err != nil {
			fmt.Println(err)
		}
		listeners = append(listeners, &conn)
	}
}
