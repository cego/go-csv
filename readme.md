Golang package for reading and writing CSV data
===============================================

Golang package for reading and writing CSV data as described in [RFC 4180](https://tools.ietf.org/html/rfc4180).

Is was written as a replacement of the stdlib package encoding/csv in cases where that interpretation of the RFC is
lacking. Specifically, we want to preserve carriage returns inside fields.

To keep this package simple, it does not attempt to implement all features of encoding/csv and be a complete drop-in
replacement. You will likely need to adjust your code slightly to use this package instead of encoding/csv.

Examples
--------

Reading CSV formatted data.

    import (
        "strings"
        "github.com/cego/go-csv"
    )

    reader := csv.NewReader(strings.NewReader("fi,fy\nfo,fum"))
    records, err := reader.ReadAll()
    if err != nil && err != io.EOF {
        panic(err)
    }

    // records is now [ ["fi","fy"], ["fo","fum"] ]

Writing CSV formatted data.

    import (
        "strings"
        "github.com/cego/go-csv"
    )

    builder := strings.Builder{}
    writer := csv.NewWriter(&builder)
    err := writer.WriteAll([][]string{ {"fi", "fy"}, {"fo", "fum"} })
    if err != nil {
        panic(err)
    }

    // builder.String() is now "fi","fy"\r\n"fo","fum"
