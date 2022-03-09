package novel

import (
	"testing"
)

func TestNote(t *testing.T) {

}

func TestCheckRuneString(t *testing.T) {
	a := "小丫小二郎"
	arr := []struct {
		str    string
		resStr string
		isTrue bool
	}{
		{a[3:15], a[3:15], true},
		{a[1:15], a[3:15], false},
		{a[1:14], a[3:12], false},
	}
	for _, v := range arr {
		s, f := CheckRuneString(v.str)
		if s != v.resStr || f != v.isTrue {
			t.Errorf("want %s,flag:%v got %s,flag:%v", v.resStr, v.isTrue, s, f)
		}
	}
}
func TestFuzzyStringMatching(t *testing.T) {
	str := `新华社北京3月8日电 综合新华社驻外记者报道：俄罗斯代表团团长梅金斯基7日说，在白俄罗斯境内举行的第三轮俄乌谈判未达预期效果，但双方的谈判还会继续。乌克兰谈判代表波多利亚克同日表示，乌俄谈判双方就人道主义走廊问题取得一些成果。

　　——俄罗斯代表团团长、总统助理梅金斯基7日说，在白俄罗斯境内别洛韦日森林举行的第三轮俄乌谈判未达预期效果，但双方的谈判还会继续。俄方希望双方能签署至少一项协议，但乌方表示将把谈判文件都带回去研究，现场未签署任何文件。

　　——乌克兰谈判代表、乌总统办公室顾问波多利亚克7日在社交媒体上说，第三轮乌俄谈判双方就人道主义走廊问题取得一些积极成果。双方将就停火、休战和结束敌对行动等问题继续进行密集磋商。

　　——据俄新社7日报道，俄总统新闻秘书佩斯科夫表示，俄方已告知乌方，若其履行俄方提出的条件，包括乌克兰必须修改宪法并放弃加入“任何集团”、承认克里米亚归属俄罗斯、承认顿涅茨克和卢甘斯克为独立共和国，那么特别军事行动随时可以停止。

　　——俄罗斯总统普京7日与欧洲理事会主席米歇尔通电话讨论乌克兰局势。普京在通话中介绍了俄方开展特别军事行动的立场，并与米歇尔详细讨论了乌当前的人道主义局势。

　　——欧盟外交与安全政策高级代表博雷利7日表示，欧盟已开始向乌克兰和摩尔多瓦提供1亿欧元，用于人道主义援助。

　　——德国总理朔尔茨7日发表声明说，过去几个月，德国政府一直在同欧盟内外的伙伴商讨俄罗斯能源的替代品，但这不可能在短期内实现，所以德国将继续与俄罗斯在能源供应领域进行合作。他说，来自俄罗斯的能源对欧洲至关重要，目前没有其他方法来保障欧洲的供暖、电力和交通能源供应。

　　——据俄罗斯国防部7日公布的材料，美国对在乌克兰境内的生物实验室资助超2亿美元。俄军辐射、化学和生物防护部队司令基里洛夫说，乌克兰境内已形成一个由30多个生物实验室组成的网络，这些“实验室工作的定购人”是美国国防部减少威胁局。`
	//t.Log(3333_4444)
	subStr := "据俄新社7日报道，俄总统新闻秘书佩斯科夫"
	_, _, matching := FuzzyStringMatching(str, subStr)
	if matching != 1 {
		t.Error("matching failed")
	}
	subStr = "——欧盟外交与安全政策高级代表博雷利7日表示，欧盟已开始向乌克兰和摩尔多瓦提供1亿欧元，用于人道主义援助。——德国总理朔尔茨7日发表声明说，过去几个月，德国政府一直在同欧盟内外的伙伴商讨俄罗斯能源的替代品，但这不可能在短期内实现，所以德国将继续与俄罗斯在能源供应领域进行合作。他说，来自俄罗斯的能源对欧洲至关重要，目前没有其他方法来保障欧洲的供暖、电力和交通能源供应。"
	_, _, matching = FuzzyStringMatching(str, subStr)
	if matching < 0.9 {
		t.Error("matching failed")
	}
}
