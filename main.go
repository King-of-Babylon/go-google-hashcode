package main

import (
    "github.com/King-of-Babylon/go-google-hashcode/pkg/loader"
    "github.com/King-of-Babylon/go-google-hashcode/pkg/processor"
    "github.com/King-of-Babylon/go-google-hashcode/pkg/writer"
)

func main() {
    filenames, err := loader.LoadFilesToProcess()
    handleError(err)

    for _, filename := range filenames {
        config, data, err := loader.Load("./input/" + filename)
        handleError(err)

        result := processor.Process(config, data)
        err = writer.Write(result, filename)
        handleError(err)
    }

    err = writer.Zip()
    handleError(err)
}

func handleError(err error) {
    if err != nil {
        panic(err)
    }
}