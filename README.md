# Gin-playground

Build to find an opinionated way for better building modern web app with your favorite language Golang XD. Therefore this is a place for idea testing.

## Getting Started

Project is build with the power of docker. All dependencies are configured at the dockerfile. So if you have a proper docker and docker-compose running then running this project should be no problem for you.

### Plans
- [ ] Dockerize with docker-compose
- [ ] Router
- [ ] Authentication
- [ ] Controller params validation
- [ ] User friendly error message
- [ ] Simple and beautiful doc auto generation
- [ ] Debugging method
- [ ] Opinions on both unit test and integration test
- [ ] Access Controls
- [ ] Middlewares integration
- [ ] Database migration
- [ ] SQL builder or better solution on database interaction
- [ ] Third party api interaction package example
- [ ] Secrets management
- [ ] RPC
- [ ] Socket usage for realtime implementation
- [ ] Microservice with package like go-kit
- [ ] Continues deployment
- [ ] I18n
- [ ] Security control
- [ ] Shell command helpers
- [ ] Elastic stack for information gathering
- [ ] Integrate API gateway service like: Kong

### Prerequisites

*** You just need docker ***

```shell
$ docker version
Version: 17.12.0-ce

$ docker-compose version
docker-compose version 1.18.0, build 8dd22a9
```

### Installing

Running the app in detached mode.

```shell
$ docker-compose up -d
```

And check with curl

```shell
curl -i -X GET http://localhost:8080/ping
```

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

## Built With

* [Golang](https://golang.org/) - Main language

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **BranLiang** - *Initial work* - [HOMEPAGE](http://www.liangboyuan.pub)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* [golang-gin-realworld-example-app](https://github.com/gothinkster/golang-gin-realworld-example-app)
