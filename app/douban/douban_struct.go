package douban

type DoubanStruct struct {
	id               string          // 条目id
	title            string          // 中文名
	original_title   string          // 原名
	aka              []string        // 又名
	alt              string          // 条目页URL
	mobile_url       string          // 移动版条目页URL
	rating           rating          // 评分，见附录
	ratings_count    int64           // 评分人数
	wish_count       int64           // 想看人数
	collect_count    int64           // 看过人数
	do_count         int64           // 在看人数，如果是电视剧，默认值为0，如果是电影值为null
	images           images          // 电影海报图，分别提供288px x 465px(大)，96px x 155px(中) 64px x 103px(小)尺寸
	subtype          string          // 条目分类, movie或者tv
	directors        directors       // 导演，数据结构为影人的简化描述，见附录
	casts            []directors     // 主演，最多可获得4个，数据结构为影人的简化描述，见附录
	writers          directors       // 编剧，数据结构为影人的简化描述，见附录
	website          string          // 官方网站
	douban_site      string          // 豆瓣小站
	pubdates         []string        // 如果条目类型是电影则为上映日期，如果是电视剧则为首Ï日期
	mainland_pubdate string          // 大陆上映日期，如果条目类型是电影则为上映日期，如果是电视剧则为首播日期
	pubdate          string          // 兼容性数据，未来会去掉，大陆上映日期，如果条目类型是电影则为上映日期，如果是电视剧则为首播日期
	year             string          // 年代
	languages        []string        // 语言
	durations        []string        // 片长
	genres           []string        // 影片类型，最多提供3个
	countries        []string        // 制片国家/地区
	summary          string          // 简介
	comments_count   int64           // 短评数量
	reviews_count    int64           // 影评数量
	seasons_count    int64           // 总季数(tv only)
	current_season   int64           // 当前季数(tv only)
	episodes_count   int64           // 当前季的集数(tv only)
	schedule_url     string          // 影讯页URL(movie only)
	trailer_urls     []string        // 预告片URL，对高级用户以上开放，最多开放4个地址
	clip_urls        []string        // 片段URL，对高级用户以上开放，最多开放4个地址
	blooper_urls     []string        // 花絮URL，对高级用户以上开放，最多开放4个地址
	photos           []user          // 电影剧照，前10张，见附录
	popular_reviews  popular_reviews // 影评，前10条，影评结构，见附录
}

type rating struct {
	max     int64   // 最高评分
	average float64 // 评分
	stars   int64   // 评星数
	min     int64   // 最低评分
}

type images struct {
	small  string
	large  string
	medium string
}

type directors struct {
	id      string // 影人条目id
	name    string // 中文名
	alt     string // 影人条目URL
	avatars images // 影人头像，分别提供420px x 600px(大)，140px x 200px(中) 70px x 100px(小)尺寸
}

type user struct {
	id    string //图片id
	alt   string //图片展示页url
	icon  string //图片地址，icon尺寸
	image string //图片地址，image尺寸
	thumb string //图片地址，thumb尺寸
	cover string //图片地址，cover尺寸
}

type popular_reviews struct {
	id         string // 影评id
	title      string // 影评名
	alt        string // 影评url
	subject_id string // 条目id
	author     user   // 上传用户，见附录
	rating     rating // 影评评分，见附录
	summary    string // 摘要，100字以内
}
