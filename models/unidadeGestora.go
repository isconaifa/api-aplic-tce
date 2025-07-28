package models

type UnidadeGestora struct {
	Codigo      int    `json:"codigo"`
	Nome        string `json:"nome"`
	Conselheiro string `json:"conselheiro"`
}
