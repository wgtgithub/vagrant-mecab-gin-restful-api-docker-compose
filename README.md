Usage: 
- Using Vagrant
    - start app: vagrant up
    - finish app: vagrant halt

- Docker-compose
    - cd ./web/
    - docker-compose --build -d

- setup flask into your PC after installed python and pip
    - pip install -upgrade pip
    - if on MacOS append --user: pip upgrade pip --user
    - export FLASK_APP=./web/app/app.py

- run flask: start app
    - flask run
    - if stop press ctrl+c

- api check
    - curl http://localhost/targetwords/NiziProjectから選ばれたNiziUの９人のMember-Mako,Rio,Maya,Rio,Riku,Ayaka,Mayuka,MiihiそしてNina 

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

```
export CGO_LDFLAGS='-L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++'
export CGO_CFLAGS='-I/usr/local/Cellar/mecab/0.996/include'
```
    