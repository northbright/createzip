# Example to Create Downloadable Zip for HTTP Server

#### Example to Create the Zip for HTTP Server

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

`sudo ./example`
