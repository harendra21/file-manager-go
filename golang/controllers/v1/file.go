package v1

import (
	"archive/zip"
	"beego/controllers"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FileController struct {
	controllers.AppController
}

type Files struct {
	Name  string `json:"name"`
	IsDir bool   `json:"is_dir"`
	Type  string `json:"type"`
	Size  int64  `json:"size"`
}

func (ctrl *FileController) AllFiles() {
	var parent string = ctrl.GetString("p")
	var dir string = "./data/" + parent + "/"
	entries, err := os.ReadDir(dir)
	if err != nil {
		ctrl.ThrowError(200, err.Error())
		return
	}
	var files []Files
	for _, e := range entries {
		var file Files
		file.Name = e.Name()
		file.IsDir = e.IsDir()

		if !e.IsDir() {
			openfile, err := os.Open(dir + e.Name())
			if err != nil {
				ctrl.ThrowError(200, err.Error())
				return
			}
			contentType, _ := ctrl.GetFileContentType(openfile)
			file.Type = contentType

			fi, err := openfile.Stat()
			if err != nil {
				ctrl.ThrowError(200, err.Error())
				return
			}
			file.Size = fi.Size()
			defer openfile.Close()

		} else {
			size, err := ctrl.DirSize(dir + e.Name())
			if err != nil {
				ctrl.ThrowError(200, err.Error())
				return
			}
			file.Size = size
		}
		files = append(files, file)
	}
	ctrl.Response(files)
}

func (ctrl *FileController) DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func (ctrl *FileController) GetFileContentType(ouput *os.File) (string, error) {

	// to sniff the content type only the first
	// 512 bytes are used.

	buf := make([]byte, 512)

	_, err := ouput.Read(buf)

	if err != nil {
		return "", err
	}

	// the function that actually does the trick
	contentType := http.DetectContentType(buf)

	return contentType, nil
}

func (ctrl *FileController) MoveRename() {
	var source string = ctrl.GetString("source")
	var destination string = ctrl.GetString("destination")
	if source == "" || destination == "" {
		ctrl.ThrowError(200, "Source and Destination are required")
		return
	}
	source = "./data/" + source
	destination = "./data/" + destination
	err := os.Rename(source, destination)

	if err != nil {
		ctrl.ThrowError(200, err.Error())
		return
	}
	ctrl.Response("Moved successfully")

}

func (ctrl *FileController) Copy() {
	var source string = ctrl.GetString("source")
	var destination string = ctrl.GetString("destination")
	if source == "" || destination == "" {
		ctrl.ThrowError(200, "Source and Destination are required")
		return
	}
	source = "./data/" + source
	destination = "./data/" + destination

	file, err := os.Open(source)
	if err != nil {
		ctrl.ThrowError(200, err.Error())
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		ctrl.ThrowError(200, err.Error())
		return
	}

	defer file.Close()

	if fileInfo.IsDir() {
		err := ctrl.CopyDirs(source, destination)
		if err != nil {
			ctrl.ThrowError(200, err.Error())
			return
		}
	} else {
		err := ctrl.CopyFiles(source, destination)
		if err != nil {
			ctrl.ThrowError(200, err.Error())
			return
		}
	}

	ctrl.Response("Copied successfully")
}

func (ctrl *FileController) CopyFiles(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func (ctrl *FileController) CopyDirs(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = ctrl.CopyDirs(srcfp, dstfp); err != nil {
				return err
			}
		} else {
			if err = ctrl.CopyFiles(srcfp, dstfp); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ctrl *FileController) Zip() {
	var source string = ctrl.GetString("source")
	var destination string = ctrl.GetString("destination")
	if source == "" || destination == "" {
		ctrl.ThrowError(200, "Source and Destination are required")
		return
	}
	source = "./data/" + source
	destination = "./data/" + destination

	err := ctrl.ZipFiles(source, destination)
	if err != nil {
		ctrl.ThrowError(200, err.Error())
		return
	}
	ctrl.Response("Zipped successfully")
}

func (ctrl *FileController) ZipFiles(source, target string) error {
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	// 2. Go through all the files of the source
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 3. Create a local file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// set compression
		header.Method = zip.Deflate

		// 4. Set relative path of a file as the header name
		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

func (ctrl *FileController) Unzip() {
	var source string = ctrl.GetString("source")
	var destination string = ctrl.GetString("destination")
	if source == "" || destination == "" {
		ctrl.ThrowError(200, "Source and Destination are required")
		return
	}
	source = "./data/" + source
	destination = "./data/" + destination

	// 1. Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		ctrl.ThrowError(200, err.Error())
		return
	}
	defer reader.Close()

	// 2. Get the absolute destination path
	destination, err = filepath.Abs(destination)
	if err != nil {
		ctrl.ThrowError(200, err.Error())
		return
	}

	// 3. Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {
		err := ctrl.UnzipFiles(f, destination)
		if err != nil {
			ctrl.ThrowError(200, err.Error())
			return
		}
	}

	ctrl.Response("Unzipped successfully")
}

func (ctrl *FileController) UnzipFiles(f *zip.File, destination string) error {
	// 4. Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// 5. Create directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// 6. Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// 7. Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}

func (ctrl *FileController) Delete() {
	var source string = ctrl.GetString("source")

	if source == "" {
		ctrl.ThrowError(200, "Source is required")
		return
	}
	source = "./data/" + source

	if err := os.RemoveAll(source); err != nil {
		ctrl.ThrowError(200, err.Error())
		return
	}

	ctrl.Response("Zipped successfully")
}

func (ctrl *FileController) Create() {
	var parent string = ctrl.GetString("parent")
	var name string = ctrl.GetString("name")
	var filetype string = ctrl.GetString("type")

	if parent == "" || name == "" || filetype == "" {
		ctrl.ThrowError(200, "Parent, Name and Type are required")
		return
	}

	parent = "./data/" + parent

	if filetype == "file" {
		f, err := os.Create(parent + name)
		if err != nil {
			ctrl.ThrowError(200, err.Error())
			return
		}
		defer f.Close()
	} else {
		if err := os.MkdirAll(parent+name, os.ModePerm); err != nil {
			ctrl.ThrowError(200, err.Error())
			return
		}
	}

	ctrl.Response("Created successfully")
}
