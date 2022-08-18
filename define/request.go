package define

type Request struct {
	Cpu    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
}

type Recommender struct {
	Cpu    []float64 `json:"cpu"`
	Memory []float64 `json:"memory"`
	CRPKI  []float64 `json:"crpki"`
	Misses []float64 `json:"misses"`
}
