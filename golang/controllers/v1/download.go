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
