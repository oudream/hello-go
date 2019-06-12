#!/usr/bin/env bash

cd /ddd/ops; git clone https://github.com/etcd-io/etcd.git
ln -s /ddd/ops/etcd/ /fff/gopath/src/go.etcd.io/etcd/v3
cd /fff/gopath/src/github.com; mkdir coreos; cd coreos
git clone https://github.com/coreos/pkg.git
cd /fff/gopath/src/go.etcd.io/etcd/v3/contrib/raftexample
go build -o raftexample # must build use package
