package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// nodemon --exec go run . --signal SIGTERM

func main() {
	configBases, configsStr := getFreshPublicProxies()

	var outboundTags []string
	for _, configBase := range configBases {
		outboundTags = append(outboundTags, configBase.Tag)
	}

	outboundsFile, err := os.OpenFile("configs/outbounds.json", os.O_RDWR, 0644)
	defer outboundsFile.Close()
	if err != nil {
		log.Fatalf("failed opening outboundsFile: %s", err)
	}

	_, err = outboundsFile.WriteAt([]byte("{\"outbounds\":"+configsStr+"}"), 0)
	if err != nil {
		log.Fatalf("failed writing to outboundsFile: %s", err)
	}

	routing := getRouting()

	for index, balancer := range routing.Routing.Balancers {
		if balancer.Tag == "public-proxies" {
			routing.Routing.Balancers[index].Selector = outboundTags
		}
	}

	routingFile, err := os.OpenFile("configs/routing.json", os.O_RDWR, 0644)
	defer routingFile.Close()
	if err != nil {
		log.Fatalf("failed writing to routingFile: %s", err)
	}

	routingStr, err := json.Marshal(routing)
	if err != nil {
		fmt.Println(err)
	}
	_, err = routingFile.WriteAt(routingStr, 0)
	if err != nil {
		log.Fatalf("failed writing to routingFile: %s", err)
	}
}

func getRouting() RoutingFile {
	jsonFile, err := os.Open("configs/routing.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var routingFile RoutingFile
	err = json.Unmarshal(byteValue, &routingFile)
	if err != nil {
		fmt.Println(err)
	}
	return routingFile
}

func getFreshPublicProxies() ([]OutboundConfigBase, string) {
	resp, err := http.Get("https://raw.githubusercontent.com/MrMohebi/xray-proxy-grabber-telegram/master/proxies_active.txt")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	bodyStr := string(body)
	bodyStr = "[" + strings.Trim(strings.Replace(bodyStr, "\n", ",", -1), ",") + "]"

	var configs []OutboundConfigBase
	err = json.Unmarshal([]byte(bodyStr), &configs)
	if err != nil {
		fmt.Println(err)
	}
	return configs, bodyStr
}
