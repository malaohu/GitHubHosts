# 目的
自动获取Github.com下的域名对应解析到的IP，针对国内用户解决无法正常访问Github的问题

# 思路
参考文章：https://51.ruyo.net/17580.html

# 源码
使用Golang编写，第一次用golang写一些小东西。有不对的地方请大佬们指点！

# 运行
```
#本地需要安装golang环境
go run .\doit.go
```

Windows 可直接运行 github_hosts.exe

Linux 可直接运行 github_hosts

# HOSTS

```bash

####################Github Start####################
140.82.113.25			alive.github.com
			live.github.com
185.199.108.154			github.githubassets.com
			central.github.com
			desktop.githubusercontent.com
185.199.108.153			assets-cdn.github.com
			camo.githubusercontent.com
185.199.108.133			github.map.fastly.net
151.101.1.194			github.global.ssl.fastly.net
			gist.github.com
			github.io
140.82.113.3			github.com
			github.blog
			api.github.com
185.199.108.133			raw.githubusercontent.com
185.199.108.133			user-images.githubusercontent.com
			favicons.githubusercontent.com
			avatars5.githubusercontent.com
185.199.108.133			avatars4.githubusercontent.com
			avatars3.githubusercontent.com
185.199.108.133			avatars2.githubusercontent.com
185.199.108.133			avatars1.githubusercontent.com
			avatars0.githubusercontent.com
185.199.108.133			avatars.githubusercontent.com
140.82.112.10			codeload.github.com
3.5.28.232			github-cloud.s3.amazonaws.com
			github-com.s3.amazonaws.com
3.5.29.54			github-production-release-asset-2e65be.s3.amazonaws.com
3.5.0.3			github-production-user-asset-6210df.s3.amazonaws.com
			github-production-repository-file-5c1aeb.s3.amazonaws.com
185.199.108.153			githubstatus.com
140.82.114.18			github.community
52.224.38.193			github.dev
185.199.108.133			media.githubusercontent.com
# Last Update Time : 2024-09-10 12:08:46 
# Github: https://github.com/malaohu/GitHubHosts 
# Article: https://51.ruyo.net/17580.html 
####################Github End####################

```
