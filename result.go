package oppopush

type AuthSendResult struct {
	Code    int    `json:"code"`              // 必选,返回码
	Message string `json:"message,omitempty"` // 可选，返回消息
	Data    struct {
		AuthToken  string `json:"auth_token"`  //权限令牌，推送消息时，需要提供 auth_token，有效期默认为 24 小时，过期后无法使用
		CreateTime int64  `json:"create_time"` //"时间毫秒数
	} `json:"data,omitempty"` // 可选，返回结果
}

type SaveSendResult struct {
	Code    int    `json:"code"`              // 必选,返回码
	Message string `json:"message,omitempty"` // 可选，返回消息
	Data    struct {
		MessageID string `json:"message_id"` //消息 ID
	} `json:"data,omitempty"` // 可选，返回结果
}

type BroadcastSendResult struct {
	Code    int    `json:"code"`              // 必选,返回码
	Message string `json:"message,omitempty"` // 可选，返回消息
	Data    struct {
		MessageID string `json:"message_id"` //消息 ID
		TaskId    string `json:"task_id"`    //推送任务 ID
	} `json:"data,omitempty"` // 可选，返回结果
}

type UnicastSendResult struct {
	Code    int    `json:"code"`              // 必选,返回码
	Message string `json:"message,omitempty"` // 可选，返回消息
	Data    struct {
		MessageID string `json:"messageId"` //消息 ID
	} `json:"data,omitempty"` // 可选，返回结果
}

type UnicastBatchSendResult struct {
	Code    int    `json:"code"`              // 必选,返回码
	Message string `json:"message,omitempty"` // 可选，返回消息
	Data    []struct {
		MessageID      string `json:"messageId"` //消息 ID
		RegistrationID string `json:"registrationId"`
		ErrorCode      int    `json:"errorCode,omitempty"`    // 失败码
		ErrorMessage   string `json:"errorMessage,omitempty"` // 失败说明
	} `json:"data,omitempty"` // 可选，返回结果
}

type FetchInvalidRegidListSendResult struct {
	Code    int    `json:"code"`              // 必选,返回码
	Message string `json:"message,omitempty"` // 可选，返回消息
	Data    struct {
		RegistrationIds []string `json:"registration_ids"`
		TotalCount      int      `json:"totalCount"`
	} `json:"data,omitempty"` // 可选，返回结果
}
