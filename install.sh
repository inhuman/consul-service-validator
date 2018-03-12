#!/usr/bin/env bash

function jsonval {
    temp=`echo $json | sed 's/\\\\\//\//g' | sed 's/[{}]//g' | awk -v k="text" '{n=split($0,a,","); for (i=1; i<=n; i++) print a[i]}' | sed 's/\"\:\"/\|/g' | sed 's/[\,]/ /g' | sed 's/\"//g' | grep -w $prop`
    echo ${temp##*|}
}

json=`curl -s -X GET https://api.github.com/repos/inhuman/consul-service-validator/releases/latest`

prop='browser_download_url'
url=`jsonval`

mkdir -p /usr/local/bin/

curl -s -L $url > /usr/local/bin/validator

chmod +x /usr/local/bin/validator

