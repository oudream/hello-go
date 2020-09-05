#!/usr/bin/env bash

go env

go build -v -work -o hello.exe


### build
-a	  # 强行对所有涉及到的代码包（包含标准库中的代码包）进行重新构建，即使它们已经是最新的了。
-n	  # 打印编译期间所用到的其它命令，但是并不真正执行它们。
-p n	# 指定编译过程中执行各任务的并行数量（确切地说应该是并发数量）。在默认情况下，该数量等于CPU的逻辑核数。但是在darwin/arm平台（即iPhone和iPad所用的平台）下，该数量默认是1。
-race	# 开启竞态条件的检测。不过此标记目前仅在linux/amd64、freebsd/amd64、darwin/amd64和windows/amd64平台下受到支持。
-v	  # 打印出那些被编译的代码包的名字。
-work	# 打印出编译时生成的临时工作目录的路径，并在编译结束时保留它。在默认情况下，编译结束时会删除该目录。
-x	  # 打印编译期间所用到的其它命令。注意它与-n标记的区别。


### install
wget https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz
sudo tar -xvf go1.12.6.linux-amd64.tar.gz
sudo mv go /usr/local

cat >> ~/.bash_profile << EOF
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
export GOPATH=/home/gopath
EOF


### go可以设置多个GOPATH
# Linux下用冒号(:)分割，例如：
GOPATH="/home/ferghs/gowork:/home/ferghs/gowork/src/project1"
# Windows使用分号分割(;)
# 1 如果使用go get 默认安装到第一个GOPATH路径
# 2 编译(go build)时，有时会报同一种类型或方法不匹配，原因是GOPATH路径中多个路径顺序不对，调换一下就OK了