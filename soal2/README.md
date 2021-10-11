# Stockbit Microservices Test

### Step to installation

Go to root soal2 directory

```
$ git mod tidy
```

### Run Mysql Docker Container

Go to root soal2 directory

```
$ bash bin/docker_run.sh
```

### Generate Proto(Additional)

Go to root soal2 directory

```
$ bash bin/generate_movie_proto.sh
```

### Run REST and RPC Service

Go to root soal2 directory

open a new terminal, then following type:

```
$ bash bin/start_rpc_movie_service.sh
```

open a new terminal, then following type:
```
$ bash bin/start_movie_rest_service.sh
```

### How to test?

example endpoint is in a directory http

```
#example get movies by searchWord 
GET http://localhost:8080/v1/movie/search/Batman/2
Accept: application/json

###
# example get details movie by omdbId
GET http://localhost:8080/v1/movie/search/details/tt0096895
Accept: application/json

###
# example incorrect imdbId
GET http://localhost:8080/v1/movie/search/details/zzzzzzzzzzz
Accept: application/json

###

GET http://localhost:8080/check/health
Accept: application/json

```