package models

type MathCompetitor struct {
	Uuid               string  `json:"uuid"`
	Cid                string  `json:"code"`
	Name               string  `json:"name"`
	Level              string  `json:"level"`
	Catergory          string  `json:"catergory"`
	Arena              string  `json:"arena"`
	School             string  `json:"school"`
	Province           string  `json:"province"`
	Region             string  `json:"region"`
	CalculationSection float32 `json:"cal_sect"`
	ProblemSection     float32 `json:"prob_sect"`
	AppliedSection     float32 `json:"applied_sect"`
	Score              float32 `json:"score"`
	ProvinceRank       int     `json:"prov_rank"`
	RegionalRank       int     `json:"reg_rank"`
}
