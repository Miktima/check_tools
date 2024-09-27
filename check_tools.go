package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	var mode string
	var parameter string

	// Ключи для командной строки
	flag.StringVar(&mode, "mode", "", "mode of the check [urldecode, urlencode, jsonfile]")
	flag.StringVar(&parameter, "p", "", "incoming value (parameter)")

	flag.Parse()

	if mode == "urldecode" {
		query, err := url.QueryUnescape(parameter)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Result: ", query)
	} else if mode == "urlencode" {
		query := url.QueryEscape(parameter)
		fmt.Println("Result: ", query)
	} else if mode == "jsonfile" {
		// Читаем файл, если файл не можем открыть, сообщаем об ошибке
		_, err := os.Stat(parameter)
		if err != nil {
			fmt.Println("ERROR! File not found: ", err)
		}
		// Open our jsonFile
		byteValue, err := os.ReadFile(parameter)
		// if we os.ReadFile returns an error then handle it
		if err != nil {
			fmt.Println("ERROR! File not readable:", err)
		}
		var jsonval map[string]interface{}
		err = json.Unmarshal(byteValue, &jsonval)
		if err != nil {
			fmt.Println("ERROR! Wrong JSON: ", err)
		} else {
			fmt.Println("Result: JSON OK")
			fmt.Println("CDN Domains:", jsonval["workflow"].(map[string]interface{})["Domains"])
		}
	} else {
		fmt.Println("ERROR: unrecognized mode, valid = [urldecode, urlencode, jsonfile]")
	}
}
