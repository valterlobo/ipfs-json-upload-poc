package nft

type MetadataNFT struct {
	Title      string `json:"title"`
	Type       string `json:"type"`
	Properties struct {
		Name struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"name"`
		Description struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"description"`
		Image struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"image"`
	} `json:"properties"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	TraitType   string `json:"trait_type"`
	Value       string `json:"value"`
	DisplayType string `json:"display_type,omitempty"`
}
