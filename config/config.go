package config

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
	"github.com/pkg/errors"
)

type Config struct {
	Data DataBases `hcl:"data"`
}

type DataBases struct {
	DataUser     string `hcl:"user"`
	DataPassword string `hcl:"password"`
	DataDB       string `hcl:"database"`
}

func ParseConfigFile(filename string, target interface{}) (err error) {
	defer func() { err = errors.Wrap(err, "config.ParseConfigFile") }()
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		err = errors.Wrap(err, "1234")
		return
	}
	fmt.Print("dskfjs")
	err = hcl.Unmarshal(b, target)
	if err != nil {
		err = errors.Wrap(err, "faild unmarshal")
	}
	return
}
