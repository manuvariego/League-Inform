package api

func NewWSConnection() *WSInfo {
	return &WSInfo{
		Seq: new(int64),
	}
}
