package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"time"
	"strings"
	"os"
	"github.com/levigross/grequests"
)

func main() {

	var url_list = [...]string{
		"alive.github.com",
		"live.github.com",
		"github.githubassets.com",
		"central.github.com",
		"desktop.githubusercontent.com",
		"assets-cdn.github.com",
		"camo.githubusercontent.com",
		"github.map.fastly.net",
		"github.global.ssl.fastly.net",
		"gist.github.com",
		"github.io",
		"github.com",
		"github.blog",
		"api.github.com",
		"raw.githubusercontent.com",
		"user-images.githubusercontent.com",
		"favicons.githubusercontent.com",
		"avatars5.githubusercontent.com",
		"avatars4.githubusercontent.com",
		"avatars3.githubusercontent.com",
		"avatars2.githubusercontent.com",
		"avatars1.githubusercontent.com",
		"avatars0.githubusercontent.com",
		"avatars.githubusercontent.com",
		"codeload.github.com",
		"github-cloud.s3.amazonaws.com",
		"github-com.s3.amazonaws.com",
		"github-production-release-asset-2e65be.s3.amazonaws.com",
		"github-production-user-asset-6210df.s3.amazonaws.com",
		"github-production-repository-file-5c1aeb.s3.amazonaws.com",
		"githubstatus.com",
		"github.community",
		"github.dev",
		"media.githubusercontent.com"}

	result := fmt.Sprintf("####################\n")
	fmt.Printf("####################\n")

	for _, v := range url_list {
		ip := get_ip(v)
		result += fmt.Sprintf("%s\t\t\t%s\n", ip, v)

		fmt.Printf("%s\t\t\t%s\n", ip, v)
	}

	result += fmt.Sprintf("# Last Update Time : %s \n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("# Last Update Time :  %s \n", time.Now().Format("2006-01-02 15:04:05"))

	result += "####################"
	fmt.Printf("####################")
	//保存文件
	ioutil.WriteFile(`hosts.txt`, []byte(result), 0666)


	content := read_tmp()
	content = strings.Replace(content, "12345", result, 1)
	write_file(content)

}



func read_tmp() string{
	file, err := os.Open("README_TEMP.md")
	if err != nil {
		panic(err)
	}	
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return string(content)
}


func write_file(content string){
	err := ioutil.WriteFile("README.md", []byte(content), 0644)
	if err != nil {
		panic(err)
	}	

}


func get_ip(url string) string {
	resp, err := grequests.Get("https://ipaddress.com/website/"+url, nil)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	reg := regexp.MustCompile(`<li>(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})</li>`)

	result := reg.FindAllStringSubmatch(resp.String(), -1)

	ip := ""

	for _, text := range result {
		ip = text[1]
		break
	}

	return ip
}
