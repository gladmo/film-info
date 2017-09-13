package douban

type ErrorRes struct {
	Msg     string
	Code    int64
	request string
}

type DoubanStruct struct {
	Id               string          // 条目id
	Title            string          // 中文名
	Original_title   string          // 原名
	Aka              []string        // 又名
	Alt              string          // 条目页URL
	Mobile_url       string          // 移动版条目页URL
	Rating           rating          // 评分，见附录
	Ratings_count    int64           // 评分人数
	Wish_count       int64           // 想看人数
	Collect_count    int64           // 看过人数
	Do_count         int64           // 在看人数，如果是电视剧，默认值为0，如果是电影值为null
	Images           images          // 电影海报图，分别提供288px x 465px(大)，96px x 155px(中) 64px x 103px(小)尺寸
	Subtype          string          // 条目分类, movie或者tv
	Directors        directors       // 导演，数据结构为影人的简化描述，见附录
	Casts            []directors     // 主演，最多可获得4个，数据结构为影人的简化描述，见附录
	Writers          directors       // 编剧，数据结构为影人的简化描述，见附录
	Website          string          // 官方网站
	Douban_site      string          // 豆瓣小站
	Pubdates         []string        // 如果条目类型是电影则为上映日期，如果是电视剧则为首Ï日期
	Mainland_pubdate string          // 大陆上映日期，如果条目类型是电影则为上映日期，如果是电视剧则为首播日期
	Pubdate          string          // 兼容性数据，未来会去掉，大陆上映日期，如果条目类型是电影则为上映日期，如果是电视剧则为首播日期
	Year             string          // 年代
	Languages        []string        // 语言
	Durations        []string        // 片长
	Genres           []string        // 影片类型，最多提供3个
	Countries        []string        // 制片国家/地区
	Summary          string          // 简介
	Comments_count   int64           // 短评数量
	Reviews_count    int64           // 影评数量
	Seasons_count    int64           // 总季数(tv only)
	Current_season   int64           // 当前季数(tv only)
	Episodes_count   int64           // 当前季的集数(tv only)
	Schedule_url     string          // 影讯页URL(movie only)
	Trailer_urls     []string        // 预告片URL，对高级用户以上开放，最多开放4个地址
	Clip_urls        []string        // 片段URL，对高级用户以上开放，最多开放4个地址
	Blooper_urls     []string        // 花絮URL，对高级用户以上开放，最多开放4个地址
	Photos           []user          // 电影剧照，前10张，见附录
	Popular_reviews  popular_reviews // 影评，前10条，影评结构，见附录
}

type rating struct {
	Max     int64   // 最高评分
	Average float64 // 评分
	Stars   int64   // 评星数
	Min     int64   // 最低评分
}

type images struct {
	Small  string
	Large  string
	Medium string
}

type directors struct {
	Id      string // 影人条目id
	Name    string // 中文名
	Alt     string // 影人条目URL
	Avatars images // 影人头像，分别提供420px x 600px(大)，140px x 200px(中) 70px x 100px(小)尺寸
}

type user struct {
	Id    string //图片id
	Alt   string //图片展示页url
	Icon  string //图片地址，icon尺寸
	Image string //图片地址，image尺寸
	Thumb string //图片地址，thumb尺寸
	Cover string //图片地址，cover尺寸
}

type popular_reviews struct {
	Id         string // 影评id
	Title      string // 影评名
	Alt        string // 影评url
	Subject_id string // 条目id
	Author     user   // 上传用户，见附录
	Rating     rating // 影评评分，见附录
	Summary    string // 摘要，100字以内
}

// douban search
type DoubanSearchOne struct {
	Episode   string
	Img       string
	Title     string
	Url       string
	Type      string
	Year      string
	Sub_title string
	Id        string
}
