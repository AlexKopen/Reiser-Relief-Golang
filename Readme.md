```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

```
go install github.com/cosmtrek/air@latest
```

```
export USELOCAL=true; air
```

```
npm install -g browser-sync
```

```
browser-sync start --proxy "localhost:8080" --files "**/*.*"
```