package utils

import (
	"io/fs"
	"os"
	"strings"
)

func UpdateEnvFile(envFilePath, key, value string) error {
	content, err := os.ReadFile(envFilePath)
	if err != nil {
		return err
	}

	envFileContent := string(content)

	lines := strings.Split(envFileContent, "\n")

	envMap := make(map[string]string)

	for _, line := range lines {
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}

	envMap[key] = value

	var updatedContent []string
	for k, v := range envMap {
		updatedContent = append(updatedContent, k+"="+v)
	}

	newContent := strings.Join(updatedContent, "\n")

	err = os.WriteFile(envFilePath, []byte(newContent), fs.ModeAppend)
	if err != nil {
		return err
	}

	return nil
}
