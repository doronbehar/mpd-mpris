package main

import (
	"flag"
	"fmt"
	"log"

	mpris "github.com/natsukagami/mpd-mpris"
	"github.com/natsukagami/mpd-mpris/mpd"
)

var (
	addr     string
	port     int
	password string

	noInstance bool
)

func init() {
	flag.StringVar(&addr, "host", "localhost", "The MPD host.")
	flag.IntVar(&port, "port", 6600, "The MPD port")
	flag.StringVar(&password, "pwd", "", "The MPD connection password. Leave empty for none.")
	flag.BoolVar(&noInstance, "no-instance", false, "Set the MPDris's interface as 'org.mpris.MediaPlayer2.mpd' instead of 'org.mpris.MediaPlayer2.mpd.instance#'")
}

func main() {
	flag.Parse()

	// Attempt to create a MPD connection
	var (
		c   *mpd.Client
		err error
	)

	if password == "" {
		c, err = mpd.Dial("tcp", fmt.Sprintf("%s:%d", addr, port))
	} else {
		c, err = mpd.DialAuthenticated("tcp", fmt.Sprintf("%s:%d", addr, port), password)
	}

	if err != nil {
		panic(err)
	}

	opts := []mpris.Option{}
	if noInstance {
		opts = append(opts, mpris.NoInstance())
	}

	instance, err := mpris.NewInstance(c, opts...)

	if err != nil {
		panic(err)
	}
	defer instance.Close()

	log.Println("mpd-mpris running")

	<-make(chan int)
}
