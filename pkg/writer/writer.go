package writer

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func Write(data []int, filename string) error {
	if _, err := os.Stat("output/"); os.IsNotExist(err) {
		err = os.Mkdir("output", 0755)
		if err != nil {
			return err
		}
	}

	file, err := os.Create("output/" + filename + strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
	if err != nil {
		return err
	}

	_, err = file.WriteString(strconv.FormatInt(int64(len(data)), 10) + "\n")
	if err != nil {
		return err
	}

	for _, row := range data {
		_, err = file.WriteString((strconv.FormatInt(int64(row), 10)) + " ")
		if err != nil {
			return err
		}
	}

	return err
}

func Zip() error {
	file, err := os.Create("output/source" + strconv.FormatInt(time.Now().UTC().UnixNano(), 10) + ".zip")
	if err != nil {
		return err
	}

	defer file.Close()

	w := zip.NewWriter(file)

	err = addFiles(w, "./", "")
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return err
}

func addFiles(w *zip.Writer, basePath, baseInZip string) error {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				return err
			}

			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				return err
			}
		} else if file.IsDir() {
			newBase := basePath + file.Name() + "/"
			err := addFiles(w, newBase, baseInZip  + file.Name() + "/")
			if err != nil {
				return err
			}
		}
	}

	return err
}