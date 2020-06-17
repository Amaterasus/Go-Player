package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	file, error := os.Open("./test.mp3")
	if error != nil {
		log.Fatal(error)
	}

	streamer, format, error := mp3.Decode(file)
	if error != nil {
		log.Fatalf("error while decoding into mp3: %v\n", error)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}