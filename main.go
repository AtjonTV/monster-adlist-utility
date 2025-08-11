package main

import (
	"flag"
	"os"

	"atjon.tv/monster/src/monster"
)

func main() {
	var sourceYaml string
	var outDir string
	var updateFor string
	var relinkBase bool

	flag.StringVar(&sourceYaml, "source", "sources.yaml", "Path to sources.yaml")
	flag.StringVar(&outDir, "out", "./", "Path to an output directory, where both monster.list and (optionally) monster.update will be written to.")
	flag.StringVar(&updateFor, "update", "monster_base.list", "Create an .update file for the given .list and the newly created .list")
	flag.BoolVar(&relinkBase, "relink", false, "Relink the monster_base.list to the newly created monster.list inside the output directory")
	flag.Parse()

	sources, err := monster.LoadSourcesFromFile(sourceYaml)
	if err != nil {
		panic(err)
	}

	err = monster.DownloadSources(&sources)
	if err != nil {
		panic(err)
	}

	err = monster.PrepareSources(&sources)
	if err != nil {
		panic(err)
	}

	newList, err := monster.BuildMonster(&sources, outDir)
	if err != nil {
		panic(err)
	}

	if updateFor != "" {
		_, err = os.Stat(outDir + string(os.PathSeparator) + updateFor)
		if os.IsNotExist(err) {
			_, err = os.Stat(updateFor)
			if os.IsNotExist(err) {
				panic("The updateFor file was not found: " + updateFor)
			}
		} else {
			updateFor = outDir + string(os.PathSeparator) + updateFor
		}

		err = monster.CreatePatch(&sources, updateFor, newList)
		if err != nil {
			panic(err)
		}
	}

	if relinkBase {
		err := os.Chdir(outDir)
		if err != nil {
			panic(err)
		}
		stat, err := os.Stat("monster_base.list")
		if stat != nil {
			err = os.Remove("monster_base.list")
			if err != nil {
				panic(err)
			}
		}

		var newRelative = newList[len(outDir)+1:]

		err = os.Symlink(newRelative, "monster_base.list")
		if err != nil {
			panic(err)
		}

		err = os.Chdir("..")
		if err != nil {
			panic(err)
		}
	}
}
