//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package params

import (
	"net/url"
)

// The CanProduceQueryParams interface exists
// to help with producing URLs with query parameters.
// <ImplementingParamStruct>.ProduceQueryParameters().Encode()
// yields a string of query parameters for an HTTP request.
type CanProduceQueryParams interface {
	ProduceQueryParameters() url.Values
}
