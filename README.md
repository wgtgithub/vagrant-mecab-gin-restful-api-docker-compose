Usage: 
- Using Vagrant
    - start app: vagrant up
    - finish app: vagrant halt

- Docker-compose
    - cd ./web/
    - docker-compose --build -d

- setup go into your MacOS
    - https://golang.org/doc/install

- run gin: start app
    - go run ./web/src/
    - if stop press ctrl+c

- api check
    - curl http://localhost:8080/targetwords?letter=NiziUのDebutは間近ですといってもまだ２ヶ月ありますが 

About Docker Image
- Dockerfile uses wgtdocker/mecab-python-docker:latest as docker pull image
- You can check the resource the following url
- https://github.com/wgtgithub/vagrant-mecab-python-docker

Note:
- I am not modified this resorces forever
- If you clone it and then something happened but I won't follow you

- use in MacOS
    - Do the following commands before run "go get github.com/bluele/mecab-golang"
    ```
    $ export CGO_LDFLAGS="-L/to/lib -lmecab -lstdc++"
    $ export CGO_CFLAGS="-I/to/include"
    $ export CC=/usr/bin/clang
    $ go get github.com/bluele/mecab-golang
    ```
    - Installed mecab-config then do the following commands,
    ```
    $ export CGO_LDFLAGS="`mecab-config --libs`"
    $ export CGO_CFLAGS="-I`mecab-config --inc-dir`"
    $ go get github.com/bluele/mecab-golang
    ```

    