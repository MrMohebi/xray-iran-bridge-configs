package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// nodemon --exec go run . --signal SIGTERM

func main() {
	var UrlProxies = "https://raw.githubusercontent.com/MrMohebi/xray-proxy-grabber-telegram/master/collected-proxies/xray-json/actives_no_403_under_1000ms.txt"
	if len(os.Getenv("URL_PROXY_FILE")) > 3 {
		UrlProxies = os.Getenv("URL_PROXY_FILE")
	}
	configBases, configsStr := getFreshPublicProxies(UrlProxies)

	var outboundTags []string
	for _, configBase := range configBases {
		outboundTags = append(outboundTags, configBase.Tag)
	}

	err := os.WriteFile("configs/outbounds.json", []byte("{\"outbounds\":"+configsStr+"}"), 0644)
	if err != nil {
		log.Fatalf("failed writing to outboundsFile: %s", err)
	}
	println("updated outbounds.json")

	routing := getRouting()
	for index, balancer := range routing.Routing.Balancers {
		if balancer.Tag == "public-proxies" {
			routing.Routing.Balancers[index].Selector = outboundTags
			routing.BurstObservatory.SubjectSelector = outboundTags
		}
	}

	routingStr, err := json.Marshal(routing)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("configs/routing.json", routingStr, 0644)
	if err != nil {
		log.Fatalf("failed writing to routingFile: %s", err)
	}
	println("updated routing.json")

	reloadXrayCore()

}

func reloadXrayCore() {
	cmdRebootXray := exec.Command("/bin/sh", "/root/xray-core/rebootXray.sh")
	err := cmdRebootXray.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("xray core Reloaded.")
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

func getFreshPublicProxies(url string) ([]OutboundConfigBase, string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	bodyStr := string(body)
	bodyStr = "[" + strings.Trim(strings.Replace(bodyStr, "\n", ",", -1), ",") + ",{\"tag\": \"direct-out\",\"protocol\": \"freedom\"}]"

	var configs []OutboundConfigBase
	err = json.Unmarshal([]byte(bodyStr), &configs)
	if err != nil {
		fmt.Println(err)
	}
	println("Got new proxies from => " + url)
	return configs, bodyStr
}
