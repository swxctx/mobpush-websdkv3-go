package mob_push_sdk

// PushFactoryExtra
type PushFactoryExtra struct {
	// 小米厂商特殊配置
	XiaomiExtra *XiaomiExtra `json:"xiaomiExtra,omitempty"`
	// OPPO厂商特殊配置
	OppoExtra *OppoExtra `json:"oppoExtra,omitempty"`
	// VIVO厂商特殊配置
	VivoExtra *VivoExtra `json:"vivoExtra,omitempty"`
	// 华为厂商特殊配置
	HuaweiExtra *HuaweiExtra `json:"huaweiExtra,omitempty"`
	// 荣耀厂商特殊配置
	HonorExtra *HonorExtra `json:"honorExtra,omitempty"`
}

type (
	XiaomiExtra struct {
		ChannelId string `json:"channel_id,omitempty"`
	}

	OppoExtra struct {
		ChannelId string `json:"channel_id,omitempty"`
	}

	VivoExtra struct {
		// VIVO消息类型 0：运营类型消息 ；1：系统类型消息
		Classification string `json:"classification,omitempty"`
	}

	HuaweiExtra struct {
		// 消息类型- LOW：资讯营销类 - NORMAL：服务与通讯类 注：资讯营销类的消息提醒方式为静默通知，仅在下拉通知栏展示。 服务与通讯类的消息提醒方式为锁屏+铃声+震动
		Importance string `json:"importance,omitempty"`
		// 作用一：完成自分类权益申请后，用于标识消息类型，确定消息提醒方式，对特定类型消息加快发送，
		// 取值如下：
		//	IM：即时聊天
		//	VOIP：音视频通话
		//	SUBSCRIPTION：订阅
		//	TRAVEL：出行
		//	HEALTH：健康
		//	WORK：工作事项提醒
		//	ACCOUNT：帐号动态
		//	EXPRESS：订单&物流
		//	FINANCE：财务
		//	DEVICE_REMINDER：设备提醒
		//	SYSTEM_REMINDER：系统提示
		//	MAIL：邮件
		//	PLAY_VOICE：语音播报（仅透传消息支持）
		//	MARKETING：内容推荐、新闻、财经动态、生活资讯、社交动态、调研、产品促销、功能推荐、运营活动（仅对内容进行标识，不会加快消息发送）
		// 作用二：申请特殊权限后，用于标识高优先级透传场景，取值如下： VOIP：音视频通话 PLAY_VOICE：语音播报
		Category string `json:"category,omitempty"`
	}

	HonorExtra struct {
		/*
			消息类型
			- LOW：资讯营销类
			- NORMAL：服务与通讯类
			注：资讯营销类的消息默认展示方式为静默通知，仅在下拉通知栏展示。 服务与通讯类的消息默认展示方式为锁屏展示+下拉通知栏展示
		*/
		Importance string `json:"importance,omitempty"`
	}
)
