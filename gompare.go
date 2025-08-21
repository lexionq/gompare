package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

var colors = map[string]string{
	"boldRed":   "\033[1;31m",
	"boldGreen": "\033[1;32m",
	"boldBlue":  "\033[1;34m",
	"reset":     "\033[0m",
}

func calculateSHA256(filePath string) (string,error){
	file, err := os.Open(filePath)
	if err != nil {
		return "[ERROR] File's hash didn't calculate: ",err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash,file); err != nil {
		return "",err
	}
	
	return hex.EncodeToString(hash.Sum(nil)),nil
}

func main(){

	if len(os.Args) != 3{
		fmt.Println(colors["boldBlue"],"Usage: ",colors["reset"])
		fmt.Println(`
	gompare <file1> <file2>
		`)
		os.Exit(0)
	}

	file1, file2 := os.Args[1], os.Args[2]

	hash1, err := calculateSHA256(file1)
	if err != nil {
		log.Fatal(colors["boldGreen"],"[ERROR] Error during calculate hash: ",colors["reset"],err)
	}

	hash2, err := calculateSHA256(file2)
	if err != nil {
		log.Fatal(colors["boldGreen"],"[ERROR] Error during calculate hash: ",colors["reset"],err)
	}


	fmt.Println(colors["boldBlue"],"File1 Hash: ",colors["reset"],hash1)

	fmt.Println(colors["boldBlue"],"File2 Hash: ",colors["reset"],hash2)

	
	if hash1 != hash2{
		fmt.Println(colors["boldRed"],"[WARNING] Hashes are not the same.",colors["reset"])
	} else {
		fmt.Println(colors["boldGreen"],"[INFO] Hashes are the same, file downloaded successfully.",colors["reset"])
	}

}