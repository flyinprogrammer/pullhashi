package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/blang/semver"
)

// GetProductList get a list of all HashiCorp products that we can download for your system.
func GetProductList(os string, arch string) (downloadProducts []DownloadProduct) {
	resp, _ := http.Get(ReleaseURL + "/index.json")
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	var products Products
	err = json.Unmarshal(b, &products)
	if err != nil {
		panic(err)
	}

	// TODO: Does not support Vagrant because the release packaging it completely different.
	for _, product := range products {
		if strings.HasPrefix(product.Name, "terraform-provider") {
			// TODO: make this optional
			continue
		}
		var versionsToSort semver.Versions
		buildsWeWant := make(map[string]Build)
		for _, version := range product.Versions {
			for _, build := range version.Builds {
				if build.OS == os {
					if build.Arch == arch {
						semV, err := semver.Make(build.Version)
						if err != nil {
							panic(err)
						} else {
							versionsToSort = append(versionsToSort, semV)
							buildsWeWant[build.Version] = build
						}
					}
				}
			}
		}
		sort.Sort(sort.Reverse(versionsToSort))
		if len(versionsToSort) > 0 {
			version := versionsToSort[0].String()
			fileName := buildsWeWant[version].Filename
			sha := product.Versions[version].Shasums
			sig := product.Versions[version].ShasumsSignature
			url := fmt.Sprintf("%s/%s/%s", ReleaseURL, product.Name, version)

			downloadProducts = append(downloadProducts, DownloadProduct{
				Name:    product.Name,
				Version: version,
				URL:     url,
				File:    fileName,
				SHASUM:  sha,
				SIG:     sig,
			})
		}
	}
	return
}
