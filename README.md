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
go run -mod=mod https://github.com/xsean2020/malignment-go -fix .
```
The above command will recursively fix memory alignment issues in all Go files within the current directory.


### Notes
Before using malignment-go, make sure to back up your code in advance to prevent any unexpected issues during the repair process.

Familiarize yourself with the relevant knowledge of memory alignment before proceeding with the repair to ensure that it aligns with your requirements.

### Contribution
Feel free to raise issues, report bugs, or contribute code. Please check the contribution guidelines before submitting issues or requests.
