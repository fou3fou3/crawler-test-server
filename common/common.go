package common

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashSHA256(index int) string {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%d", index)))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func GeneratePagesHashes(nPages int) []string {
	var pagesHashes []string

	for nPage := 0; nPage < nPages; nPage++ {
		pageHash := HashSHA256(nPage)
		pagesHashes = append(pagesHashes, pageHash)
	}

	return pagesHashes
}
