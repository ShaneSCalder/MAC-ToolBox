package main

import (
	"MERKLETREE/pkg/merkletree" // Adjust the import path to match your project structure
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	chunkDir := "data/output"
	chunkFiles, err := ioutil.ReadDir(chunkDir)
	if err != nil {
		log.Fatalf("Failed to read chunk directory: %v", err)
	}

	var hashes []string
	for _, f := range chunkFiles {
		data, err := ioutil.ReadFile(filepath.Join(chunkDir, f.Name()))
		if err != nil {
			log.Fatalf("Failed to read chunk file: %v", err)
		}

		hash := sha256.Sum256(data)
		hashes = append(hashes, hex.EncodeToString(hash[:]))
	}

	// Assuming BuildMerkleTree is a function of the merkletree package
	tree := merkletree.BuildMerkleTree(hashes)

	// Assuming SaveTreeAsJSON is a method that takes a MerkleTree pointer
	err = merkletree.SaveTreeAsJSON(tree, "data/merkle_tree.json")
	if err != nil {
		log.Fatalf("Failed to save Merkle tree as JSON: %v", err)
	}

	fmt.Println("Merkle tree created and saved as JSON.")
}
