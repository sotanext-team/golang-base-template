package models

type Shop struct {
	Model

	Name          string `json:"name" gorm:"size:500"`
	DefaultDomain string `json:"defaultDomain" gorm:"size:500;index:shop_default_domain_idx,unique"`
	CustomDomain  string `json:"customDomain" `
	ShopName      string `json:"shopName" `

	// Users          []*User          `json:"users" gorm:"many2many:user_shops"`
	// ShopMetas      []*ShopMeta      `json:"shopMetas"`
	// ShopEvents     []*ShopEvent     `json:"shopEvents"`
	// Themes         []*Theme         `json:"theme"`
	// ThemeTemplates []*ThemeTemplate `json:"themeTemplates"`
}

type ShopToken struct {
	ID            string `json:"id"`
	DefaultDomain string `json:"defaultDomain"`
	UserID        uint   `json:"userId"`
}
