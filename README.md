# Introduction
Web service that lists of pickup locations for available in customer area so that he can have his package
delivered to the nearest location in his neighbourhood.

## Table of Contents

- [Installation](#installation)
---
## Installation
To get the project up and running you should have the following prerequisite
- `git` which comes preinstalled with Linux & Mac on Windows you can Download Git[ Here]()
- `docker` you can download Docker [Here](https://www.docker.com/get-started)

### Clone
When you have all the requirement installed you can start by cloning the Repo
- Run ```git clone https://github.com/amine-bambrik-p8/go-lang-web-service```
### Run Docker Container
To run the `docker` container you have:
- First run ```cd go-lang-web-service```
- Then  build ```docker build --pull --rm -f "Dockerfile" -t golangwebservice:latest "."``` to build your image
- Finally you can run the container using ```docker run --rm -it  -p 8081:8081/tcp golangwebservice:latest```
- You can stop the container by pressing ```Ctrl + C``` 
