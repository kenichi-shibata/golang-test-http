Golang-Http-test
============

Testing Golang's Http capabilities without HTTP Packages

Only uses packages for sqlite3 stateful db for local dev.

Dependencies
------------
* Glog - Google's logging used by Kubernetes
* Go-sqlite3 - Local DB for lightweight sql

Running test
---------------
```
make test
```


Development
------------
Local Run
```
make run
```

Tidying up dependencies
```
go mod tidy
```

Docker
---------
Build
```

make build
```
Run
```
docker run -it kenichishibata/golang-http-test:<GIT_HASH> 
```


Deploy Remotely
```
make deploy TYPE=kubernetes
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

Deployment and Rolling Updates 
---------------------
The recommendation is to separate the stateful DB from the Application. As such having the App on a Kubernetes cluster and the DB on RDS or other stateful provider makes sense. This is to loosely couple the application from the DB. However we have to keep in mind any schema changes and makes sure that any newer updates are backwards compatible and easily understand by well known frameworks such as semver. 

Here is an example of how you would do a Rolling Update for the application with zero downtime
```
apiVersion:
kind: 
```

Architecture on AWS Deployment using Kops
======================