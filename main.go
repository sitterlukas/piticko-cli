package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var cmdPiticko = &cobra.Command{
		Use:   "pojd",
		Short: "Send message to Kim when piticko",
		Long:  `Send message to Kim when we are going to get piticko.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			sendRequest(args[0])
			fmt.Printf("Piticko za %v\n", args[0])
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPiticko)
	rootCmd.Execute()
}

func sendRequest(time string) []byte {
	values := map[string]string{"time": time}
	jsonData, err := json.Marshal(values)

	request, err := http.NewRequest(
		http.MethodPost, //method
		"https://hook.eu1.make.com/fjs2jxbu52d7m1dal2pxc7dmgc93dbv5", //url
		bytes.NewBuffer(jsonData),                                    //body
	)

	request.Header.Set("X-Custom-Header", "myvalue")
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Printf("Request failed. %v", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}
