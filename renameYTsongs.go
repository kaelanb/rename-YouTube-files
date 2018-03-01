package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "/your/directory"
	//filenames to skip if any
	skip := []string{"blablabla.txt", "blabla2.txt"}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//check filename to see if it needs to be skipped
		fname := info.Name()
		if fname == skip[0] || fname == skip[1] || info.IsDir() {
			return nil
		}

		//split filename from full path name
		dir, file := filepath.Split(path)
		n := len(file) - 16

		//strip 16 characters and append mp3
		file2 := file[:n] + ".mp3"
		newpath := dir + file2		
		os.Rename(path, newpath)
		
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
	}
}
