#!/bin/sh

# restore mecab setuped
cp -r /root/mecab-backup/lib/mecab /usr/local/lib/
cp /root/mecab-backup/lib/libmecab.* /usr/local/lib/
cp /root/mecab-backup/bin/* /usr/local/bin/

export CGO_LDFLAGS="`mecab-config --libs`"
export CGO_CFLAGS="-I`mecab-config --inc-dir`"
export CC=/usr/bin/clang

go get -u github.com/gin-gonic/gin
go get -u github.com/bluele/mecab-golang

go run ./src/
