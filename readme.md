## GO-lang revision + GitHub history 

### To run in docker env use or create alias
Pull latest goLang image

```
docker run --rm -it -v "$(pwd)":/app -w /app golang:1.22 go run main.go
```

alias
```
alias try='docker run --rm -it -v "$(pwd)":/app -w /app golang:1.22 go run main.go'
```