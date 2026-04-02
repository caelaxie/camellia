package camellia

import "testing"

func TestSuggestedName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		input      string
		want       string
		wantChange bool
	}{
		{name: "empty", input: "", want: "", wantChange: false},
		{name: "api error", input: "APIError", want: "ApiError", wantChange: true},
		{name: "full uppercase abbreviation", input: "URL", want: "Url", wantChange: true},
		{name: "user id", input: "UserID", want: "UserId", wantChange: true},
		{name: "http client", input: "HTTPClient", want: "HttpClient", wantChange: true},
		{name: "mixed lower prefix", input: "userID", want: "userId", wantChange: true},
		{name: "trailing abbreviation", input: "MyURL", want: "MyUrl", wantChange: true},
		{name: "digit after abbreviation", input: "HTTP2Client", want: "Http2Client", wantChange: true},
		{name: "digit before trailing abbreviation", input: "Version2API", want: "Version2Api", wantChange: true},
		{name: "abbreviations at both ends", input: "URLParserAPI", want: "UrlParserApi", wantChange: true},
		{name: "valid camel case", input: "ApiError", want: "ApiError", wantChange: false},
		{name: "single uppercase", input: "Client", want: "Client", wantChange: false},
		{name: "oauth style", input: "OAuthToken", want: "OAuthToken", wantChange: false},
		{name: "oauth plus trailing abbreviation", input: "OAuthAPI", want: "OAuthApi", wantChange: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, changed := SuggestedName(tc.input)
			if got != tc.want {
				t.Fatalf("SuggestedName(%q) = %q, want %q", tc.input, got, tc.want)
			}

			if changed != tc.wantChange {
				t.Fatalf("SuggestedName(%q) changed = %v, want %v", tc.input, changed, tc.wantChange)
			}
		})
	}
}
