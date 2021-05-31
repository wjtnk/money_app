/*
apiレスポンスメッセージ
*/
package message

type ResponseMessage struct {
	Message string `json:"message"`
}

func (resultMsg ResponseMessage) GetSuccessMessage() ResponseMessage {
	message := ResponseMessage{
		Message: "success",
	}
	return message
}

func (resultMsg ResponseMessage) GetFailedMessage() ResponseMessage {
	message := ResponseMessage{
		Message: "failed",
	}
	return message
}

func (resultMsg ResponseMessage) GetErrorMessage(error string) ResponseMessage {
	message := ResponseMessage{
		Message: error,
	}
	return message
}
