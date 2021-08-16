/*
@Author :
@Group :Rosefintech
@Description :
@Data: 2021/06/18 14:09
*/

package tools

import (
	"gopkg.in/yaml.v2"
	"os"
	"rosefintech-rosl2/common"
)

func ReadYamlConfig(path string) error {
	common.Conf = new(common.Config)
	if f, err := os.Open(path); err != nil {
		return err
	} else {
		yaml.NewDecoder(f).Decode(common.Conf)
	}
	return nil
}

// Init init config.
func init() {
	err := ReadYamlConfig("./conf.yml")
	if err != nil {
		panic(err.Error())
	}
	common.Conf.Chain.PrivateKey = ""

	return
}
