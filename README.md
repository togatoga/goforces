# goforces

[![Build Status](https://travis-ci.org/togatoga/goforces.svg?branch=master)](https://travis-ci.org/togatoga/goforces) [![GoDoc](https://godoc.org/github.com/togatoga/goforces?status.svg)](https://godoc.org/github.com/togatoga/goforces)

goforces is go package for the codeforces(<http://codeforces.com/>) api. Check the [Usage](#usage) and the [codeforces api help](http://codeforces.com/api/help)

## Install

```
go get github.com/togatoga/goforces
```

## Usage

```go
package main

import (
    "context"
    "fmt"

    "github.com/togatoga/goforces"
)

func main() {
    ctx := context.Background()

    //Codeforces client
    api, _ := goforces.NewClient(nil)
    options := map[string]interface{}{
        "tags": []string{"dp"},
    }
    //Problems
    problems, _ := api.GetProblemSetProblems(ctx, options)
    fmt.Printf("%+v\n", problems)

    //Contest list
    contestList, _ := api.GetContestList(ctx, nil)
    fmt.Printf("%+v\n", contestList)

    //If you use authorized methods, you must set your key and secret
    api.SetApiKey("<your key>")
    api.SetApiSecret("<your secret>")
    //User friends
    friends, _ := api.GetUserFriends(ctx, nil)
    fmt.Printf("%+v\n", friends)
}
```

## Documention

Read [Godoc](https://godoc.org/github.com/togatoga/goforces)

## License

[MIT License](LICENSE)
