package crypto

import "testing"

func TestGithubSign(t *testing.T) {
	var tests = []struct {
		signature string
		secret    string
		payload   string
	}{
		{
			"sha1=6d499ea5273857814a448cd5a1e961b20fabf5a4",
			"231a9f49a1592ec348663c37e53885164fe840a0b5e464e017c2ed6f06686",
			`{"received_events_url":"https://api.github.com/users/phramz/received_events","type":"User","site_admin":false}`,
		},
		{
			"sha1=fbdb1d1b18aa6c08324b7d64b71fb76370690e1d",
			"",
			"",
		},
		{
			"sha1=c3c4bf4cfa0e12f325e6f9f429309c70700f7ae8",
			"231a9f49a1592ec348663c37e53885164fe840a0b5e464e017c2ed6f06686",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.signature, func(t *testing.T) {
			s := GithubSign([]byte(tt.payload), []byte(tt.secret))
			if s != tt.signature {
				t.Errorf("got %q, want %q", s, tt.signature)
			}
		})
	}
}
