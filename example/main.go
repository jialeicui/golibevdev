package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jialeicui/golibevdev"
)

var flagKeyboard = flag.String("keyboard", "HHKB", "Keyboard name to watch")

func main() {
	flag.Parse()

	paths, err := getAllInputDevices()
	if err != nil {
		panic(err)
	}

	var theKeyboard *golibevdev.InputDev

	for _, path := range paths {
		dev, err := golibevdev.NewInputDev(path)
		if err != nil {
			fmt.Printf("Failed to open %s: %s\n", path, err)
			continue
		}
		fmt.Printf("Path: %s, Name: %s\n", path, dev.Name())

		if strings.Contains(dev.Name(), *flagKeyboard) {
			theKeyboard = dev
			break
		}

		dev.Close()
	}

	if theKeyboard == nil {
		fmt.Printf("%s not found\n", *flagKeyboard)
		return
	}

	fmt.Printf("Watching %s\n", theKeyboard.Name())

	udev, err := golibevdev.NewVirtualKeyboard("golibevdev")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer udev.Close()
	watchEvent(theKeyboard, udev)
}

func getAllInputDevices() ([]string, error) {
	var paths []string
	devices, err := os.ReadDir("/dev/input")
	if err != nil {
		return nil, fmt.Errorf("failed to read /dev/input: %w", err)
	}
	for _, device := range devices {
		if device.IsDir() {
			continue
		}
		paths = append(paths, "/dev/input/"+device.Name())
	}
	return paths, nil
}

func watchEvent(dev *golibevdev.InputDev, out *golibevdev.UInputDev) {
	if err := dev.Grab(); err != nil {
		fmt.Println(err)
		return
	}

	for {
		ev, err := dev.NextEvent(golibevdev.ReadFlagNormal)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("Type: %v, Code: %v, Value: %v\n", ev.Type, ev.Code, ev.Value)

		// replace key b to a
		if ev.Type == golibevdev.EvKey {
			if ev.Code == golibevdev.KeyB {
				ev.Code = golibevdev.KeyA
			}
			if ev.Code == golibevdev.KeyL {
				continue
			}
		}

		if err := out.WriteEvent(ev.Type, ev.Code, ev.Value); err != nil {
			fmt.Println(err)
			break
		}
	}
}
