package helpers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/sergerad/incremental-merkle-tree/imt"
	log "github.com/sirupsen/logrus"
)

func UpdateMerkleTree(sortedData []string, tree *imt.IncrementalMerkleTree) (*imt.IncrementalMerkleTree, error) {
	log.Debugln("current hash: ", common.Bytes2Hex(tree.RootDigest()))
	for _, value := range sortedData {
		err := tree.AddLeaf([]byte(value))
		if err != nil {
			log.Errorf("Error adding merkle tree leaf: %s\n", err.Error())
			return nil, err
		}
	}

	return tree, nil
}
