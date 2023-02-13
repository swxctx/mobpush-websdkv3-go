package mob_push_sdk

type PushCallback struct {
	// 回调地址
	Url string `json:"url,omitempty"`
	// JSON对象自定义参数例： { "key": "value" }
	Params map[string]string `json:"params,omitempty"`
}
