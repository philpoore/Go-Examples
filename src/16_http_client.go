package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://api-v3.easyfundraising.org.uk/version")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	var data map[string]interface{}
	for scanner.Scan() {
		json.Unmarshal(scanner.Bytes(), &data)
		fmt.Println(data)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
