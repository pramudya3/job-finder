package response

type (
	ResSuccess struct {
		Data interface{} `json:"data"`
		Meta interface{} `json:"meta,omitempty"`
	}

	ResFailed struct {
		Message interface{} `json:"message"`
	}
)

func ResponseSuccess(data, meta interface{}) *ResSuccess {
	return &ResSuccess{
		Data: data,
		Meta: meta,
	}
}

func ResponseFailed(message interface{}) *ResFailed {
	return &ResFailed{
		Message: message,
	}
}
