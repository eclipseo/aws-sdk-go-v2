// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	smithytime "github.com/awslabs/smithy-go/time"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_XmlTimestamps_awsRestxmlSerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *XmlTimestampsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Tests how normal timestamps are serialized
		"XmlTimestamps": {
			Params: &XmlTimestampsInput{
				Normal: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <normal>2014-04-29T18:30:38Z</normal>
			</XmlTimestampsInputOutput>
			`))
			},
		},
		// Ensures that the timestampFormat of date-time works like normal timestamps
		"XmlTimestampsWithDateTimeFormat": {
			Params: &XmlTimestampsInput{
				DateTime: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <dateTime>2014-04-29T18:30:38Z</dateTime>
			</XmlTimestampsInputOutput>
			`))
			},
		},
		// Ensures that the timestampFormat of epoch-seconds works
		"XmlTimestampsWithEpochSecondsFormat": {
			Params: &XmlTimestampsInput{
				EpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <epochSeconds>1398796238</epochSeconds>
			</XmlTimestampsInputOutput>
			`))
			},
		},
		// Ensures that the timestampFormat of http-date works
		"XmlTimestampsWithHttpDateFormat": {
			Params: &XmlTimestampsInput{
				HttpDate: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/XmlTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareXMLReaderBytes(actual, []byte(`<XmlTimestampsInputOutput>
			    <httpDate>Tue, 29 Apr 2014 18:30:38 GMT</httpDate>
			</XmlTimestampsInputOutput>
			`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.XmlTimestamps(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if actualReq.Body != nil {
				defer actualReq.Body.Close()
			}
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_XmlTimestamps_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *XmlTimestampsOutput
	}{
		// Tests how normal timestamps are serialized
		"XmlTimestamps": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <normal>2014-04-29T18:30:38Z</normal>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				Normal: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of date-time works like normal timestamps
		"XmlTimestampsWithDateTimeFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <dateTime>2014-04-29T18:30:38Z</dateTime>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				DateTime: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of epoch-seconds works
		"XmlTimestampsWithEpochSecondsFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <epochSeconds>1398796238</epochSeconds>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				EpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of http-date works
		"XmlTimestampsWithHttpDateFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlTimestampsInputOutput>
			    <httpDate>Tue, 29 Apr 2014 18:30:38 GMT</httpDate>
			</XmlTimestampsInputOutput>
			`),
			ExpectResult: &XmlTimestampsOutput{
				HttpDate: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, vs := range c.Header {
					for _, v := range vs {
						w.Header().Add(k, v)
					}
				}
				if len(c.BodyMediaType) != 0 && len(w.Header().Values("Content-Type")) == 0 {
					w.Header().Set("Content-Type", c.BodyMediaType)
				}
				if len(c.Body) != 0 {
					w.Header().Set("Content-Length", strconv.Itoa(len(c.Body)))
				}
				w.WriteHeader(c.StatusCode)
				if len(c.Body) != 0 {
					if _, err := io.Copy(w, bytes.NewReader(c.Body)); err != nil {
						t.Errorf("failed to write response body, %v", err)
					}
				}
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params XmlTimestampsInput
			result, err := client.XmlTimestamps(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
