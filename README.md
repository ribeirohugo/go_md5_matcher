# Go Md5 Matcher

Md5 matcher tool using Golang.

## 1. Run Commands

Start application by running the following command:

``run ./cmd/console/main.go``

## 2. Initialize CSV matcher

Instantiate a CSV Matcher by running NewCsvMatcher method.

```
matcher.NewCsvMatcher(
    "data.csv",     // data file path
    3,              // data file column to be compared
    ';'             // data file delimiter
    "md5.csv",      // encoded file path
    2,              // encoded file column to be compared
    ';'             // encoded file delimiter
)
```
