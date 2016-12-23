package article

import (
	"donn/config"
	"donn/entity"
)

//var dbmap = config.GetDbMap()
var dbmap = config.GetDbMap()

func Get(key int64) *entity.Article {
	if dbmap == nil {
		dbmap = config.GetDbMap()
	}
	obj, _ := dbmap.Get(entity.Article{}, key)
	inst := obj.(*entity.Article)
	return inst
}

func List() []entity.Article {
	if dbmap == nil {
		dbmap = config.GetDbMap()
	}

	var articles []entity.Article
	dbmap.Select(&articles, "select * from articles")
	return articles
}

func Save(inst *entity.Article) {
	if dbmap == nil {
		dbmap = config.GetDbMap()
	}
	dbmap.Insert(inst)
}

func Update(inst *entity.Article) int64 {
	if dbmap == nil {
		dbmap = config.GetDbMap()
	}
	count, _ := dbmap.Update(inst)
	return count
}

func Delete(inst *entity.Article) int64 {
	if dbmap == nil {
		dbmap = config.GetDbMap()
	}
	count, _ := dbmap.Delete(inst)
	return count
}
