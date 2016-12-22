package article

import (
	"donn/config"
	"donn/entity"
	"fmt"
)

//var dbmap = config.GetDbMap()
var dbmap = config.GetDbMap()

func Save(inst *entity.Article) {
	if dbmap == nil {
		dbmap = config.GetDbMap()
	}
	dbmap.Insert(inst)
}
