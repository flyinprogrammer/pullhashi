package pullhashi

// Products a map of products
type Products map[string]Product

// Product is a single tool like consul
type Product struct {
	Name     string             `json:"name"`
	Versions map[string]Version `json:"versions"`
}

// Version contains the version metadata
type Version struct {
	Name             string `json:"name"`
	Version          string `json:"version"`
	Shasums          string `json:"shasums"`
	ShasumsSignature string `json:"shasums_signature"`
	Builds           Builds `json:"builds"`
}

// Build Contains the multi-os & multi-arch metadata for a version.
type Build struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

// DownloadProduct All the metadata needed to download and verify a file.
type DownloadProduct struct {
	Name    string
	URL     string
	File    string
	SHASUM  string
	SIG     string
	Version string
}

// Builds a collection of Build objects.
type Builds []Build
