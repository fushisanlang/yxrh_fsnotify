#!/bin/bash
#by https://github.com/fushisanlang/


## judge user
Username=`whoami`
if [[ ${Username} != root ]]; then
    echo -e "\e[1;31mYou are not root. \e[0m" 
    exit 1
fi

## judge install 
echo -en "\e[1;33mReady to install yxrh_fsnotify,is this ok??? [y/n]  \e[0m" 
read judgement
if [[ ${judgement} != y ]]; then
    echo -e "\e[1;32mOK, Bye. \e[0m" 
    exit 1
else
    echo -e "\e[1;32mLet's go. \e[0m"
fi

cd /usr/local/
rm -fr yxrh_fsnotify.tar.gz
cd /usr/local/yxrh_fsnotify
chmod +x sbin/*
cp src/install/yxrh_fsnotify.sh /etc/init.d/yxrh_fsnotify
chmod +x /etc/init.d/yxrh_fsnotify
echo -e "\e[1;32mSuccessful installation." 
echo -e "you can learn more at https://github.com/fushisanlang/own_note/tree/master/go/fsnotify .\e[0m "
