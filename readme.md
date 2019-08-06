Golang-Http-test
============

Testing Golang's Http capabilities without HTTP Packages

Only uses packages for sqlite3 stateful db for local dev.

Dependencies
------------
* Glog - Google's logging used by Kubernetes
* Go-sqlite3 - Local DB for lightweight sql

Usage Local Docker Setup
--------
```
git clone git@github.com:kenichi-shibata/golang-http-test
cd golang-http-test
mkdir db
docker run -it -v $(pwd)/db:/app/db/ -p 8080:8080 quay.io/kenichi_shibata/golang-http-test:15813e9
```
If you are running on mac you might need to enable mounts on docker. https://docs.docker.com/docker-for-mac/osxfs/#namespaces

Once you have the application running you can run a series of request for checking.

```
sh hack/test.curl
```

The DB Created via sqlite3 will be stored on your $(pwd)/db

Development
------------
Once you have local docker environment setup. Development can be done locally
as well via Makefiles

```
make run
make test
```

### Tidying up dependencies

```
go mod tidy
```

### Build Docker image
```

make build
```
### Run Docker image
```
docker run -it -v $(pwd)/db:/app/db -p 8080:8080 quay.io/kenichi_shibata/golang-http-test:<GIT_HASH>
```
### Registry
```
docker push quay.io/kenichi_shibata/golang-http-test:<GIT_HASH>
```

Prerequisites
------------
* go version 1.12+
* make
* docker
* kubectl
* helm

Deployment
-----------------
Deploy Remotely via Kops
```
make deploy TYPE=kubernetes
```

Deploy Remotely via EKS
```
make deploy TYPE=eks
```

Changing DB
--------------------
By default the application uses sqlite file created as `users.db` on the local mounts. If you want to use a remote SQL DB like RDS. Please specify this in the environment variables like
```
docker run -it kenichishibata/golang-http-test --env-file
```
Where your env file will have
```
POSTGRES_ENV_POSTGRES_PASSWORD='foo'
POSTGRES_ENV_POSTGRES_USER='bar'
POSTGRES_ENV_DB_NAME='mysite_staging'
POSTGRES_PORT_5432_TCP_ADDR='docker-db-1.hidden.us-east-1.rds.amazonaws.com'
```

The DBs need to created with the following SQL Statement
```
CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, birthdate TEXT)
```

Deployment and Rolling Updates
---------------------
The recommendation is to separate the stateful DB from the Application. As such having the App on a Kubernetes cluster and the DB on RDS or other stateful provider makes sense. This is to loosely couple the application from the DB. However we have to keep in mind any schema changes and makes sure that any newer updates are backwards compatible and easily understand by well known frameworks such as semver.

Here is an example of how you would do a Rolling Update for the application with zero downtime
```
apiVersion:
kind:
```

Helm Chart Deployment
------------

Architecture on AWS Deployment using Kops
======================
