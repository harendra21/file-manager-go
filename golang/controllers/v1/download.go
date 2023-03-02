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
)

type DownloadController struct {
	controllers.AppController
}

type downloadFileRequest struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

func (ctrl *DownloadController) DownloadFiles() {
	var parent string = ctrl.GetString("p")
	if parent == "" {
		ctrl.ThrowError(200, "Parent is required")
		return
	}

	var fileReq downloadFileRequest
	json.Unmarshal(ctrl.Ctx.Input.RequestBody, &fileReq)

	getUrl := fileReq.Url
	if getUrl == "" {
		ctrl.ThrowError(200, "Url is required")
		return
	}
	u, err := url.ParseRequestURI(getUrl)
	if err != nil {
		ctrl.ThrowError(200, err.Error())
		return
	}

	fileName := fileReq.Name
	if fileName == "" {
		r, _ := http.NewRequest("GET", getUrl, nil)
		fileName = path.Base(r.URL.Path)
	}
	if fileName == "" || fileName == "." {
		ctrl.ThrowError(200, "Invalid url")
		return
	}

	err = ctrl.downloadFile(parent, fileName, getUrl)
	if err != nil {
		ctrl.ThrowError(200, err.Error())
		return
	}

	ctrl.Response(u)
}

func (ctrl *DownloadController) downloadFile(parent, filepath, url string) (err error) {

	// Create the file
	out, err := os.Create("./data/" + parent + filepath)
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
