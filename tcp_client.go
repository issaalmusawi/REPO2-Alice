package main

import (
	"net"
	"log"
	"os"
	"github.com/issaalmusawi/repo3-crypt/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.2:8080")
	if err != nil {
		log.Fatal(err)
	}
    	
	//nytt

	message :=[]rune(os.Args[1])
	kryptertMelding, err := mycrypt.Krypter(message, 4)
	if err != nil{
		log.Fatal(err)
	}

	log.Println("Kryptert melding: ", string(kryptertMelding))

 	_, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 
	response := string(buf[:n])
	log.Printf("reply from proxy: %s", response)
}
