package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Fardan Nozami Ajitama"
		fmt.Println("selesai mengirim data ke channel")
	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Fardan Nozami Ajitama"
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Halo Dunia!"
}

func OnlyOut(channel <-chan string) {
	fmt.Println(<-channel)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Fardan"
	channel <- "Nozami"
	channel <- "Ajitama"

	fmt.Println(<-channel) // Fardan
	fmt.Println(<-channel) // Nozami
	fmt.Println(<-channel) // Ajitama
	fmt.Println("selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i+1)
		}
		close(channel)
	}()

	// for loop
	for data := range channel {
		fmt.Println("menerima data", data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima dari channel 1 : ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Menerima dari channel 2 : ", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima dari channel 1 : ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Menerima dari channel 2 : ", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
