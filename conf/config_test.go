package conf_test

import (
	"fmt"
	"os"
	"restful-api/conf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		fmt.Println((conf.C().App.Name))
	}
}
func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	os.Setenv("MySQl_DATABASE", "unit_test")
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("unit_test", conf.C().MySQL.Database)
		fmt.Println((conf.C().MySQL.Database))
	}
}

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		conf.C().MySQL.GetDB()
	}
}
