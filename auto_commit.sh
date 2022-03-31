./output/github_hosts

time=$(date "+%Y-%m-%d %H:%M:%S")
git commit -am "auto commit hosts "

echo "auto commit hosts "${time}

git push
