package awql

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Client struct for AWQL caller
type Client struct {
	Auth
}

// Format type for AWQL format
type Format string

// AWQL formats
const (
	FormatCSVForExcel Format = "CSVFOREXCEL"
	FormatCSV         Format = "CSV"
	FormatTSV         Format = "TSV"
	FormatXML         Format = "XML"
	FormatGzippedCSV  Format = "GZIPPED_CSV"
	FormatGzippedXML  Format = "GZIPPED_XML"
)

// Request struct for awql request
type Request struct {
	Query                  string
	Format                 Format
	SkipReportHeader       bool
	SkipColumnHeader       bool
	SkipReportSummary      bool
	IncludeZeroImpressions bool
	UseRawEnumValues       bool
}

// APIURL the AWQL API's URL
const APIURL = "https://adwords.google.com/api/adwords/reportdownload/v201605"

// NewAWQLClient is a constructor for AWQLClient
func NewAWQLClient(auth *Auth) *Client {
	return &Client{Auth: *auth}
}

// Download downloads a report by awql request
func (c *Client) Download(awqlReq Request) (io.ReadCloser, error) {
	req, err := http.NewRequest(
		"POST",
		APIURL,
		strings.NewReader(url.Values{"__rdquery": {awqlReq.Query}, "__fmt": {string(awqlReq.Format)}}.Encode()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("developerToken", c.DeveloperToken)
	req.Header.Add("clientCustomerId", c.AdwordsID)
	req.Header.Add("skipReportHeader", strconv.FormatBool(awqlReq.SkipReportHeader))
	req.Header.Add("skipColumnHeader", strconv.FormatBool(awqlReq.SkipColumnHeader))
	req.Header.Add("skipReportSummary", strconv.FormatBool(awqlReq.SkipReportSummary))
	req.Header.Add("includeZeroImpressions", strconv.FormatBool(awqlReq.IncludeZeroImpressions))
	req.Header.Add("useRawEnumValues", strconv.FormatBool(awqlReq.UseRawEnumValues))
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var content string
		respBody, errRead := ioutil.ReadAll(resp.Body)
		if errRead == nil {
			content = string(respBody[:])
		}
		return nil, fmt.Errorf("unexpected StatusCode [%d] with Status [%s] Response Body [%s]", resp.StatusCode, resp.Status, content)
	}

	return resp.Body, nil
}
