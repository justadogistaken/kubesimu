package define

type Node struct {
	Cpu         float64     `json:"cpu"`
	Memory      float64     `json:"memory"`
	Oversell    float64     `json:"oversell"`
	Pods        []Pod       `json:"pods"`
	Recommender Recommender `json:"recommender"`
}

func (n *Node) ActualCPU() {
	return n.Cpu
}

func (n *Node) ActualCPU() {

}

func (n *Node) ActualCPU() {

}

func (n *Node) ActualCPU() {

}
