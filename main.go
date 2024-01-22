package main

import (
	"github.com/xsean2020/malignment-go/malignment"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(malignment.Analyzer)
}
