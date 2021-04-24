# Go Md5 Matcher

Md5 matcher tool using Golang.

## 1. Run Commands

Start application by running the following command:

``run ./cmd/console/main.go``

## 2. Configurations

Configurations must be loaded by setting values at ``config.toml`` in the main project path.

| Parameter | Description | Type | Default | Required |
|:---|:---|:---|:---|:---|
| ``encoded_column`` | Encoded column field to be added into the output file. | `int` | `-1` | **NO** |
| ``output_name`` | Output file name with match results. | `int` | `<current_timestamp>.csv` | **NO** |
| ``[data_csv]`` | Add a ``CsvFile`` config concerning the following topic. | `CsvFile` | ` ` | **YES** |
| ``[data_encoded]`` | Add a ``CsvFile`` config concerning the following topic. | `CsvFile` | ` ` | **YES** |

### 2.1. Csv File

| Parameter | Description | Type | Default | Required |
|:---|:---|:---|:---|:---|
| ``field_delimiter`` | Field character that delimits row fields. Usually is `;` or `,` . | `string` | `;` | **NO** |
| ``file_path`` | Csv file path. | `string` | ` ` | **YES** |
| ``match_column`` | Csv file column that will be used as match comparison. Counting starts with `0.` | `int` | `0` | **NO** |
| ``start_line`` | Csv line number that will be used to start comparison. Counting starts with `0.` | `int` | `0` | **NO** |

## 3. Guidelines

Instantiate ``CsvFile`` structs for the data CSV file and for the CSV file with the encoded fields.

```
matcher.CsvFile{
    Delimiter: ';',
    FilePath: 'data.csv',
    MatchColumn: 3,
}
```

### 3.1. Instantiate Csv Matcher

Instantiate a Csv Matcher by running ``NewCsvMatcher`` method.
Insert two ``CsvFile`` structs: one for Csv data file; and the other one for the encoded data Csv.

```
matcher.NewCsvMatcher(
    matcher.CsvFile{
        Delimiter: ';',
        FilePath: "data.csv",
        MatchColumn: 3,
        StartLine: 1,
    },
    matcher.CsvFile{
        Delimiter: ';',
        FilePath: "md5.csv",
        MatchColumn: 2,
        StartLine: 1,
    },
    "fileName.csv",
    0
)
```

### 3.2. Run Matcher

After instantiate a ``CsvMatcher`` struct, you may run ``Match()`` method as in the following excerpt.

```
matchStruct := matcher.NewCsvMatcher( [...] )

err := matchstruct.Match()
```

It will generate a new Csv file with the data file rows that matched with the encoded ones.
