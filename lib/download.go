package lib

import (
	"errors"
	"io/ioutil"
	"os"
	"sync"

	"github.com/rs/zerolog/log"
)

// DownloadAll downloads all the HashiCorp tools
func DownloadAll(os, arch, bindir string) {
	downloadProducts(os, arch, bindir)
}

func downloadProducts(osName, archName, bindir string) {
	EnsureDirExists(bindir)

	dlP := GetProductList(osName, archName)

	tmpDir, err := ioutil.TempDir("", "pullhashi")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir) // clean up

	var wg sync.WaitGroup
	wg.Add(len(dlP))

	for _, p := range dlP {
		go func(p DownloadProduct) {
			defer wg.Done()
			if skipProduct(osName, archName, p) {
				return
			}
			downloadFiles(tmpDir, bindir, p)
		}(p)
	}

	wg.Wait()
}

func downloadFiles(tmpdir, bindir string, p DownloadProduct) {
	files := [...]string{p.File, p.SHASUM, p.SIG}
	// Download files
	for _, f := range files {
		Curl(tmpdir+"/"+f, p.URL+"/"+f)
	}

	// Verify GPG
	err := ShaIsSigned(tmpdir+"/"+p.SIG, tmpdir+"/"+p.SHASUM)
	if err != nil {
		panic(err)
	} else {
		log.Debug().
			Str("product", p.Name).
			Msg("shasums were signed by hashicorp")
	}

	// Verify SHA
	same, err := PassSumTest(tmpdir+"/"+p.SHASUM, tmpdir+"/"+p.File)
	if err != nil {
		panic(err)
	}
	if !same {
		panic(errors.New("file did not pass sha sum test"))
	} else {
		log.Debug().
			Str("product", p.Name).
			Msg("file sha matched shasums")
	}

	// Unzip
	unzipFiles, err := Unzip(tmpdir+"/"+p.File, bindir)
	if err != nil {
		panic(err)
	}
	for _, f := range unzipFiles {
		log.Info().
			Str("product", p.Name).
			Msg("created: " + f)
	}
}
