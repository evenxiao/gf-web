package define

// Service查询栏目树形列表
type CategoryApiGetTreeReq struct {
	// 栏目类型：topic, question, article
	// 当传递空时表示获取所有类型的栏目
	ContentType string
}
