package alyaze

import (
	"net/http"
)

// Job may consist only of a URL, in which case webanalyse will
// proceed to download from that URL, or it may consist of the
// Body and Headers of a request to a URL and the URL itself,
// in which case these fields will be trusted and used for
// analysis without further network traffic.
// If a Job is constructed using the OfflineJob constructor
// then a flag will be set to prevent downloading regardless
// of the contents (or absence) of the Body or Headers fields.
type Job struct {
	URL              string
	Body             []byte
	Headers          http.Header //map[string][]string
	Cookies          []*http.Cookie
	Crawl            int
	SearchSubdomain  bool
	forceNotDownload bool
	followRedirect   bool
}

// NewOfflineJob constructs a job out of the constituents of a
// webanalyzer analysis; a URL, a body, and response headers.
// This constructor also sets a flag to explicitly prevent
// fetching from the URL even if the body and headers are nil
// or empty. Use this for...offline jobs.
func NewOfflineJob(url, body string, headers map[string][]string) *Job {
	return &Job{
		URL:              url,
		Body:             []byte(body),
		Headers:          headers,
		Crawl:            0,
		SearchSubdomain:  false,
		forceNotDownload: true,
		followRedirect:   false,
	}
}

// NewOnlineJob constructs a job that may either have a URL only,
// or a URL, Body and Headers. If it contains at least a URL and Body,
// then webanalyzer will not re-download the data, but if a Body is
// absent then downloading will be attempted.
func NewOnlineJob(url, body string, headers map[string][]string, crawlCount int, searchSubdomain bool, redirect bool) *Job {
	return &Job{
		URL:              url,
		Body:             []byte(body),
		Headers:          headers,
		Crawl:            crawlCount,
		SearchSubdomain:  searchSubdomain,
		forceNotDownload: false,
		followRedirect:   redirect,
	}
}
