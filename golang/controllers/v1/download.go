package v1

import (
	"beego/controllers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

type Files struct {
	Name  string `json:"name"`
	IsDir bool   `json:"is_dir"`
	Type  string `json:"type"`
	Size  int64  `json:"size"`
}

type DownloadController struct {
	controllers.AppController
}

type downloadFileRequest struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

func (ctrl *DownloadController) AllFiles() {
	var parent string = ctrl.GetString("p")

	var dir string = "./data/" + parent + "/"
	entries, err := os.ReadDir(dir)
	if err != nil {
		ctrl.ThrowError(200, err.Error())
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
			}
			defer openfile.Close()
			contentType, err := ctrl.GetFileContentType(openfile)
			file.Type = contentType

			fi, err := openfile.Stat()
			if err != nil {
				ctrl.ThrowError(200, err.Error())
			}
			file.Size = fi.Size()

		} else {
			size, err := ctrl.DirSize(dir)
			if err != nil {
				ctrl.ThrowError(200, err.Error())
			}
			file.Size = size
		}
		files = append(files, file)
	}
	ctrl.Response(files)

}
func (ctrl *DownloadController) DirSize(path string) (int64, error) {
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
func (ctrl *DownloadController) GetFileContentType(ouput *os.File) (string, error) {

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

func (ctrl *DownloadController) DownloadFiles() {

	var fileReq downloadFileRequest
	json.Unmarshal(ctrl.Ctx.Input.RequestBody, &fileReq)

	getUrl := fileReq.Url
	if getUrl == "" {
		ctrl.ThrowError(1003, "Url is required")
	}
	u, err := url.ParseRequestURI(getUrl)
	if err != nil {
		ctrl.ThrowError(200, err.Error())
	}

	fileName := fileReq.Name
	if fileName == "" {
		r, _ := http.NewRequest("GET", getUrl, nil)
		fileName = path.Base(r.URL.Path)
	}
	if fileName == "" || fileName == "." {
		ctrl.ThrowError(200, "Invalid url")
	}

	err = ctrl.downloadFile(fileName, getUrl)
	if err != nil {
		fmt.Println(err)
		ctrl.ThrowError(200, err.Error())
	}

	ctrl.Response(u)
}

func (ctrl *DownloadController) downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create("./data/" + filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
