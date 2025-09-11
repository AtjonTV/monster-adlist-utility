/*
 * Copyright (c) 2025 Thomas Obernosterer. Licensed under the EUPL-1.2 or later.
 */

package utils

import (
	"io"
	"net/http"
	"os"
)

func DownloadFileToPath(url string, path string) (err error) {
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
