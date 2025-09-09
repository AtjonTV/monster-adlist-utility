/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import (
	"io"
	"net/http"
	"os"
	"strings"
)

var tmpDir = os.TempDir()
var pathSeparator = string(os.PathSeparator)

func (m *Monster) DownloadSourceLists() error {
	for i := range m.Sources.Allow {
		err := processList(&m.Sources.Allow[i], "allow")
		if err != nil {
			// TODO: We ignore download errors, maybe we should retry and write a log
			//return err
		}
	}

	for i := range m.Sources.Block {
		err := processList(&m.Sources.Block[i], "block")
		if err != nil {
			// TODO: We ignore download errors, maybe we should retry and write a log
			//return err
		}
	}

	return nil
}

func processList(list *SourceList, suffix string) error {
	list.TempFile = tmpDir + pathSeparator + list.Name + "." + suffix
	if list.Url == "" && list.Data != nil {
		err := writeFile(list.TempFile, []byte(strings.Join(list.Data, "\n")))
		if err != nil {
			return err
		}
		return nil
	}
	err := downloadFile(list.Url, list.TempFile)
	if err != nil {
		// TODO: We ignore download errors, maybe we should retry and write a log
		//return err
	}
	return nil
}

func writeFile(path string, data []byte) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	_, err = out.Write(data)
	return err
}

func downloadFile(url string, path string) (err error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	_, err = io.Copy(out, resp.Body)

	return err
}
