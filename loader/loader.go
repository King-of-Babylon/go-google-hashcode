package loader

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func Load(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if file == nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	config := loadConfig(reader)
	lines := loadData(reader)

	return config, lines
}

func LoadFilesToProcess() []string {
	files, _ := ioutil.ReadDir("./src")

	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}

func loadConfig(reader *bufio.Reader) []string {
	config, _ := reader.ReadString('\n')
	config = strings.Replace(config, "\n", "", -1)
	configArr := strings.Split(config, " ")

	return configArr
}

func loadData(reader *bufio.Reader) []string {
	lines := []string{}
	for {
		line, err := reader.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)

		if line == "" {
			break
		}

		columns := strings.Split(line, " ")
		for _, element := range columns {
			lines = append(lines, element)
		}

		if err != nil {
			break
		}
	}

	return lines
}