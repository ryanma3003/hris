package models

type Casbin_rule struct {
	Id     int64  `json:"id"`
	P_type string `json:"p_type"`
	V0     string `json:"v0"`
	V1     string `json:"v1"`
	V2     string `json:"v2"`
	V3     string `json:"v3"`
	V4     string `json:"v4"`
	V5     string `json:"v5"`
}
