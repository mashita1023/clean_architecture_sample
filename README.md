# clean_architecture_sample_app

# use it
```
$ docker compose build
$ docker compose run --rm gin_api go get -u
$ docker compose up
```

# get go modules
```
$ docker compose run --rm gin_api go get -u [modules]
$ docker compose run --rm gin_api go mod tidy
