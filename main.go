package main

import (
	"fmt"

	"github.com/prometheus/prometheus/pkg/rulefmt"
	"github.com/prometheus/prometheus/config"
)

func main() {
	_, errs := rulefmt.ParseFile("")
	myconf := config.Config{}

	fmt.Printf("Errs: %+v", errs)
	fmt.Printf("Conf: %+v", myconf)
}
