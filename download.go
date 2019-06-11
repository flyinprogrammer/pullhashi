package pullhashi

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
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
	err = removeEmptyBinBin(bindir + "/bin")
}

func downloadFiles(tmpdir, bindir string, p DownloadProduct) {
	files := [...]string{p.File, p.SHASUM, p.SIG}
	// Download files
	for _, f := range files {
		Curl(tmpdir+"/"+f, p.URL+"/"+f)
	}

	// Verify GPG
	log.Debug().Str("product", p.Name).Msg("about to validate signing")
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
	err = cleanUpNestedBin(unzipFiles)
	if err != nil {
		panic(err)
	}
}

func cleanUpNestedBin(files []string) error {
	for _, f := range files {
		if strings.Contains(f, "/bin/bin/") {
			// we have a nested bin
			newPath := strings.Replace(f, "bin/", "", 1)
			log.Info().Str("og", f).Str("new", f).Msg("moving file")
			err := os.Rename(f, newPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func removeEmptyBinBin(dir string) error {
	empty, err := isEmptyDir(dir)

	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil
		}
		log.Debug().Bool("empty", empty).Err(err).Msg("error testing if empty dir")
		return err
	}
	if empty {
		log.Debug().Str("dir", dir).Msg("removing empty dir")
		os.RemoveAll(dir)
	}
	return nil
}

func isEmptyDir(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
