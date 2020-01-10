package main

import (
    "github.com/King-of-Babylon/hashcode2020/loader"
    "github.com/King-of-Babylon/hashcode2020/processor"
)

func main() {
    filenames := loader.LoadFilesToProcess()

    for _, filename := range filenames {
        config, data := loader.Load("./src/" + filename)
        processor.Process(config, data)
    }
}