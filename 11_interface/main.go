package main

import (
	"fmt"
	"golang/11_interface/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}

}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixTape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixTape)
	player = gadget.TapeRecorder{}
	playList(player, mixTape)
	fmt.Println("----------------------------------------------------------")
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
