# Iris App

This is a quick sample with Iris Web Framework. (Golang)

## How to use

1. clone repo
```
git clone https://github.com/kilfu0701/iris-app.git
```

2. set `GOPATH`
```
cd iris-app
export GOPATH=`pwd`
```

3. install `dep`
```
brew install dep
# or
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

4. install vendor packages
```
cd iris-app/src/app
dep ensure
```

5. run local server and check http://localhost:9000
```
make serve
```

## Lincense

MIT
