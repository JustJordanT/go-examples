package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	file := "api.json"
	sha := "c25cf40c2ebedbea04b317f55a9cee9ed35074a4886b60376b1421832c5e3389"
	fmt.Println(CompareSHA256(file, sha))
}

func CompareSHA256(filename string, expectedSHA string) (bool, error) {
	// Read the contents of the file
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return false, err
	}

	// Calculate the SHA256 checksum of the file
	sha256Bytes := sha256.Sum256(fileBytes)
	actualSHA := hex.EncodeToString(sha256Bytes[:])

	// Compare the calculated SHA256 with the expected value
	if actualSHA != expectedSHA {
		return false, nil
	}

	return true, nil
}
