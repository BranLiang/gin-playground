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
- [ ] Database setup
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

### Database

Use [database/sql](https://golang.org/pkg/database/sql/) and [pg](https://github.com/lib/pq). The decision is made because as a ruby on rails developer, I am so used to the ActiveRecord as if I almost forgot how to use the native sql driver. The problem then comes when you want to more effciently control your database, the ORMs normally can hardly help and they also add too much complexity on top of the database layer. So I choose to not use ORM for this project as I want to get more close to my database. Here is another good article which also expresses my opinion - [golang-orms-and-why-im-still-not-using-one](http://www.hydrogen18.com/blog/golang-orms-and-why-im-still-not-using-one.html).

Database has its own package called db, when db is imported its `init` function will be called and a database connection will be made. You can retrive the database instance which is actually a pointer through `db.DB` method.

Database commands like `db setup`, `db create/drop` or `db migrate` are not implemented for the moment, because the migration function is a heavy wrapper on top the database layer, The responsiblity if left for the user to properly maintain the database schema. My aim for this project is not to introduce a lot new commands which could lead to harder maintaince but a simple glue for modules.

### CLI

[Cobra](https://github.com/spf13/cobra) is choosen as the tool for building the application commands. Reason is simply because I am more familiar with this cli framework.

* Command arguments are checked, both on argument number and name.

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

Start the server with command

```shell
$ gin-playground server
```

And check with curl

```shell
curl -i -X GET http://localhost:8080/ping
```

## Development

Create the database

```shell
$ psql
psql (10.1)
Type "help" for help.

bran=# CREATE DATABASE gin_playground;
bran=# \q
```

Install the extension [pgcrypto](https://starkandwayne.com/blog/uuid-primary-keys-in-postgresql/) for UUID usage

```shell
$ psql gin_playground
psql (10.1)
Type "help" for help.

gin_playground=# CREATE EXTENSION pgcrypto;
gin_playground=# \q
```

Load the setup tables from `db/schema.sql`
```shell
$ psql -f ./db/schema.sql gin_playground
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
* [pg](https://github.com/lib/pq) - Database driver
* [Cobra](https://github.com/spf13/cobra) - CLI framework

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **BranLiang** - *Initial work* - [HOMEPAGE](http://www.liangboyuan.pub)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* [golang-gin-realworld-example-app](https://github.com/gothinkster/golang-gin-realworld-example-app)
* [gin-boilerplate](https://github.com/Massad/gin-boilerplate)
