package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

const (
	debugFlag = "JIKE_SDK_DEBUG"
)

func ShowInformation(name string, info ...interface{}) {
	fmt.Printf("=== %s ===\n", name)
	spew.Dump(info...)
	fmt.Println("")
}

func IsDebug() bool {
	return strings.ToLower(os.Getenv(debugFlag)) == "true"
}
