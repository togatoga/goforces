# goforces

[![Build Status](https://travis-ci.org/togatoga/goforces.svg?branch=master)](https://travis-ci.org/togatoga/goforces) [![GoDoc](https://godoc.org/github.com/togatoga/goforces?status.svg)](https://godoc.org/github.com/togatoga/goforces)[![Go Report Card](https://goreportcard.com/badge/github.com/togatoga/goforces)](https://goreportcard.com/report/github.com/togatoga/goforces)

goforces is go package for the codeforces(<http://codeforces.com/>) api.
Check the [Usage](#usage)

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
    //logger
    logger := log.New(os.Stderr, "*** ", log.LstdFlags)
    //Codeforces client
    api, _ := goforces.NewClient(logger)
    //Problems
    problems, _ := api.GetProblemSetProblems(ctx, &goforces.ProblemSetProblemsOptions{Tags: []string{"dp"}})
    fmt.Printf("%+v\n", problems)

    //Contest list
    contestList, _ := api.GetContestList(ctx, nil)
    fmt.Printf("%+v\n", contestList)

    //If you use authorized methods, you must set your key and secret
    api.SetAPIKey("<your key>")
    api.SetAPISecret("<your secret>")
    //User friends
    friends, _ := api.GetUserFriends(ctx, nil)
    fmt.Printf("%+v\n", friends)
}
```

The official codeforces api documentation is [here](http://codeforces.com/api/help)

## Documention

Read [Godoc](https://godoc.org/github.com/togatoga/goforces)

## License

[MIT License](LICENSE)
