package pullhashi

import (
	"os"
	"strings"

	"golang.org/x/crypto/openpgp"
)

// ShaIsSigned validates the our SHASUMS file is properly signed.
func ShaIsSigned(sig, sha string) error {

	keyRingReader := strings.NewReader(HashicorpKey)

	signature, err := os.Open(sig)
	if err != nil {
		return err
	}
	defer signature.Close()

	verificationTarget, err := os.Open(sha)
	if err != nil {
		return err
	}
	defer verificationTarget.Close()

	keyring, err := openpgp.ReadArmoredKeyRing(keyRingReader)
	if err != nil {
		return err
	}
	_, err = openpgp.CheckDetachedSignature(keyring, verificationTarget, signature)
	if err != nil {
		return err
	}
	return nil
}
