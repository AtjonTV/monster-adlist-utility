/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"atjon.tv/monster/src/monster"
)

func main() {
	var sourceYaml string
	var outDir string
	var makeDiff bool
	var diffAgainst string
	var relinkBase bool
	var disableRewrite bool
	var doRewrite bool
	var disableCleanup bool
	var doCleanup bool
	var doVerboseLog bool

	flag.StringVar(&sourceYaml, "source", "sources.yaml", "Path to sources.yaml")
	flag.StringVar(&outDir, "out", "./", "Path to an output directory, where both monster.list and monster.update (diff) will be written to.")
	flag.BoolVar(&makeDiff, "diff", false, "Create an .update (diff) file")
	flag.StringVar(&diffAgainst, "diff-file", "monster_base.list", "Create an .update (diff) file for the given .list and the newly created .list")
	flag.BoolVar(&relinkBase, "relink", false, "Relink the monster_base.list to the newly created monster.list inside the output directory")
	flag.BoolVar(&disableRewrite, "no-rewrite", false, "Explicitly disable the rewrite feature, even when enabled in sources.yaml")
	flag.BoolVar(&doRewrite, "rewrite", false, "Explicitly enable the rewrite feature, even when disabled in sources.yaml; Forces --no-rewrite to be false")
	flag.BoolVar(&disableCleanup, "no-cleanup", false, "Explicitly disable the cleanup feature, even when enabled in sources.yaml")
	flag.BoolVar(&doCleanup, "cleanup", false, "Explicitly enable the cleanup feature, even when disabled in sources.yaml; Forces --no-cleanup to be false")
	flag.BoolVar(&doVerboseLog, "verbose", false, "Enable verbose (debug) logging")
	flag.Parse()

	sources, err := monster.LoadSourcesFromFile(sourceYaml)
	if err != nil {
		panic(err)
	}

	monster := monster.New(sources, doVerboseLog)

	err = monster.DownloadSources()
	if err != nil {
		panic(err)
	}

	err = monster.PrepareSources()
	if err != nil {
		panic(err)
	}

	stat, err := os.Stat(outDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(outDir, 0755)
		if err != nil {
			panic(err)
		}
	} else if !stat.IsDir() {
		panic(fmt.Sprintf("The output directory '%s' is not a directory.\n", outDir))
	}

	sources.Rewrite.Enable = doRewrite || (sources.Rewrite.Enable && !disableRewrite)

	newList, err := monster.BuildMonster(outDir)
	if err != nil {
		panic(err)
	}

	if makeDiff {
		var doCreateDiff = true
		_, err = os.Stat(outDir + string(os.PathSeparator) + diffAgainst)
		if os.IsNotExist(err) {
			_, err = os.Stat(diffAgainst)
			if os.IsNotExist(err) {
				fmt.Printf("WARN: The .list file ('%s') to diff against was not found, skipping\n", diffAgainst)
				doCreateDiff = false
			}
		} else {
			diffAgainst = outDir + string(os.PathSeparator) + diffAgainst
		}

		if doCreateDiff {
			err = monster.CreateDiffFile(diffAgainst, newList)
			if err != nil {
				fmt.Printf("WARN: Failed to create diff due to an error: %s\n", err)
			}
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

	sources.CleanRule.Enable = doCleanup || (sources.CleanRule.Enable && !disableCleanup)

	if sources.CleanRule.Enable {
		err = monster.CleanUp(outDir)
		if err != nil {
			panic(err)
		}
	}
}
