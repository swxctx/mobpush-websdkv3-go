package mob_push_sdk

// PushResponse 推送响应
type PushResponse struct {
	// 状态
	Status int `json:"status"`
	Res    struct {
		// 任务ID
		BatchID string      `json:"batchId"`
		Fetched interface{} `json:"fetched"`
		// 卸载
		Uninstalls interface{} `json:"uninstalls"`
		// 关闭通知
		Closes interface{} `json:"closes"`
		// 找不到设备
		NotFounds interface{} `json:"notFounds"`
	} `json:"res"`
	Error interface{} `json:"error"`
}

// PushNumDetails 详情数据
type PushNumDetails struct {
	// 返回的设备数量
	FetchNum int `json:"fetchNum"`
	// 实际下发的设备数量
	DeliverNum int `json:"deliverNum"`
	// 下发失败的数量
	DeliverFailNum int `json:"deliverFailNum"`
	// 下发后回执的数量
	ReportNum int `json:"reportNum"`
	// 下发后回执失败的数量
	ReportFailNum int `json:"reportFailNum"`
	// 点击通知的数量
	ClickNum int `json:"clickNum"`
}

// PushDetailsResponse 推送详情响应
type PushDetailsResponse struct {
	// 状态码
	Status int `json:"status"`
	// 详细信息
	Res DetailsInfo `json:"res"`
}

// detailsResponse
type DetailsInfo struct {
	// iOS 统计
	IOS PushNumDetails `json:"ios"`
	// 安卓统计
	Android PushNumDetails `json:"android"`
	// 通过厂商推送的统计
	Factory PushNumDetails `json:"factory"`
	// 通过MobPush通道下发的统计
	MobPush PushNumDetails `json:"mobpush"`
	// 通过MobPush通道下发iOS的统计
	IOSTcp PushNumDetails `json:"iostcp"`
	// 通过MobPush通道下发Android的统计
	AndroidTcp PushNumDetails `json:"androidtcp"`
	// apns 厂商统计
	Apns PushNumDetails `json:"apns"`
	// 华为 厂商统计
	Huawei PushNumDetails `json:"huawei"`
	// 小米 厂商统计
	Xiaomi PushNumDetails `json:"xiaomi"`
	// fcm 厂商统计
	Fcm PushNumDetails `json:"fcm"`
	// oppo 厂商统计
	Oppo PushNumDetails `json:"oppo"`
	// vivo 厂商统计
	Vivo PushNumDetails `json:"vivo"`
	// sms 厂商统计
	Sms PushNumDetails `json:"sms"`
	// 离线IOS
	OfflineIOS PushNumDetails `json:"offlineIos"`
	// 立项安卓
	OfflineAndroid PushNumDetails `json:"offlineAndroid"`
}
