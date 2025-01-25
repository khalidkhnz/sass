package dto

type ApiResponse struct {
	Success bool
	Message string
}

type ApiResponseWithData struct {
	Success bool
	Message string
	Data    any
}

func ApiSuccessMsg(Message string) *ApiResponse {
	return &ApiResponse{
		Success: true,
		Message: Message,
	}
}

func ApiFailedMsg(Message string) *ApiResponse {
	return &ApiResponse{
		Success: false,
		Message: Message,
	}
}

type SuccessMsgWithDataProps struct {
	Message string
	Data    any
}

func ApiSuccessMsgWithData(Props SuccessMsgWithDataProps) *ApiResponseWithData {
	return &ApiResponseWithData{
		Success: true,
		Message: Props.Message,
		Data:    Props.Data,
	}
}