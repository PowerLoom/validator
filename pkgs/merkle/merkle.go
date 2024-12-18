package merkle

import (
	"github.com/sergerad/incremental-merkle-tree/imt"
	log "github.com/sirupsen/logrus"
)

// UpdateMerkleTree adds all provided IDs to the Merkle tree as leaves and returns the updated tree
func UpdateMerkleTree(ids []string, tree *imt.IncrementalMerkleTree) (*imt.IncrementalMerkleTree, error) {
	for _, id := range ids {
		err := tree.AddLeaf([]byte(id))
		if err != nil {
			log.Errorf("Error adding leaf to Merkle tree: %s\n", err.Error())
			return nil, err
		}
	}
	return tree, nil
}
