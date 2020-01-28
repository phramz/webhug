package tpl

import (
	"testing"
)

func TestMustRender(t *testing.T) {
	var tests = []struct {
		name   string
		env    []string
		test   string
		expect string
	}{
		{
			name: "test1",
			env: []string{
				"FOO=bar",
				"BAR==foo",
				"BAZZ=b=ar",
			},
			test:   `hello {{ .Request.Username }}!`,
			expect: "hello victoria!",
		},
		{
			name: "test2",
			env: []string{
				"FOO=bar",
				"BAR==foo",
				"BAZZ=b=ar",
			},
			test:   `hello {{ .Env.FOO }}!`,
			expect: "hello bar!",
		},
		{
			name: "test3",
			env: []string{
				"FOO=bar",
				"BAR==foo",
				"BAZZ=b=ar",
			},
			test:   `{{ index .Request.Header "x-test" }}!`,
			expect: "hello world!",
		},
		{
			name: "test4",
			env: []string{
				"FOO=bar",
				"BAR==foo",
				"BAZZ=b=ar",
			},
			test:   `hey {{ .Request.Json.test.FOO }}!`,
			expect: "hey bar!",
		},
		{
			name: "test5",
			env: []string{
				"FOO=bar",
				"BAR==foo",
				"BAZZ=b=ar",
			},
			test:   `{{ index .Request.Json.more 0 }} {{ index .Request.Json.moore 0 "foo" }}!`,
			expect: "hola bar!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			ctx := AddRequestContext(DefaultContext(tt.env), newRequest(tt.name))
			actual := MustRender(tt.test, ctx)

			if actual != tt.expect {
				t.Errorf("rendered template does look like expected for wanted %q, got: %q", tt.expect, actual)
			}
		})
	}
}
