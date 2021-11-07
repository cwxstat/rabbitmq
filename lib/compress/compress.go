package compress

import (
	"archive/tar"
	"compress/gzip"

	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Compress used on complete directory
func Compress(sourceDir string, removePrefix string) error {

	destinationFile := sourceDir + ".tar.gz"

	tarfile, err := os.Create(destinationFile)

	if err != nil {
		return err
	}

	defer tarfile.Close()
	var fileWriter io.WriteCloser = tarfile

	if strings.HasSuffix(destinationFile, ".gz") {
		fileWriter = gzip.NewWriter(tarfile) // add a gzip filter
		defer fileWriter.Close()             // if user add .gz in the destination filename
	}

	tarfileWriter := tar.NewWriter(fileWriter)
	defer tarfileWriter.Close()

	loop(sourceDir, tarfileWriter, removePrefix)

	return nil
}

func UnCompress(sourceFile, destinationDir string) error {

	file, err := os.Open(sourceFile)
	if err != nil {
		return err
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(sourceFile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {

			return err

		}
		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files

	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err

		}

		// get the individual filename and extract to the current directory
		//err = os.MkdirAll(destinationDir, os.FileMode(header.Mode))
		filename := destinationDir + header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			fmt.Println("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				return err

			}

		case tar.TypeReg:
			err = os.MkdirAll(filepath.Dir(filename), os.FileMode(0755))
			if err != nil {
				return err
			}
			// handle normal file
			fmt.Println("Untarring :", filename)
			writer, err := os.Create(filename)

			if err != nil {
				return err

			}

			_, err = io.Copy(writer, tarBallReader)
			if err != nil {
				return err
			}

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				return err

			}

			writer.Close()
		default:
			fmt.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
		}
	}

	return nil
}

func loop(sourceDir string, tarfileWriter *tar.Writer, prefixDirToRemove string) error {

	dir, err := os.Open(sourceDir)

	if err != nil {
		return err
	}

	defer dir.Close()

	files, err := dir.Readdir(0) // grab the files list

	if err != nil {
		return err
	}

	for _, fileInfo := range files {

		if fileInfo.IsDir() {

			err = loop(sourceDir+"/"+fileInfo.Name(), tarfileWriter, prefixDirToRemove)
			if err != nil {
				return err
			}
			continue
		}

		file, err := os.Open(dir.Name() + string(filepath.Separator) + fileInfo.Name())

		if err != nil {
			return err
		}

		defer file.Close()
		// prepare the tar header

		header := new(tar.Header)
		header.Name = file.Name()
		if strings.Index(header.Name, prefixDirToRemove) == 0 {
			header.Name = strings.Replace(header.Name, prefixDirToRemove, "", 1)
		}
		header.Size = fileInfo.Size()
		header.Mode = int64(fileInfo.Mode())
		header.ModTime = fileInfo.ModTime()

		err = tarfileWriter.WriteHeader(header)

		if err != nil {
			return err
		}

		_, err = io.Copy(tarfileWriter, file)

		if err != nil {
			return err
		}

	}
	return nil

}
