package merkletree

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
)

// Node represents a node in the Merkle tree.
type Node struct {
	Hash  string `json:"hash"`
	Left  *Node  `json:"left,omitempty"`
	Right *Node  `json:"right,omitempty"`
}

// MerkleTree represents the entire Merkle tree.
type MerkleTree struct {
	Root *Node `json:"root"`
}

// hashData computes the SHA256 hash of the given data.
func hashData(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// NewNode creates a new tree node.
func NewNode(left, right *Node) *Node {
	node := &Node{Left: left, Right: right}
	// Concatenate left and right hashes and hash them to get the parent hash.
	combinedHash := left.Hash
	if right != nil { // Safely handle when there's no right child.
		combinedHash += right.Hash
	}
	node.Hash = hashData(combinedHash)
	return node
}

// BuildMerkleTree constructs a Merkle tree from the given hashes and returns it.
func BuildMerkleTree(hashes []string) *MerkleTree {
	if len(hashes) == 0 {
		return nil // Handle empty hash slice case.
	}

	var nodes []*Node
	for _, hash := range hashes {
		nodes = append(nodes, &Node{Hash: hash})
	}

	for len(nodes) > 1 {
		var level []*Node

		for i := 0; i < len(nodes); i += 2 {
			var right *Node
			if i+1 < len(nodes) {
				right = nodes[i+1]
			}
			level = append(level, NewNode(nodes[i], right))
		}

		nodes = level
	}

	return &MerkleTree{Root: nodes[0]}
}

// SaveTreeAsJSON saves the Merkle tree to a JSON file.
func SaveTreeAsJSON(tree *MerkleTree, filePath string) error {
	data, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}
