package main

import "fmt"

type token struct {
	data      string
	recipient int
	ttl       int
}

func node(in <-chan token, out chan<- token, id int){ 
	TToken := <-in
	TToken.ttl--
	fmt.Println("Channel: ", id)
	//fmt.Println(TToken.ttl)
	if id == TToken.recipient {
		fmt.Println("I have got ur data: - ", TToken.data)
		return
	}else if TToken.ttl > 0{
		out <- TToken
	}else {
	    fmt.Println("Time is over")
	    return
	}
}

func main() {
	const N int = 10
	data := "Very important data"
	recipient := 3
    ttl := 5
	
	var chans [N]chan token 
	for i := 0; i < N-1; i++ {
		chans[i] = make(chan token)
	}
	for i := 0; i < N-1; i++ {
		go node(chans[i], chans[i+1], i)
	}
	chans[0] <- token{data, recipient, ttl}
	<-chans[0]
}