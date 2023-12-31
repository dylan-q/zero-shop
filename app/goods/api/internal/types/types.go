// Code generated by goctl. DO NOT EDIT.
package types

type ListReq struct {
	Page     int64 `form:"page"`
	PageSize int64 `form:"page_size"`
}

type DetailReq struct {
	Id int64 `form:"id"`
}

type ListResp struct {
	Total int64       `json:"total"`
	Data  []GoodsInfo `json:"data"`
}

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GoodsInfo struct {
	Id          int64    `json:"id"`
	GoodsName   string   `json:"goods_name"`   // 商品名称
	CategoryId  int64    `json:"category_id"`  // 商品分类id
	Category    Category `json:"category"`     //分类
	OnSale      int64    `json:"on_sale"`      // 是否在销售 1 是 0 否
	ShipFree    int64    `json:"ship_free"`    // 是否免运费 1 是 0 否
	IsNew       int64    `json:"is_new"`       // 是否最新 1 是 0 否
	IsHot       int64    `json:"is_hot"`       // 是否最热 1 是 0 否
	ClickNum    int64    `json:"click_num"`    // 查看次数
	SoldNum     int64    `json:"sold_num"`     // 售出数量
	FavNum      int64    `json:"fav_num"`      // 收藏数量
	MarketPrice float64  `json:"market_price"` // 市场价
	ShopPrice   float64  `json:"shop_price"`   // 售价
	GoodsBrief  string   `json:"goods_brief"`  // 商品简介
	GoodsDetail string   `json:"goods_detail"` // 商品详情
	CreateTime  int64    `json:"create_time"`
	UpdateTime  int64    `json:"update_time"`
}
