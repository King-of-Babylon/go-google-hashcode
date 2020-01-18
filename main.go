package main

import (
    "github.com/King-of-Babylon/go-google-hashcode/pkg/loader"
    "github.com/King-of-Babylon/go-google-hashcode/pkg/processor"
    "github.com/King-of-Babylon/go-google-hashcode/pkg/writer"
)

func main() {
    filenames := loader.LoadFilesToProcess()

    for _, filename := range filenames {
        config, data := loader.Load("./input/" + filename)
        result := processor.Process(config, data)
        writer.Write(result, filename)
    }

    writer.Zip()
}