package main

import (
	"crypto/md5"
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
	// urlencode - кодирует url escape символами
	// urldecode - декодирует url
	// jsonfile - проверяет синтаксис json файла и одновременно выводит параметры CDN если есть
	// md5 - преобразует строку в md5 hash
	// help - вывод помощи
	flag.StringVar(&mode, "mode", "", "mode of the check [urldecode, urlencode, jsonfile, md5, h]")
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
	} else if mode == "md5" {
		hash := md5.Sum([]byte(parameter))
		fmt.Printf("Result: %x\n", hash)
	} else if mode == "help" {
		fmt.Println("urlencode - кодирует url escape символами")
		fmt.Println("urldecode - декодирует url")
		fmt.Println("jsonfile - проверяет синтаксис json файла и одновременно выводит параметры CDN если есть")
		fmt.Println("md5 - преобразует строку в md5 hash")
		fmt.Println("help - вывод помощи")
		fmt.Println("ПРИМЕР: ./check_tools -mode urlencode -p 'ABC!@23%'")
	} else {
		fmt.Println("ERROR: unrecognized mode, valid = [urldecode, urlencode, jsonfile, md5, h]")
	}
}
