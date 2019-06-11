package pullhashi

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
	case product.Name == "docker-basetool":
		log.Debug().
			Str("product", product.Name).
			Msg("skipping because this tool is fairly useless outside of Docker")
		return true
	case product.Name == "otto":
		log.Debug().
			Str("product", product.Name).
			Msg("skipping because: https://www.hashicorp.com/blog/decommissioning-otto.html")
		return true
	case product.Name == "atlas-upload-cli":
		log.Debug().
			Str("product", product.Name).
			Msg("skipping because: https://www.hashicorp.com/blog/vagrant-cloud-migration-announcement")
		return true
	case product.Name == "envconsul" && product.Version == "0.8.0":
		log.Debug().
			Str("product", product.Name).
			Msg("skipping because: https://github.com/hashicorp/envconsul/issues/214")
		return true
	default:
		return false
	}
}
