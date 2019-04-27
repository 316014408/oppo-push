package oppopush

// 通知栏消息
type NotificationMessage struct {
	AppMessageID        string `json:"app_message_id,omitempty"`        //App 开发者自定义消息 Id，OPPO 推送平台根据此 ID 做去重处理，对于广播推送相同app_message_id只会保存一次，对于单推相同 app_message_id 只会推送一次。
	Title               string `json:"title"`                           //设置在通知栏展示的通知栏标题, 【字数限制 1~32，中英文均以一个计算】
	SubTitle            string `json:"sub_title,omitempty"`             //子标题 设置在通知栏展示的通知栏标题, 【字数限制 1~10，中英文均以一个计算】
	Content             string `json:"content"`                         //设置在通知栏展示的通知的内容,【必填，字数限制 200 以内，中英文均以一个计算】
	ClickActionType     int    `json:"click_action_type,omitempty"`     //点击动作类型 0，启动应用；1，打开应用内页（activity 的 intent action）；2，打开网页；4，打开应用内页（activity）；【非必填，默认值为 0】;5,Intentscheme URL
	ClickActionActivity string `json:"click_action_activity,omitempty"` //应用内页地址【click_action_type 为1/4/时必填，长度 500】
	ClickActionURL      string `json:"click_action_url,omitempty"`      //网页地址或【click_action_type 为 2与 5 时必填，长度 500】
	ActionParameters    string `json:"action_parameters,omitempty"`     //动作参数，打开应用内页或网页时传递给应用或网页【JSON 格式，非必填】，字符数不能超过 4K
	ShowTimeType        int    `json:"show_time_type,omitempty"`        //展示类型 (0, “即时”),(1, “定时”)
	ShowStartTime       int64  `json:"show_start_time,omitempty"`       //定时展示开始时间（根据 time_zone 转换成当地时间），时间的毫秒数
	ShowEndTime         int64  `json:"show_end_time,omitempty"`         //定时展示结束时间（根据 time_zone 转换成当地时间），时间的毫秒数
	OffLine             bool   `json:"off_line,omitempty"`              //是否进离线消息,【非必填，默认为True】
	OffLineTTL          int    `json:"off_line_ttl,omitempty"`          //离线消息的存活时间(time_to_live)(单位：秒), 【off_line 值为 true 时，必填，最长 3 天】
	PushTimeType        int    `json:"push_time_type,omitempty"`        // 否 0 定时推送 (0, “即时”),(1, “定时”), 【只对全部用户推送生效】
	PushStartTime       int64  `json:"push_start_time,omitempty"`       //定时推送开始时间（根据 time_zone 转换成当地时间）,【push_time_type 为1 必填】，时间的毫秒数
	TimeZone            string `json:"time_zone,omitempty"`             //时区，默认值：（GMT+08:00）北京，香港，新加坡
	FixSpeed            bool   `json:"fix_speed,omitempty"`             //是否定速推送,【非必填，默认值为false】
	FixSpeedRate        int64  `json:"fix_speed_rate,omitempty"`        //定速速率【fixSpeed 为 true 时，必填】
	NetworkType         int    `json:"network_type,omitempty"`          //0：不限联网方式, 1：仅 wifi 推送；
	CallBackURL         string `json:"call_back_url,omitempty"`         //*仅支持 registrationId 或 aliasName两种推送方式*应用接收消息到达回执的回调 URL，字数限制200以内，中英文均以一个计算。OPPO Push 服务器 POST 一个 JSON 数据到 call_back_url；
	CallBackParameter   string `json:"call_back_parameter,omitempty"`   //App 开发者自定义回执参数，字数限制50 以内，中英文均以一个计算。
	ChannelID           string `json:"channel_id,omitempty"`            //通知栏通（NotificationChannel），从Android9 开始发送通知消息必须要指定通道 Id
}

// 广播推送
type Broadcast struct {
	MessageID   string `json:"message_id"`             //消息Id
	TargetType  int    `json:"target_type"`            //目标类型，1:ALL;2:registration_id
	TargetValue string `json:"target_value,omitempty"` //推送目标用户【多个以英文分号(;)分隔，最大1000个】，可以替代registration_id
}

// 单推
type Unicast struct {
	Messages Message `json:"messages"` //通知栏消息 JSON String
}

// 单推消息
type Message struct {
	TargetType     int                 `json:"target_type"`               //目标类型 2:registration_id ;3: 别名推送alias_name;
	RegistrationID string              `json:"registration_id,omitempty"` //应用级设备注册唯一标识符【target_type 为 2 必填】
	TargetValue    string              `json:"target_value"`              //推送目标用户
	Notification   NotificationMessage `json:"notification"`              //请参见通知栏消息
}

// 通知栏消息内容体
func NewMessage(notificationTitle, notificationContent string) *Message {
	return &Message{
		Notification: NotificationMessage{
			Title:   notificationTitle,
			Content: notificationContent,
		},
	}
}

// 保存通知栏消息内容体
func NewSaveMessageContent(notificationTitle, notificationContent string) *NotificationMessage {
	return &NotificationMessage{
		Title:   notificationTitle,
		Content: notificationContent,
	}
}

