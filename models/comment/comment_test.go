package comment

import "testing"

func TestComment_Add(t *testing.T) {
	NewComment().Add(1, 1, 0, 1, "Hello World!")
	NewComment().Add(2, 1, 10, 1, "Hello zxf!")
	NewComment().Add(3, 1, 10, 1, "Hello lxq!")
}
