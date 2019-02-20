## Building Minimal Docker Containers for Go Applications
### build go hello world app
```sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```
### build docker image 
```sh
docker build -t hello-world .
```
### show list docker image 
```sh
docker image list
```
### run docker image 
```sh
docker run -it -p 8080:8080 hello-world
```
### remove docker image 
```sh
docker rmi --force hello-world:latest
```