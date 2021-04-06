package tm

import "os"

func LoadYamlFile(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return LoadYamlReader(file)
}
