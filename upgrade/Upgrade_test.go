package upgrade

import "testing"

func TestIsPass(t *testing.T) {
	cases := []struct {
		userid   int
		excepted bool
	}{
		{1000, false},
		{11111, true},
		{0, false},
	}

	for _, c := range cases {
		result := IsPass(c.userid)
		if result != c.excepted {
			t.Fatalf("userid: %d,  execpted:%v, result:%v", c.userid, c.excepted, result)
		}
	}
}
