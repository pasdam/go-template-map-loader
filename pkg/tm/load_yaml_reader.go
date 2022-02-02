package tm

import (
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func LoadYamlReader(reader io.Reader) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	currentMap := map[string]interface{}{}
	if err := yaml.Unmarshal(bytes, &currentMap); err != nil {
		return nil, errors.Wrapf(err, "failed to parse yaml")
	}
	return currentMap, nil
}
