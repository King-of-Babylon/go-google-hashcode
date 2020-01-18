package loader

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Load(filename string) ([]string, []string, error) {
	var config, lines []string

	file, err := os.Open(filename)
	if err != nil {
		return config, lines, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	config, err = loadConfig(reader)
	if err != nil {
		return config, lines, err
	}

	lines, err = loadData(reader)

	return config, lines, err
}

func LoadFilesToProcess() ([]string, error) {
	var fileNames []string

	files, err := ioutil.ReadDir("./input")
	if err != nil {
		return fileNames, err
	}

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, err
}

func loadConfig(reader *bufio.Reader) ([]string, error) {
	var configArr []string

	config, err := reader.ReadString('\n')
	if err != nil {
		return configArr, err
	}

	config = strings.Replace(config, "\n", "", -1)
	if config == "" {
		return configArr, errors.New("config can not be empty")
	}
	configArr = strings.Split(config, " ")

	return configArr, err
}

func loadData(reader *bufio.Reader) ([]string, error) {
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return lines, err
		}

		line = strings.Replace(line, "\n", "", -1)
		if line == "" {
			break
		}

		columns := strings.Split(line, " ")
		for _, element := range columns {
			lines = append(lines, element)
		}
	}

	return lines, nil
}