package oppopush

// 文档地址 http://storepic.oppomobile.com/openplat/resource/201904/03/OPPO%E6%8E%A8%E9%80%81%E5%B9%B3%E5%8F%B0%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI-V1.6.pdf
const (
	PushHost     = "https://api.push.oppomobile.com"
	FeedbackHost = "https://feedback.push.oppomobile.com"
)

const (
	AuthURL                  = "/server/v1/auth"                                      // 鉴权（auth）
	SaveMessageContentURL    = "/server/v1/message/notification/save_message_content" // 广播推送-保存通知栏消息内容体
	MessageBroadcastURL      = "/server/v1/message/notification/broadcast"            // 广播推送-通知栏消息
	MessageUnicastURL        = "/server/v1/message/notification/unicast"              // 单推-通知栏消息推送
	MessageUnicastBatchURL   = "/server/v1/message/notification/unicast_batch"        // 批量单推-通知栏消息推送
	FetchInvalidRegidListURL = "/server/v1/feedback/fetch_invalid_regidList"          // Feedback-获取失效的 registration_id 列表
)
