# Go Md5 Matcher

Md5 matcher tool using Golang.

## 1. Run Commands

Start application by running the following command:

``run ./cmd/console/main.go``

## 2. Guidelines

Instantiate ``CsvFile`` structs for the data CSV file and for the CSV file with the encoded fields.

```
matcher.CsvFile{
    Delimiter: ';',
    FilePath: 'data.csv',
    MatchColumn: 3,
}
```

### 2.1. Instantiate Csv Matcher

Instantiate a Csv Matcher by running ``NewCsvMatcher`` method.
Insert two ``CsvFile`` structs: one for 

```
matcher.NewCsvMatcher(
    matcher.CsvFile{
        Delimiter: ';',
        FilePath: 'data.csv',
        MatchColumn: 3,
    },
    matcher.CsvFile{
        Delimiter: ';',
        FilePath: 'md5.csv',
        MatchColumn: 2,
    }
)
```

### 2.2. Run Matcher

After instantiate a ``CsvMatcher`` struct, you may run ``Match()`` method as in the following excerpt.

```
matchStruct := matcher.NewCsvMatcher( [...] )

err := matchstruct.Match()
```
