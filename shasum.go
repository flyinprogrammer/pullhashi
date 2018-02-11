package pullhashi

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// PassSumTest sees if a file has the correct sha.
func PassSumTest(shaFileName, fileToTest string) (bool, error) {
	sum, err := getSum(shaFileName, fileToTest)
	if err != nil {
		return false, err
	}
	realSum, err := getRealSum(fileToTest)
	if err != nil {
		return false, err
	}
	if sum == realSum {
		return true, nil
	}

	return false, nil
}

func getSum(shaFileName string, fileName string) (string, error) {
	shaFile, err := os.Open(shaFileName)
	if err != nil {
		return "", err
	}
	defer shaFile.Close()

	scanner := bufio.NewScanner(shaFile)
	justName := filepath.Base(fileName)
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), justName) {
			parts := strings.Split(scanner.Text(), "  ")
			return parts[0], nil
		}
	}
	return "", errors.New("could not find sha for file specified")
}

func getRealSum(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	bytes := sha256.Sum256(data)
	return hex.EncodeToString(bytes[:]), nil
}
