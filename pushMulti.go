package mob_push_sdk

type PushMulti struct {
	PushMulti []PushWork `json:"pushWork,omitempty"`
}

type PushWork struct {
	WorkNo string `json:"workno,omitempty"`
	Source string `json:"source,omitempty"`
	// 推送目标
	PushTarget PushTarget `json:"pushTarget,omitempty"`
	// 推送展示细节配置
	PushNotify PushNotify `json:"pushNotify,omitempty"`
	// 运营保障相关配置
	PushOperator PushOperator `json:"pushOperator,omitempty"`
	// link 相关打开配置
	PushForward PushForward `json:"pushForward,omitempty"`
	// 推送回调配置
	PushCallback PushCallback `json:"pushCallback,omitempty"`
}
