package writer

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func Write(data []int, filename string) {
	file, err := os.Create("output/" + filename + strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(strconv.FormatInt(int64(len(data)), 10) + "\n")
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range data {
		_, err = file.WriteString((strconv.FormatInt(int64(row), 10)) + " ")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Zip() {
	file, err := os.Create("output/source" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + ".zip")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	w := zip.NewWriter(file)

	addFiles(w, "./", "")

	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				log.Fatal(err)
			}

			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				log.Fatal(err)
			}
		} else if file.IsDir() {
			newBase := basePath + file.Name() + "/"
			addFiles(w, newBase, baseInZip  + file.Name() + "/")
		}
	}
}