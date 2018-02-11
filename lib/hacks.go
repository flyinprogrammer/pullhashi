package lib

import (
	"github.com/rs/zerolog/log"
)

func skipProduct(os, arch string, product DownloadProduct) bool {
	switch {
	case product.Name == "vault-ssh-helper" && product.Version == "0.1.1":
		log.Debug().
			Str("product", product.Name).
			Msg("skipping because: https://github.com/hashicorp/vault-ssh-helper/issues/33")
		return true
	default:
		return false
	}
}
