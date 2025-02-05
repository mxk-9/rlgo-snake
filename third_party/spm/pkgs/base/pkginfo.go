package pkginfo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Package struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	UsePackageManager bool   `json:"use_package_manager"`
}

func NewPackageItem(pathToPackage string) (pkg *Package, err error) {
	var (
		filePackage *os.File
		buf         *bufio.Scanner
	)

	pkg = &Package{}

	if filePackage, err = os.Open(pathToPackage); err != nil {
		err = fmt.Errorf("Failed to open package's file:\n%s\n", err)
		return
	}
	defer filePackage.Close()

	buf = bufio.NewScanner(filePackage)

	if err = json.Unmarshal(buf.Bytes(), pkg); err != nil {
		err = fmt.Errorf("Failed to unmarshal package:\n%s\n", err)
		return
	}

	return
}
