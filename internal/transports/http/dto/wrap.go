package dto

type Wrap struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Err    string      `json:"err"`
}

var succeed = Wrap{Status: 0}

func CreateSucceedWithData(i interface{}) *Wrap {
	r := Wrap{Status: 0, Data: i}
	return &r
}

func CreateSucceed() *Wrap {
	return &succeed
}

func CreateFailWithError(s int, err string) *Wrap {
	r := Wrap{Status: s, Err: err}
	return &r
}
