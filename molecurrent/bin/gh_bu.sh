#!/bin/sh
set -x
user="$1"
gh_user="$2"
test -z "$user" && echo "username and github username expected" && exit 1
test -z "$gh_user" && echo "github username expected" && exit 1
mkdir -p $HOME/Documents/gh_bu
cd $HOME/Documents/gh_bu
curl --silent "https://api.github.com/users/$gh_user/repos?per_page=100" | grep html_url | sort | /opt/plan9/bin/9 sed 's/.*"html_url": "(.*)",/\1/g' | uniq -u | parallel git clone {}
ls | parallel 'cd {} && git pull'
cd -
