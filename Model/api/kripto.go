package api

type Kripto struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Harga   float32 `json:"harga"`
	Uang    string  `json:"uang"`
}
