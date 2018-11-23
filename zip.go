package main

import (
    "archive/zip"
    "io"
    "log"
    "os"
)

func main() {

    // Files to Zip
    files := []string{"temp.go","main.go", "output.txt","users.json","zip.go"}
    output := "done.zip"

    err := ZipFiles(output, files)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Zipped File: " + output)
}

// ZipFiles compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
func ZipFiles(filename string, files []string) error {

    newZipFile, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer newZipFile.Close()

    zipWriter := zip.NewWriter(newZipFile)
    defer zipWriter.Close()

    // Add files to zip
    for _, file := range files {

        zipfile, err := os.Open(file)
        if err != nil {
            return err
        }
        defer zipfile.Close()

        // Get the file information
        info, err := zipfile.Stat()
        if err != nil {
            return err
        }

        header, err := zip.FileInfoHeader(info)
        if err != nil {
            return err
        }

        // Using FileInfoHeader() above only uses the basename of the file. If we want 
        // to preserve the folder structure we can overwrite this with the full path.
        header.Name = file

        // Change to deflate to gain better compression
        // see http://golang.org/pkg/archive/zip/#pkg-constants
        header.Method = zip.Deflate

        writer, err := zipWriter.CreateHeader(header)
        if err != nil {
            return err
        }
        if _, err = io.Copy(writer, zipfile); err != nil {
            return err
        }
    }
    return nil
}