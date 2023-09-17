package params

import (
	"net/url"
	"strconv"
)

// The BibleSearchParams struct contains all optional parameters
// for the API call to the `bibles/{bibleId}/search` end point
type BibleSearchParams struct {
	Query     string
	Limit     int
	Offset    int
	Sort      string
	Range     string
	Fuzziness string
}

func (params *BibleSearchParams) ProduceQueryParameters() url.Values {
	values := url.Values{}

	if params.Query != "" {
		values.Add("query", params.Query)
	}
	if params.Limit != 0 {
		values.Add("limit", strconv.Itoa(params.Limit))
	}
	if params.Offset != 0 {
		values.Add("offset", strconv.Itoa(params.Offset))
	}
	if params.Sort != "" {
		values.Add("sort", params.Sort)
	}
	if params.Range != "" {
		values.Add("range", params.Range)
	}
	if params.Fuzziness != "" {
		values.Add("fuzziness", params.Fuzziness)
	}

	return values
}
