package tpl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestAddRequestContext(t *testing.T) {
	var tests = []struct {
		name   string
		env    []string
		expect string
	}{
		{
			name: "test1",
			env: []string{
				"FOO=bar",
				"BAR==foo",
				"BAZZ=b=ar",
			},
			expect: `{"Env":{"BAR":"=foo","BAZZ":"b=ar","FOO":"bar"},"Webhook":{"Name":"","Format":""},"Request":{"Body":"{\"test\":{\"BAR\":\"=foo\",\"BAZZ\":\"b=ar\",\"FOO\":\"bar\"},\"more\":[\"hola\",\"hello\",\"hallo\"],\"moore\":[{\"foo\": \"bar\", \"bar\": \"bazz\"}, {}]}","Method":"POST","Json":{"moore":[{"bar":"bazz","foo":"bar"},{}],"more":["hola","hello","hallo"],"test":{"BAR":"=foo","BAZZ":"b=ar","FOO":"bar"}},"Uri":"/test1/webhug?name=test1\u0026e[]=m\u0026e[]=c^2","Host":"example.com","RemoteAddr":"127.0.0.1:12345","Query":"name=test1\u0026e[]=m\u0026e[]=c^2","Scheme":"https","Username":"victoria","Password":"secret","Header":{"authorization":"Basic dmljdG9yaWE6c2VjcmV0","content-type":"application/vnd.api+json","cookie":"session=test; test=hello","x-test":"hello world","x-test-2":"hello dude"},"Get":{"e[]":["m","c^2"],"name":["test1"]},"Cookie":{"session":{"Name":"session","Value":"test","Path":"","Domain":"","Expires":"0001-01-01T00:00:00Z","RawExpires":"","MaxAge":0,"Secure":false,"HttpOnly":false,"SameSite":0,"Raw":"","Unparsed":null},"test":{"Name":"test","Value":"hello","Path":"","Domain":"","Expires":"0001-01-01T00:00:00Z","RawExpires":"","MaxAge":0,"Secure":false,"HttpOnly":false,"SameSite":0,"Raw":"","Unparsed":null}}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := AddRequestContext(DefaultContext(tt.env), newRequest(tt.name))
			j, _ := json.Marshal(s)

			if fmt.Sprintf("%s", j) != tt.expect {
				t.Errorf("context does look like expected for %q, got: %s", tt.name, j)
			}
		})
	}
}

func newRequest(name string) *http.Request {
	turi := fmt.Sprintf("/%s/webhug?name=%s&e[]=m&e[]=c^2", name, name)
	tbody := `{"test":{"BAR":"=foo","BAZZ":"b=ar","FOO":"bar"},"more":["hola","hello","hallo"],"moore":[{"foo": "bar", "bar": "bazz"}, {}]}`
	trq, _ := http.NewRequest("POST", fmt.Sprintf("https://victoria:secret@example.com%s", turi), strings.NewReader(tbody))
	trq.SetBasicAuth("victoria", "secret")
	trq.AddCookie(&http.Cookie{Name: "session", Value: "test"})
	trq.AddCookie(&http.Cookie{Name: "test", Value: "hello"})
	trq.Header.Add("content-type", "application/vnd.api+json")
	trq.Header.Add("x-test", "hello dude")
	trq.Header.Add("x-test", "hello world")
	trq.Header.Add("x-test-2", "hello dude")
	trq.RemoteAddr = "127.0.0.1:12345"
	trq.Host = "example.com"
	trq.RequestURI = turi

	return trq
}
