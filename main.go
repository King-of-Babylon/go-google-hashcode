package main

import (
    "github.com/King-of-Babylon/hashcode2020/loader"
    "github.com/King-of-Babylon/hashcode2020/processor"
    "github.com/King-of-Babylon/hashcode2020/writer"
)

func main() {
    filenames := loader.LoadFilesToProcess()

    for _, filename := range filenames {
        config, data := loader.Load("./src/" + filename)
        result := processor.Process(config, data)
        writer.Write(result, filename)
    }

    writer.Zip()
}