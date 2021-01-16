# Simple-microservices

A simple microservices project. Everything is simply for demo purpose.

![](https://img.shields.io/badge/Environment-Docker-blue)
[![](https://img.shields.io/badge/Owner-minhluantran017-darkviolet)](mailto:minhluantran017@gmail.com)

![](https://github.com/minhluantran017/simple-microservices/workflows/Check%20Docker%20images/badge.svg)

## What services inside?

- Python
- NodeJS
- Golang
- and so on, depend on my free time. :)

## How to run it?

### Local with Docker

TBD

### Local with Docker compose

I use Docker compose and Traefik proxy to eliminate the port append:

```sh
docker-compose up
```

And then open your browser and oint to these endpoints:

* http://127.0.0.1/python
* http://127.0.0.1/nodejs
* http://127.0.0.1/golang

### CI/CD

When you make any change to this repository, GitHub Actions will be triggered to build Docker images for these services, deploy and curl to each services to get content.