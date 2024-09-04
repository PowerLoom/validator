package helpers

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

func HandleAttestationSubmissionError(err error, multiplier int, id string) int {
	log.Debugf("Found error: %s proceeding with adjustments for attestation submission\n", err.Error())
	if strings.Contains(err.Error(), "transaction underpriced") {
		log.Errorf("Could not submit batch: %s error: %s\n", id, err.Error())
		multiplier++
		UpdateGasPrice(multiplier)
		log.Debugln("Retrying with gas price: ", Auth.GasPrice.String())
	} else if strings.Contains(err.Error(), "nonce too low") {
		log.Errorf("Nonce too low for batch: %s error: %s\n", id, err.Error())
		UpdateAuth(1)
		log.Debugln("Retrying with nonce: ", Auth.Nonce.String())
	} else if strings.Contains(err.Error(), "nonce too high") {
		log.Errorf("Nonce too low for batch: %s error: %s\n", id, err.Error())
		UpdateAuth(-1)
		log.Debugln("Retrying with nonce: ", Auth.Nonce.String())
	} else {
		// Handle other errors
		log.Errorf("Unexpected error: %v", err)
	}
	return multiplier
}
