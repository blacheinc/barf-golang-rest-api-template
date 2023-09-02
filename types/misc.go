package types

type Home struct {
	Status      bool   `json:"status"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Website     string `json:"website"`
}

type Health struct {
	Name    string `json:"name"`
	Status  bool   `json:"status"`
	Version string `json:"version"`
}
