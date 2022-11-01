package score_api

type ScoreRes struct {
	Status  int     `json:"status"`
	Msg     string  `json:"msg"`
	Data    int     `json:"data,omitempty"`
	Metrics metrics `json:"metrics,omitempty"`
}
type metrics struct {
	Ap, Ar, Af1, Af05, Af2, Acc, Err interface{}
}
