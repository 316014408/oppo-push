package oppopush

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

type OppoPush struct {
	appKey       string
	masterSecret string
}

func NewClient(appKey, masterSecret string) *OppoPush {
	return &OppoPush{
		appKey:       appKey,
		masterSecret: masterSecret,
	}
}

// 保存通知栏消息内容体
func (c *OppoPush) SaveMessageContent(msg *NotificationMessage) (*SaveSendResult, error) {
	tokenInstance, err := GetToken(c.appKey, c.masterSecret)
	if err != nil {
		return nil, err
	}
	params := defaultForm(msg)
	params.Add("auth_token", tokenInstance.AccessToken)
	bytes, err := doPost(PushHost+SaveMessageContentURL, params)
	if err != nil {
		return nil, err
	}
	var result SaveSendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 广播推送-通知栏消息
func (c *OppoPush) Broadcast(broadcast *Broadcast) (*BroadcastSendResult, error) {
	tokenInstance, err := GetToken(c.appKey, c.masterSecret)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("message_id", broadcast.MessageID)
	params.Add("target_type", strconv.Itoa(broadcast.TargetType))
	params.Add("target_value", broadcast.TargetValue)
	params.Add("auth_token", tokenInstance.AccessToken)
	bytes, err := doPost(PushHost+MessageBroadcastURL, params)
	if err != nil {
		return nil, err
	}
	var result BroadcastSendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, errors.New(result.Message)
	}
	return &result, nil
}

// 单推-通知栏消息推送
func (c *OppoPush) Unicast(message *Message) (*UnicastSendResult, error) {
	tokenInstance, err := GetToken(c.appKey, c.masterSecret)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("message", message.String())
	params.Add("auth_token", tokenInstance.AccessToken)
	bytes, err := doPost(PushHost+MessageUnicastURL, params)
	if err != nil {
		return nil, err
	}
	var result UnicastSendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, errors.New(result.Message)
	}
	return &result, nil
}

// 批量单推-通知栏消息推送
func (c *OppoPush) UnicastBatch(messages []Message) (*UnicastBatchSendResult, error) {
	jsons, err := json.Marshal(messages)
	if err != nil {
		return nil, err
	}
	tokenInstance, err := GetToken(c.appKey, c.masterSecret)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("messages", string(jsons))
	params.Add("auth_token", tokenInstance.AccessToken)
	bytes, err := doPost(PushHost+MessageUnicastBatchURL, params)
	if err != nil {
		return nil, err
	}
	var result UnicastBatchSendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, errors.New(result.Message)
	}
	return &result, nil
}

// 获取失效的registration_id列表
func (c *OppoPush) FetchInvalidRegidList() (*FetchInvalidRegidListSendResult, error) {
	tokenInstance, err := GetToken(c.appKey, c.masterSecret)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("auth_token", tokenInstance.AccessToken)
	bytes, err := doGet(FeedbackHost+FetchInvalidRegidListURL, "?"+params.Encode())
	if err != nil {
		return nil, err
	}
	var result FetchInvalidRegidListSendResult
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	if result.Code != 0 {
		return nil, errors.New(result.Message)
	}
	return &result, nil
}

func defaultForm(msg *NotificationMessage) url.Values {
	form := url.Values{}
	if msg.AppMessageID != "" {
		form.Add("app_message_id", msg.AppMessageID)
	}
	form.Add("title", msg.Title)
	if msg.SubTitle != "" {
		form.Add("sub_title", msg.SubTitle)
	}
	form.Add("content", msg.Content)
	if msg.ClickActionType > 0 {
		form.Add("click_action_type", strconv.Itoa(msg.ClickActionType))
	}
	if msg.ClickActionType == 1 || msg.ClickActionType == 4 {
		form.Add("click_action_activity", msg.ClickActionActivity)
	}
	if msg.ClickActionType == 2 || msg.ClickActionType == 5 {
		form.Add("click_action_url", msg.ClickActionURL)
	}
	if msg.ActionParameters != "" {
		form.Add("action_parameters", msg.ActionParameters)
	}
	if msg.ShowTimeType > 0 {
		form.Add("show_time_type", strconv.Itoa(msg.ShowTimeType))
	}
	if msg.ShowTimeType > 0 {
		form.Add("show_start_time", strconv.FormatInt(msg.ShowStartTime, 10))
	}
	if msg.ShowTimeType > 0 {
		form.Add("show_end_time", strconv.FormatInt(msg.ShowEndTime, 10))
	}
	if !msg.OffLine {
		form.Add("off_line", strconv.FormatBool(msg.OffLine))
	}
	if msg.OffLine && msg.OffLineTTL > 0 {
		form.Add("off_line_ttl", strconv.Itoa(msg.OffLineTTL))
	}
	if msg.PushTimeType > 0 {
		form.Add("push_time_type", strconv.Itoa(msg.PushTimeType))
	}
	if msg.PushTimeType > 0 {
		form.Add("push_start_time", strconv.FormatInt(msg.PushStartTime, 10))
	}
	if msg.TimeZone != "" {
		form.Add("time_zone", msg.TimeZone)
	}
	if msg.FixSpeed {
		form.Add("fix_speed", strconv.FormatBool(msg.FixSpeed))
	}
	if msg.FixSpeed {
		form.Add("fix_speed_rate", strconv.FormatInt(msg.FixSpeedRate, 10))
	}
	if msg.NetworkType > 0 {
		form.Add("network_type", strconv.Itoa(msg.NetworkType))
	}
	if msg.CallBackURL != "" {
		form.Add("call_back_url", msg.CallBackURL)
	}
	if msg.CallBackParameter != "" {
		form.Add("call_back_parameter", msg.CallBackParameter)
	}
	if msg.ChannelID != "" {
		form.Add("channel_id", msg.ChannelID)
	}
	return form
}

func (u *Message) String() string {
	bytes, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
