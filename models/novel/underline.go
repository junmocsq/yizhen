package novel

import (
	"github.com/agnivade/levenshtein"
	"strings"
	"unicode/utf8"
)

// 每用户划线。取消划线时，UnderlineContentId为0；评论删除时，需要维护CommentId字段。

type Underline struct {
	Id                 int    `json:"id"`
	Nid                int    `json:"nid"`
	Cid                int    `json:"cid"`
	Uid                int    `json:"uid"`
	UnderlineContentId int    `json:"underline_content_id"`
	CommentId          int    `json:"comment_id" gorm:"comment:评论id"`
	HasNote            uint8  `json:"has_note" gorm:"comment:0 没有笔记 1 有笔记;NOT NULL;default:0"`
	Note               string `json:"content" gorm:"comment:笔记;size:3000;NOT NULL;default:''"`
	CreatedAt          int64  `json:"created_at" gorm:"comment:autoCreateTime"`
}

func (Underline) TableName() string {
	return "nv_underline"
}

// CheckRuneString 校验字符串是不是正确的unicode编码，返回正确编码字符串和是否正确编码
func CheckRuneString(str string) (string, bool) {
	ret := make([]rune, utf8.RuneCountInString(str))
	var n int
	flag := true
	for _, v := range str {
		if v == 0xfffd {
			flag = false
			continue
		}
		ret[n] = v
		n++
	}
	return string(ret[:n]), flag
}

// FuzzyStringMatching 字符串模糊匹配
func FuzzyStringMatching(str, substr string) (start, end int, matching float64) {
	// 准确匹配
	index := strings.Index(str, substr)
	if index >= 0 {
		return index, index + len(substr), 1
	}

	// 模糊匹配
	totalLength := utf8.RuneCountInString(str)
	subLength := utf8.RuneCountInString(substr)
	newStr := []rune(str)
	minDistance := subLength
	start = -1
	end = -1
	// 相等长度比较
	for i := 0; i < totalLength-subLength; i++ {
		testStr := string(newStr[i : i+subLength])
		distance := levenshtein.ComputeDistance(testStr, substr)
		if distance < minDistance {
			minDistance = distance
			start = i
			end = i + subLength
		}
	}
	if start < 0 || end <= start {
		return 0, 0, 0
	}
	// 在等长最好匹配的地方再向后移动10分之一的子串长度比较，进一步优化匹配
	newEnd := end + subLength/10
	for i := end; i < newEnd; i++ {
		testStr := string(newStr[start:i])
		distance := levenshtein.ComputeDistance(testStr, substr)
		if distance < minDistance {
			minDistance = distance
			end = i
		}
	}
	patStr := string(newStr[start:end])
	index = strings.Index(str, patStr)
	return index, index + len(patStr), 1 - float64(minDistance)/float64(end-start)
}
