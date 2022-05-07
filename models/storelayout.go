package models

type StorefrontLayout struct {
	Data []elem `json:"data"`
}

type elem struct {
	Section     string       `json:"section"`
	Storefronts []storefront `json:"storefronts"`
}

type storefront struct {
	Name string `json:"name"`
	Item string `json:"item"`
}
