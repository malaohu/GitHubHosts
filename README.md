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
			alive.github.com
			live.github.com
			github.githubassets.com
			central.github.com
			desktop.githubusercontent.com
			assets-cdn.github.com
			camo.githubusercontent.com
			github.map.fastly.net
			github.global.ssl.fastly.net
			gist.github.com
			github.io
			github.com
			github.blog
			api.github.com
			raw.githubusercontent.com
			user-images.githubusercontent.com
			favicons.githubusercontent.com
			avatars5.githubusercontent.com
			avatars4.githubusercontent.com
			avatars3.githubusercontent.com
			avatars2.githubusercontent.com
			avatars1.githubusercontent.com
			avatars0.githubusercontent.com
			avatars.githubusercontent.com
			codeload.github.com
			github-cloud.s3.amazonaws.com
			github-com.s3.amazonaws.com
			github-production-release-asset-2e65be.s3.amazonaws.com
			github-production-user-asset-6210df.s3.amazonaws.com
			github-production-repository-file-5c1aeb.s3.amazonaws.com
			githubstatus.com
			github.community
			github.dev
			media.githubusercontent.com
# Last Update Time : 2024-08-09 00:01:45 
# Github: https://github.com/malaohu/GitHubHosts 
# Article: https://51.ruyo.net/17580.html 
####################Github End####################

```
