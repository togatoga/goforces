# goforces
[![Build Status](https://travis-ci.org/togatoga/goforces.svg?branch=master)](https://travis-ci.org/togatoga/goforces) [![GoDoc](https://godoc.org/github.com/togatoga/goforces?status.svg)](https://godoc.org/github.com/togatoga/goforces)

goforces is go package for the codeforces(http://codeforces.com/) api.  
Check the [Usage](#usage) and the [codeforces api help](http://codeforces.com/api/help)

## Install
```
go get github.com/togatoga/goforces
```

## Usage


```go
ctx := context.Background()

//Codeforces client
api, err := goforces.NewClient(nil)
if err != nil {
  panic(err)
}

options := map[string]interface{}{
  "tags": []string{"dp"},
}
//Problems
problems, err := api.GetProblemSetProblems(ctx, options)
if err != nil {
  panic(err)
}
fmt.Printf("%+v\n", problems)

//Contest list
contestList, err := api.GetContestList(ctx, nil)
if err != nil {
  panic(err)
}
fmt.Printf("%+v\n", contestList)

//If you use authorized methods, you must set your key and secret
api.SetApiKey("<your key>")
api.SetApiSecret("<your secret>")
//User friends
friends, err := api.GetUserFriends(ctx, nil)
if err != nil {
  panic(err)
}
fmt.Printf("%+v\n", friends)

```
## Documention
Read [Godoc](https://godoc.org/github.com/togatoga/goforces)

## License

[MIT License](LICENSE)
