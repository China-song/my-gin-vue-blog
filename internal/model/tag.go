package model

import "gorm.io/gorm"

// Tag 文章标签
// many2many 一个标签下有多个文章 一篇文章有多个标签
// TODO: 掌握many2many gorm会创建关联表吗 2024/03/25 16:42
type Tag struct {
	Model
	Name string `gorm:"unique;type:varchar(20);not null" json:"name"`

	Articles []*Article `gorm:"many2many:article_tag;" json:"articles,omitempty"`
}

type TagVO struct {
	Tag
	ArticleCount int `json:"article_count"`
}

// GetTagList 还需要统计每个tag下的文章数目
func GetTagList(db *gorm.DB, page, size int, keyword string) (list []TagVO, total int64, err error) {
	db = db.Table("tag t").
		Joins("LEFT JOIN article_tag at ON t.id = at.tag_id").
		Select("t.id", "t.name", "COUNT(at.article_id) AS article_count", "t.created_at", "t.updated_at")

	if keyword != "" {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}

	result := db.
		Group("t.id").Order("t.updated_at DESC").
		Count(&total).
		Scopes(Paginate(page, size)).
		Find(&list)

	return list, total, result.Error
}

func SaveOrUpdateTag(db *gorm.DB, id int, name string) (*Tag, error) {
	tag := Tag{
		Model: Model{ID: id},
		Name:  name,
	}

	var result *gorm.DB
	if id > 0 {
		result = db.Updates(tag)
	} else {
		result = db.Create(&tag)
	}

	return &tag, result.Error
}

func GetTagOption(db *gorm.DB) ([]OptionVO, error) {
	list := make([]OptionVO, 0)
	result := db.Model(&Tag{}).Select("id", "name").Find(&list)
	return list, result.Error
}
