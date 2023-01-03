package resp

type Resp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ToStruct(data interface{}, err error) *Resp {
	if err == nil {
		return &Resp{
			Code: 0,
			Msg:  "success",
			Data: data,
		}
	}
	return &Resp{
		Code: -1,
		Msg:  err.Error(),
		Data: nil,
	}
}
