package hyphae

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/bouncepaw/mycorrhiza/mimetype"
	"github.com/bouncepaw/mycorrhiza/util"
)

// Index finds all hypha files in the full `path` and sends them to the channel. Handling of duplicate entries and attachment and counting them is up to the caller.
func Index(path string, nestLevel uint, ch chan *Hypha) {
	nodes, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {
		// If this hypha looks like it can be a hypha path, go deeper. Do not touch the .git and static folders for they have an admnistrative importance!
		if node.IsDir() &&
			util.IsCanonicalName(node.Name()) &&
			node.Name() != ".git" &&
			!(nestLevel == 0 && node.Name() == "static") {
			Index(filepath.Join(path, node.Name()), nestLevel+1, ch)
			continue
		}

		var (
			hyphaPartPath           = filepath.Join(path, node.Name())
			hyphaName, isText, skip = mimetype.DataFromFilename(hyphaPartPath)
			hypha                   = &Hypha{Name: hyphaName}
		)
		if !skip {
			if isText {
				hypha.TextPath = hyphaPartPath
			} else {
				hypha.BinaryPath = hyphaPartPath
			}
			ch <- hypha
		}

	}
}