// 广播推送
func NewBroadcast(messageID string) *Broadcast {
	return &Broadcast{
		MessageID: messageID,
	}
}

//--------------------------------------------------------------------------------

func (n *NotificationMessage) SetID(appMessageID string) *NotificationMessage {
	n.AppMessageID = appMessageID
	return n
}

func (n *NotificationMessage) SetSubTitle(subTitle string) *NotificationMessage {
	n.SubTitle = subTitle
	return n
}

func (m *Message) SetSubTitle(subTitle string) *Message {
	m.Notification.SubTitle = subTitle
	return m
}

func (n *NotificationMessage) SetClickActionType(clickActionType int) *NotificationMessage {
	n.ClickActionType = clickActionType
	return n
}

func (m *Message) SetClickActionType(clickActionType int) *Message {
	m.Notification.ClickActionType = clickActionType
	return m
}

func (n *NotificationMessage) SetClickActionActivity(clickActionActivity string) *NotificationMessage {
	n.ClickActionActivity = clickActionActivity
	return n
}

func (m *Message) SetClickActionActivity(clickActionActivity string) *Message {
	m.Notification.ClickActionActivity = clickActionActivity
	return m
}

func (n *NotificationMessage) SetClickActionUrl(clickActionUrl string) *NotificationMessage {
	n.ClickActionURL = clickActionUrl
	return n
}

func (m *Message) SetClickActionUrl(clickActionUrl string) *Message {
	m.Notification.ClickActionURL = clickActionUrl
	return m
}

func (n *NotificationMessage) SetActionParameters(actionParameters string) *NotificationMessage {
	n.ActionParameters = actionParameters
	return n
}

func (m *Message) SetActionParameters(actionParameters string) *Message {
	m.Notification.ActionParameters = actionParameters
	return m
}

func (n *NotificationMessage) SetShowTimeType(showTimeType int) *NotificationMessage {
	n.ShowTimeType = showTimeType
	return n
}

func (n *NotificationMessage) SetShowStartTime(showStartTime int64) *NotificationMessage {
	n.ShowStartTime = showStartTime
	return n
}

func (n *NotificationMessage) SetShowEndTime(showEndTime int64) *NotificationMessage {
	n.ShowEndTime = showEndTime
	return n
}

func (n *NotificationMessage) SetOffLine(offLine bool) *NotificationMessage {
	n.OffLine = offLine
	return n
}

func (m *Message) SetOffLine(offLine bool) *Message {
	m.Notification.OffLine = offLine
	return m
}

func (n *NotificationMessage) SetOffLineTtl(offLineTtl int) *NotificationMessage {
	n.OffLineTTL = offLineTtl
	return n
}

func (m *Message) SetOffLineTtl(offLineTtl int) *Message {
	m.Notification.OffLineTTL = offLineTtl
	return m
}

func (n *NotificationMessage) SetPushTimeType(pushTimeType int) *NotificationMessage {
	n.PushTimeType = pushTimeType
	return n
}

func (n *NotificationMessage) SetPushStartTime(pushStartTime int64) *NotificationMessage {
	n.PushStartTime = pushStartTime
	return n
}

func (n *NotificationMessage) SetTimeZone(timeZone string) *NotificationMessage {
	n.TimeZone = timeZone
	return n
}

func (m *Message) SetTimeZone(timeZone string) *Message {
	m.Notification.TimeZone = timeZone
	return m
}

func (n *NotificationMessage) SetFixSpeed(fixSpeed bool) *NotificationMessage {
	n.FixSpeed = fixSpeed
	return n
}

func (n *NotificationMessage) SetFixSpeedRate(fixSpeedRate int64) *NotificationMessage {
	n.FixSpeedRate = fixSpeedRate
	return n
}

func (n *NotificationMessage) SetNetworkType(networkType int) *NotificationMessage {
	n.NetworkType = networkType
	return n
}

func (n *NotificationMessage) SetCallBackUrl(callBackUrl string) *NotificationMessage {
	n.CallBackURL = callBackUrl
	return n
}

func (m *Message) SetCallBackUrl(callBackUrl string) *Message {
	m.Notification.CallBackURL = callBackUrl
	return m
}

func (n *NotificationMessage) SetCallBackParameter(callBackParameter string) *NotificationMessage {
	n.CallBackParameter = callBackParameter
	return n
}

func (m *Message) SetCallBackParameter(callBackParameter string) *Message {
	m.Notification.CallBackParameter = callBackParameter
	return m
}

func (n *NotificationMessage) SetChannelId(channelId string) *NotificationMessage {
	n.ChannelID = channelId
	return n
}

func (m *Message) SetChannelId(channelId string) *Message {
	m.Notification.ChannelID = channelId
	return m
}

//--------------------------------------------------------------------------------

func (b *Broadcast) SetTargetType(targetType int) *Broadcast {
	b.TargetType = targetType
	return b
}

func (m *Message) SetTargetType(targetType int) *Message {
	m.TargetType = targetType
	return m
}

func (b *Broadcast) SetTargetValue(targetValue string) *Broadcast {
	b.TargetValue = targetValue
	return b
}

func (m *Message) SetTargetValue(targetValue string) *Message {
	m.TargetValue = targetValue
	return m
}
