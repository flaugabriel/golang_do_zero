package modelos

type Senha struct {
	Atual string `json:"atual,omitempty"`
	Nova  string `json:"nova,omitempty"`
}
