package models

type AssetNode struct {
	Criticality Criticality `json:"criticality"`
}

type Criticality string

const (
	CriticalityA Criticality = "criticality_a"
	CriticalityB Criticality = "criticality_b"
	CriticalityC Criticality = "criticality_c"
)
