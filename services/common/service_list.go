package common

import (
	"gorm.io/gorm"
	"server/global"
	"server/models"
)

type Option struct {
	models.PageInfo
	Debug bool //看日志，如果debug等于true就去看日志
}

// CommList 将列表和总页数返回出去
// CommList 完成一个简单的分页，问题是怎么调用这个方法（函数）。
func CommList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	// 设置一个默认值
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
		// option.Sort = "created_at asc"  // 默认按照时间往后排
	}
	query := DB.Where(model) //这样就能将model里面的参数查进去
	// 列表页会有分页
	//DB.Model(model).Count(&count)
	count = query.Select("id").Find(&list).RowsAffected // 查所有的数据,数据量大的话，这里会有点慢
	// 这里的query会受上面query的影响，需要手动复位
	query = DB.Where(model)                    // 这里如果没有这一句query就被前面的覆盖，加这一句是为了复位
	offset := option.Limit * (option.Page - 1) //通过当前的页数计算得出
	if offset < 0 {
		offset = 0
	}
	// 默认按照时间往前排 Order(option.Sort)
	if option.Limit == 0 {
		err = query.Offset(offset).Order(option.Sort).Find(&list).Error // 完成分页 （Limit(1).Offset(1)只拿一个从第一页开始）将参数一一映射上去
	} else {
		err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error // 完成分页 （Limit(1).Offset(1)只拿一个从第一页开始）将参数一一映射上去
	}

	return list, count, err
}
