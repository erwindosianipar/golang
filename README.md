# Golang

## Download Golang
Download dan install Golang pada PC
```
https://golang.org/dl/
```

## Visual Studio Code & Go Extensions
Download Visual Studio Code dan lakukan instalasi Go Extensions

## Membuat dan memulai project
Buat folder untuk project Golang
```
D:\golang-project-name
```

## GOPATH dan Workspace
Cara setting GOPATH pada Windows
```
system -> advanced system settings -> environment variabel -> PATH_TO_GOLANG_PROJECT
```
Workspace Golang project
Buat folder untuk project Golang <code>bin</code>, <code>pkg</code>, <code>src</code>
Folder src adalah folder dimana project Golang disimpan

```
$ cd PATH_TO_GOLANG_PROJECT

$ go get -v github.com/go-delve/delve/dlv
$ go install
```
## Debugging
```
$ go run file.go
```
