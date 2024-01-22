# malignment-go
**FORK FROM** https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/fieldalignment   
Suport anonymous structure！

malignment-go is a tool designed to address memory alignment issues in Go language structs. With this tool, you can quickly identify and adjust the memory alignment of structs to improve memory utilization.

### Features
Struct Memory Alignment Fix: malignment-go can detect and fix memory alignment issues in Go language structs, ensuring that struct fields are aligned in the most optimal way to enhance memory utilization.

Simple and Easy to Use: Simply use the command malignment-go -fix . to perform memory alignment fixes on all Go files in the specified directory. No manual intervention is required as the tool automatically handles the repair process.

Usage
```bash
malignment-go -fix .
// or
go run -mod=mod github.com/xsean2020/malignment-go -fix .
```
The above command will recursively fix memory alignment issues in all Go files within the current directory.

### Example
```go
type A struct {
        C int64
        D struct {
                Bb int16
                Cc int32
                AA int8
        }
        Array []int32
        Age   int16
        Map   map[string]int
        Bool  bool
}

```

```
+----+--------------------------------+-----------+-----------+
| ID |           FIELDTYPE            | FIELDNAME | FIELDSIZE |
+----+--------------------------------+-----------+-----------+
| A  | int64                          | C         | 8         |
| B  | struct { Bb int16; Cc int32;   | D         | 12        |
|    | AA int8 }                      |           |           |
| C  | []int32                        | Array     | 24        |
| D  | int16                          | Age       | 2         |
| E  | map[string]int                 | Map       | 8         |
| F  | bool                           | Bool      | 1         |
+----+--------------------------------+-----------+-----------+
---- Memory layout ----
|A|A|A|A|A|A|A|A|
|B|B|B|B|B|B|B|B|
|B|B|B|B| | | | |
|C|C|C|C|C|C|C|C|
|C|C|C|C|C|C|C|C|
|C|C|C|C|C|C|C|C|
|D|D| | | | | | |
|E|E|E|E|E|E|E|E|
|F|
```
After run command:
```bash
malignment-go -fix .

```

```
+----+--------------------------------+-----------+-----------+
| ID |           FIELDTYPE            | FIELDNAME | FIELDSIZE |
+----+--------------------------------+-----------+-----------+
| A  | map[string]int                 | Map       | 8         |
| B  | []int32                        | Array     | 24        |
| C  | int64                          | C         | 8         |
| D  | struct { Cc int32; Bb int16;   | D         | 8         |
|    | AA int8 }                      |           |           |
| E  | int16                          | Age       | 2         |
| F  | bool                           | Bool      | 1         |
+----+--------------------------------+-----------+-----------+
---- Memory layout ----
|A|A|A|A|A|A|A|A|
|B|B|B|B|B|B|B|B|
|B|B|B|B|B|B|B|B|
|B|B|B|B|B|B|B|B|
|C|C|C|C|C|C|C|C|
|D|D|D|D|D|D|D|D|
|E|E|F|
```

### Notes
Before using malignment-go, make sure to back up your code in advance to prevent any unexpected issues during the repair process.

Familiarize yourself with the relevant knowledge of memory alignment before proceeding with the repair to ensure that it aligns with your requirements.

### Contribution
Feel free to raise issues, report bugs, or contribute code. Please check the contribution guidelines before submitting issues or requests.
