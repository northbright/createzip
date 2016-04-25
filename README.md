# createzip

createzip is a [Golang](http://golang.org) package which creates local zip or downloadable zip for HTTP server then add files into the zip file.

[![Build Status](https://travis-ci.org/northbright/createzip.svg?branch=master)](https://travis-ci.org/northbright/createzip)

#### Example to Create the Zip for HTTP Server  

You may find the [example](./example/main.go) in [./example/main.go](./example/main.go)

    func hello(w http.ResponseWriter, r *http.Request) {
        // Create zip file for HTTP server
        zf := createzip.NewForHTTP(w, "new.zip")
        // Add file into zip file. 1st param is the real file path, 2nd param is the file name in the zip file.
        // If file name in the zip file is empty, it will put the file in the root dir of the zip.
        zf.AddFile("./README.md", "")
        zf.AddFile("./main.go", "example/main.go")
        zf.Close()
    }

#### Run Example

`./example`


#### Documentation
* [API Reference](http://godoc.org/github.com/northbright/createzip)

#### License
* [MIT License](./LICENSE)
