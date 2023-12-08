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

	result := fmt.Sprintf("####################Github Start####################\n")
	fmt.Printf("####################\n")

	for _, v := range url_list {
		ip := get_ip(v)
		result += fmt.Sprintf("%s\t\t\t%s\n", ip, v)

		fmt.Printf("%s\t\t\t%s\n", ip, v)
	}

	result += fmt.Sprintf("# Last Update Time : %s \n", time.Now().Format("2006-01-02 15:04:05"))
	result += "# Github: https://github.com/malaohu/GitHubHosts \n"
	result += "# Article: https://51.ruyo.net/17580.html \n"

	fmt.Printf("# Last Update Time :  %s \n", time.Now().Format("2006-01-02 15:04:05"))

	result += "####################Github End####################\n"
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
	
	ro := &grequests.RequestOptions{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36",

		//Headers: map[string]string{
		//	"Accept":          "*/*",
		//	"Accept-Encoding": "gzip, deflate, br",
		//	"Connection":      "keep-alive",
		//	"Cookie":          "ezoadgid_280870=-1; ezoref_280870=; ezosuibasgeneris-1=c20e8ba9-e495-4415-5c34-e3daa3d80321; ezoab_280870=mod1; lp_280870=https://sites.ipaddress.com/alive.github.com/; ezovuuid_280870=c991149d-3189-4f59-77e0-a1762c2c8385; ezds=ffid%3D1%2Cw%3D1920%2Ch%3D1080; ezohw=w%3D1920%2Ch%3D923; _ga=GA1.1.1440838470.1702002556; __qca=P0-410538264-1702002563689; active_template::280870=pub_site.1702002635; ezopvc_280870=4; ezovuuidtime_280870=1702002635; _ga_3GT2CLN45N=GS1.1.1702002556.1.1.1702002635.0.0.0; ezux_lpl_280870=1702002636435|32f307ab-be94-465d-499d-a271e85bf881|false",
		//},
	}

	resp, err := grequests.Get("https://ipaddress.com/website/"+url, ro)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	reg := regexp.MustCompile(`>(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})</a>`)

	result := reg.FindAllStringSubmatch(resp.String(), -1)

	ip := ""

	
	for _, text := range result {
		ip = text[1]
		break
	}

	return ip
}
