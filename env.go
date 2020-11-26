package go_env

import (
	"io/ioutil"
	"strings"
)

type EnvParams struct {
	FilePath string
	Results  map[string][]map[string]string
}

func (env *EnvParams) ParseEnvFile(fileName string) {
	buf, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	content := string(buf)
	lines := strings.Split(content, "\r\n")
	result := make(map[string][]map[string]string, 0)
	var moduleName string
	moduleMap := make([]map[string]string, 0)
	for i := 0; i < len(lines); i++ {

		thisModule := make(map[string]string, 0)
		line := strings.Replace(lines[i], "\r", "", -1)
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// has module name
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			module := line[1 : len(line)-1]
			if module == "" {
				continue
			}
			moduleName = strings.ToLower(module)
			moduleMap = nil
			continue
		}

		pos := strings.Index(line, "=")
		if pos > 0 && pos < len(line)-1 {
			key := line[:pos]
			val := line[pos+1:]

			key = strings.TrimSpace(strings.ToLower(key))
			val = strings.TrimSpace(strings.TrimSpace(val))

			thisModule[key] = val
		}

		moduleMap = append(moduleMap, thisModule)
		result[moduleName] = moduleMap
	}
	env.Results = result
}

func (env *EnvParams) GetSection(moduleName, keyName string) string {
	moduleName = strings.ToLower(strings.TrimSpace(moduleName))
	keyName = strings.ToLower(strings.TrimSpace(keyName))
	config := env.Results
	_, ok := config[moduleName]
	if !ok {
		panic("Invalid module name")
	}
	module := config[moduleName]
	for _, m := range module {
		for key, name := range m {
			if key == keyName {
				return name
			}
		}
	}
	return ""
}
