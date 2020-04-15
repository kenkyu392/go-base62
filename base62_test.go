package base62

import "testing"

func Test(t *testing.T) {
	t.Run("case=01", func(t *testing.T) {
		var src = []byte("hello")
		const want = "7tQLFHz"
		if got := StdEncoding.EncodeToString(src); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=02", func(t *testing.T) {
		var src = []byte("مرحبا بالعالم!")
		const want = "e0ioTSb2h3yqg1CNlbdJ4SqQMKXJQYZz7jN"
		if got := StdEncoding.EncodeToString(src); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})

	t.Run("case=03", func(t *testing.T) {
		var src = []byte("<b>HELLO</b>WORLD!")
		const want = "VKaviXC9ehWqj2drr1wBIoe9"
		if got := StdEncoding.EncodeToString(src); got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})
}
