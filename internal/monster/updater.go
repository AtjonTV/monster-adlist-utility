/*
 * Copyright (c) 2025 Thomas Obernosterer. Licensed under the EUPL-1.2 or later.
 */

package monster

import (
	"os"
	"strings"

	"atjon.tv/monster/internal/utils"
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
		err := utils.WriteDataToFile(list.TempFile, []byte(strings.Join(list.Data, "\n")))
		if err != nil {
			return err
		}
		return nil
	}
	err := utils.DownloadFileToPath(list.Url, list.TempFile)
	if err != nil {
		// TODO: We ignore download errors, maybe we should retry and write a log
		//return err
	}
	return nil
}
