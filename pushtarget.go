package mob_push_sdk

type TargetType int

const (
	// 广播
	TARGET_ALL TargetType = 1
	// 别名
	TARGET_ALIAS TargetType = 2
	// 标签
	TARGET_TAGS TargetType = 3
	// regid
	TARGET_RIDS TargetType = 4
	// 地理位置
	TARGET_AREAS TargetType = 9
)

type PushTarget struct {
	// 目标类型：1广播；2别名；3标签；4regid；5地理位置；6用户分群；9复杂地理位置推送
	Target int `json:"target,omitempty"`
	// target:3 => 设置推送标签集合["tag1","tag2"]集合元素限制1000个以内
	Tags []string `json:"tags,omitempty"`
	// target:3 => 标签组合方式：1并集；2交集；3补集(3暂不考虑)
	TagsType string `json:"tagsType,omitempty"`
	// arget:2 => 设置推送别名集合["alias1","alias2"]集合元素限制1000个以内
	Alias []string `json:"alias,omitempty"`
	// target:4 => 设置推送Registration Id集合["id1","id2"]集合元素限制1000个以内
	Rids []string `json:"rids,omitempty"`

	// target:6 => 用户分群ID（暂时不支持）
	Block string `json:"block,omitempty"`
	// target:5 => 推送地理位置 城市，地理位置推送时（示例：上海市注意要带“市”），city, province, country 必须有一个不为空
	City string `json:"city,omitempty"`
	// target:5 => 推送地理位置 国家，地理位置推送时，city, province, country 必须有一个不为空
	Country string `json:"country,omitempty"`
	// target:5 => 推送地理位置 省份，地理位置推送时（示例：浙江省注意要带“省”），city, province, country 必须有一个不为空
	Province string `json:"province,omitempty"`
	// target:9 时必传，复杂地理位置
	PushAreas *PushAreas `json:"pushAreas,omitempty"`

	// 指定推送的包名列表
	AppPackages []string `json:"appPackages,omitempty"`
}

type PushLabel struct {
	LabelIds []string `json:"labelIds,omitempty"`
	MobId    string   `json:"mobId,omitempty"`
	Type     int      `json:"type,omitempty"`
}

type PushAreas struct {
	// 国家列表
	Countries []PushCountry `json:"countries,omitempty"`
}

type PushCountry struct {
	// 国家名称
	Country string `json:"country,omitempty"`
	// 指定需要推送的省份列表
	Provinces []PushProvince `json:"provinces,omitempty"`
}

type PushProvince struct {
	// 省份名称
	Province string `json:"province,omitempty"`
	// 需要推送的城市列表
	Cities []string `json:"cities,omitempty"`
	// 指定不需要推送的城市列表，当指定[cities]时，这个字段不起作用
	ExcludeCities []string `json:"excludeCities,omitempty"`
	// 指定不需要推送的省份列表，当设置[provinces]时，这个不起作用
	ExcludeProvinces []string `json:"excludeProvinces,omitempty"`
}

func (p *Push) setTarget(targetType TargetType) *Push {
	p.PushTarget.Target = int(targetType)
	return p
}

func (p *Push) setAlias(alias []string) *Push {
	p.PushTarget.Alias = alias
	return p
}

func (p *Push) setTags(tags []string) *Push {
	p.PushTarget.Tags = tags
	return p
}

func (p *Push) setRids(rids []string) *Push {
	p.PushTarget.Rids = rids
	return p
}

func (p *Push) setPushAreas(pushAreas PushAreas) *Push {
	p.PushTarget.PushAreas = &pushAreas
	return p
}
