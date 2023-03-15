package encrypt

import "testing"

func Test(t *testing.T) {
	items := []struct {
		content string
		salt    string
		tail    uint
		result  string
	}{
		{
			content: "",
			salt:    "8$afzFlfM1y7NTN2z^J3nsfoAt7dW%VV",
			tail:    5,
			result:  "28ea087b2da70a2112ffec079a78abe6",
		},

		{
			content: "hello world",
			tail:    15,
			result:  "ac8220c5f01121a5343992492a4e42e0",
		},

		{
			content: "18048574657",
			tail:    27,
			result:  "7ea8d4501df29aca658d26e8cd54f467",
		},

		{
			content: "jztech",
			salt:    "w#O71y3U2jXIC1FxW!$$&3Jb!qnPZLNT",
			tail:    21,
			result:  "b908aa5ea83db74e8e4ac7d44842197f",
		},
	}
	for _, item := range items {
		t.Run(item.content, func(t *testing.T) {
			result := Encrypt(item.content, item.salt, item.tail)
			if result != item.result {
				t.Errorf("invalid result, content: [%v], salt: [%v], expect result: [%v], actually result: [%v]", item.content, item.salt, item.result, result)
			}
		})
	}
}
