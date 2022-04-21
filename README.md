# Jsonvalue - A Go package for fast and convient JSON marshal and unmarshaling

[![Workflow](https://github.com/Andrew-M-C/go.jsonvalue/actions/workflows/go_test_general.yml/badge.svg)](https://github.com/Andrew-M-C/go.jsonvalue/actions/workflows/go_test_general.yml)
[![Coveralls](https://coveralls.io/repos/github/Andrew-M-C/go.jsonvalue/badge.svg?branch=master)](https://coveralls.io/github/Andrew-M-C/go.jsonvalue)
[![Go report](https://goreportcard.com/badge/github.com/Andrew-M-C/go.jsonvalue)](https://goreportcard.com/report/github.com/Andrew-M-C/go.jsonvalue)
[![Codebeat](https://codebeat.co/badges/ecf87760-2987-48a7-a6dd-4d9fcad57256)](https://codebeat.co/projects/github-com-andrew-m-c-go-jsonvalue-master)<br>
[![GoDoc](https://godoc.org/github.com/Andrew-M-C/go.jsonvalue?status.svg)](https://pkg.go.dev/github.com/Andrew-M-C/go.jsonvalue@v1.2.0)
[![Latest](https://img.shields.io/badge/latest-v1.2.0-blue.svg)](https://github.com/Andrew-M-C/go.jsonvalue/tree/v1.2.0)
[![Beta](https://img.shields.io/badge/beta-v1.2.1-yellow.svg)](https://github.com/Andrew-M-C/go.jsonvalue/tree/feature/v1.2.1)
[![License](https://img.shields.io/badge/license-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

[中文版](./README_cn.md)，[中文版详细文档](./docs/zh-cn/README.md)

**Jsonvalue** is a Golang package for JSON parsing. It is used in situations those Go structures cannot achieve, or `map[string]interface{}` could not do properbally.

## Import

Use following statements to import jsonvalue:

```go
import (
    jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
)
```

## Quick Start

Sometimes we want to create a complex JSON object like:

```json
{
    "someObject": {
        "someObject": {
            "someObject": {
                "message": "Hello, JSON!"
            }
        }
    }
}
```

With `jsonvalue`, It is quite simple to achieve this:

```go
    v := jsonvalue.NewObject()
    v.SetString("Hello, JSON").At("someObject", "someObject", "someObject", "message")
    fmt.Println(v.MustMarshalString())
    // Output:
    // {"someObject":{"someObject":{"someObject":{"message":"Hello, JSON!"}}}
```

Similarly, it is quite easy to create sub-arrays like:

```json
[
    {
        "someArray": [
            "Hello, JSON!"
        ]
    }
]
```

```go
    v := jsonvalue.NewArray()
    v.SetString("Hello, JSON").At(0, "someObject", 0)
    fmt.Println(v.MustMarshalString())
    // Output:
    // [{"someObject":["Hello, JSON"]}]
```

In opposite, to parse and read the first JSON above, you can use jsonvalue like this:

```go
	const raw = `{"someObject": {"someObject": {"someObject": {"message": "Hello, JSON!"}}}}`
	s := jsonvalue.MustUnmarshalString(s).GetString("someObject", "someObject", "someObject", "message")
	fmt.Println(v.MustMarshalString())
	// Output:
	// Hello, JSON!
```

However, it is quite complex and annoying in automatically creating array. I strongly suggest using `SetArray()` to create the array first, then use `Append()` or `Insert()` to set array elements. Please refer go [godoc](https://godoc.org/github.com/Andrew-M-C/go.jsonvalue).
