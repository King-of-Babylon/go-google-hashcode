package loader

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Load(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if file == nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	config := loadConfig(reader, file)
	lines := loadData(reader)

	return config, lines
}

func LoadFilesToProcess() []string {
	files, err := ioutil.ReadDir("./input")
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}

func loadConfig(reader *bufio.Reader, file *os.File) []string {
	config, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	config = strings.Replace(config, "\n", "", -1)

	if config == "" {
		log.Fatal("Could not read config from file: " + file.Name())
	}
	configArr := strings.Split(config, " ")

	return configArr
}

func loadData(reader *bufio.Reader) []string {
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
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

	return lines
}