package novel

// 作品分类
import (
	"fmt"
	"github.com/junmocsq/yizhen/models/dbcache"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

type Category struct {
	Id        int    `json:"id"`
	Name      string `json:"name" gorm:"size:50;NOT NULL;default:''"`
	Pid       int    `json:"pid" gorm:"NOT NULL;default:0;comment:上级pid"`
	SortIndex int    `json:"sort_index" gorm:"NOT NULL;default:0;comment:排序码，大则权重大"`
	State     uint8  `json:"state" gorm:"NOT NULL;default:1;comment:1 启用 2 停用"`
}

func (Category) TableName() string {
	return "nv_category"
}

type RetCategory struct {
	Category
	Children []RetCategory `json:"children"`
}

type category struct {
}

func NewCategory() *category {
	return &category{}
}

func (c *category) Tag() string {
	return "novel_category"
}

// Add 添加
func (c *category) Add(name string, pid, sortIndex int) int {
	// 检验父级是否存在
	if pid > 0 {
		parentCategory := c.one(pid)
		if parentCategory == nil || parentCategory.Id == 0 {
			logrus.Errorf("novel category pid %d not exists", pid)
			return 0
		}
	}
	// 同一父级下不能重复
	existCategory := c.oneByNameAndPid(name, pid)
	if existCategory != nil && existCategory.Id > 0 {
		return existCategory.Id
	}
	db := dbcache.NewDb()
	var ct = Category{
		Name:      name,
		Pid:       pid,
		SortIndex: sortIndex,
	}
	stmt := db.DryRun().Create(&ct).Statement
	_, err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&ct)
	if err != nil {
		logrus.Errorf("novel category add err:%s", err)
		return 0
	}
	return ct.Id
}

// Open 分类打开
func (c *category) Open(id int) bool {
	return c.updateState(id, 1) > 0
}

// Close 分类关闭
func (c *category) Close(id int) bool {
	return c.updateState(id, 2) > 0
}

func (c *category) updateState(id int, state uint8) int64 {
	db := dbcache.NewDb()
	stmt := db.DryRun().Model(Category{}).Where("id=?", id).Update("state", state).Statement
	n, err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("novel category update err:%s", err)
		return 0
	}
	return n
}

// Update 修改分类名称和权重
func (c *category) Update(id, sortIndex int, name string) bool {
	ca := c.one(id)
	if ca == nil || ca.Id == 0 {
		logrus.Errorf("novel category update err:%d 不存在", id)
		return false
	}
	// 修改类型，需要校验
	if ca.Name != name {
		checkCategory := c.oneByNameAndPid(name, ca.Pid)
		if checkCategory != nil && checkCategory.Id > 0 {
			logrus.Errorf("novel category update err:同一等级下分类 %s 已存在", name)
			return false
		}
	}
	db := dbcache.NewDb()
	stmt := db.DryRun().Model(ca).Updates(map[string]interface{}{"name": name, "sort_index": sortIndex}).Statement
	n, err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).EXEC()
	if err != nil {
		logrus.Errorf("novel category update err:%s", err)
		return false
	}
	return n > 0

}

// 获取所有分类列表，是否需要包含关闭的分类，默认不包含
func (c *category) all(includeClosed ...bool) []Category {
	db := dbcache.NewDb()
	var ret []Category
	isAll := false
	if len(includeClosed) > 0 {
		isAll = includeClosed[0]
	}
	var stmt *gorm.Statement
	if isAll {
		stmt = db.DryRun().Order("sort_index desc,id asc").Find(&ret).Statement
	} else {
		stmt = db.DryRun().Where("state=1").Order("sort_index desc,id asc").Find(&ret).Statement
	}
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&ret)
	if err != nil {
		logrus.Errorf("novel category find all err:%s", err)
		return nil
	}
	return ret
}
func (c *category) FormatList(isAll ...bool) []RetCategory {
	a := c.all(isAll...)
	var f func(ca *RetCategory)
	f = func(ca *RetCategory) {
		for _, v := range a {
			if v.Pid == ca.Id {
				temp := RetCategory{
					v,
					[]RetCategory{},
				}
				f(&temp)
				ca.Children = append(ca.Children, temp)
			}
		}
	}

	var ret []RetCategory
	for _, v := range a {
		if v.Pid == 0 {
			temp := RetCategory{
				v,
				[]RetCategory{},
			}
			f(&temp)
			ret = append(ret, temp)
		}
	}
	return ret
}

func (c *category) List(isAll ...bool) []Category {
	return c.all(isAll...)
}

func (c *category) Print() {
	var fp func(ret []RetCategory, level int)
	fp = func(ret []RetCategory, level int) {
		for _, v := range ret {
			fmt.Println(strings.Repeat("|___", level) + v.Name)
			if len(v.Children) > 0 {
				fp(v.Children, level+1)
			}
		}
	}
	fp(c.FormatList(), 0)
}

// one 获取单条记录
func (c *category) one(id int) *Category {
	db := dbcache.NewDb()
	var ret Category
	stmt := db.DryRun().Where("id=?", id).Find(&ret).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&ret)
	if err != nil {
		logrus.Errorf("novel category find all err:%s", err)
		return nil
	}
	return &ret
}

func (c *category) oneByNameAndPid(name string, pid int) *Category {
	db := dbcache.NewDb()
	var ret Category
	stmt := db.DryRun().Where("name=? AND pid=?", name, pid).Find(&ret).Statement
	err := db.PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&ret)
	if err != nil {
		logrus.Errorf("novel category find all err:%s", err)
		return nil
	}
	return &ret
}
