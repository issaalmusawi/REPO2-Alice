package main

import (
	"net"
	"log"
	"os"
	"github.com/issaalmusawi/repo3-crypt/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.4:8080")
	if err != nil {
		log.Fatal(err)
	}
    	
	//nytt

	message :=[]rune(os.Args[1])
	encryptedMessage, err := mycrypt.Krypter(message, 4)
	if err != nil{
		log.Fatal(err)
	}

	log.Println("Kryptert melding: ", string(message))

 	_, err = conn.Write([]byte(string(encryptedMessage)))
	if err != nil {
		log.Fatal(err)
		return
	}
	
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil{
		log.Println(err)
		return
	}

	encryptedResponse :=[]rune(string(buf[:n]))
	decryptedResponse, err := mycrypt.Krypter(encryptedResponse, -4)
	if err != nil{
		log.Fatal(err)
	}
	
	log.Println("reply from proxy: ", string(decryptedResponse))



/*	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 
	response := string(buf[:n])
	log.Printf("reply from proxy: %s", response)*/
}
