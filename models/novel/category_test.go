package novel

import (
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestCategory(t *testing.T) {
	arr := make(map[string]map[string][]string)
	arr["男生"] = map[string][]string{
		"玄幻": []string{"东方玄幻", "异世大陆", "王朝争霸", "高武世界"},
		"武侠": []string{"古典仙侠", "现代修真", "新派武侠", "传统武侠", "奇幻修真"},
		"科幻": []string{"进化变异", "星际争霸", "游戏攻略", "虚拟网游", "游戏小说", "科技时代"},
		"历史": []string{"穿越历史", "架空历史", "历史传记"},
	}

	arr["女生"] = map[string][]string{
		"言情": []string{"青春校园", "婚恋家庭", "白领职场", "都市重生", "豪门总裁"},
		"武侠": []string{"奇幻玄幻", "灵异悬疑", "科幻游戏", "武侠仙侠"},
		"古代": []string{"历史传奇", "穿越时空", "架空历史"},
	}

	c := NewCategory()
	if len(c.all(true)) == 0 {
		for k, v := range arr {
			pid := c.Add(k, 0, rand.Intn(999))
			for _k, _v := range v {
				pid := c.Add(_k, pid, rand.Intn(999))
				for _, __v := range _v {
					c.Add(__v, pid, rand.Intn(999))
				}
			}
		}
	}
	id := c.Add("category_test", 0, 989)
	ca := c.one(id)
	Convey("Category", t, func() {
		Convey("Add", func() {
			So(ca, ShouldNotBeNil)
			So(ca.Id, ShouldEqual, id)
		})
		Convey("Close", func() {
			So(c.Close(id), ShouldBeTrue)
			So(c.one(id).State, ShouldEqual, 2)
		})
		Convey("Open", func() {
			So(c.Open(id), ShouldBeTrue)
			So(c.one(id).State, ShouldEqual, 1)
		})

		Convey("Update", func() {
			So(c.Update(id, 889, "update_temp"), ShouldBeTrue)
			So(c.Update(id, 887, "category_test"), ShouldBeTrue)
		})

		Convey("List", func() {
			So(c.List(), ShouldNotBeNil)
			So(c.FormatList(), ShouldNotBeNil)
		})
	})
	//c.Print()
}
