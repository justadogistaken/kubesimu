package define

type Nodes struct {
	Number int     `json:"number"`
	Cpu    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
}

type Cluster struct {
	Nodes    []Nodes `json:"nodes"`
	Oversell float64 `json:"oversell"`
}
