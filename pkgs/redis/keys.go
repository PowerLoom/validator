package redis

import (
	"fmt"
	"strings"
	"validator/pkgs"
)

func ValidatorSet(dataMarketAddress, epochID string) string {
	return fmt.Sprintf("%s.%s.%s", pkgs.ValidatorSetKey, strings.ToLower(dataMarketAddress), epochID)
}

func SnapshotSubmissionValidatorKey(dataMarketAddress, epochID, rootHash string) string {
	return fmt.Sprintf("%s.%s.%s.%s", pkgs.ValidatorKey, strings.ToLower(dataMarketAddress), epochID, rootHash)
}
