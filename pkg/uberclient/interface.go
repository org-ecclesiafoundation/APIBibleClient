//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package uberclient

import (
	"www.ecclesiafoundation.org/apibibleclient/pkg/params"
	"www.ecclesiafoundation.org/apibibleclient/pkg/uberclient/internal"
)

func GetBibleReference(apiKey string, bibleId string, reference string, params *params.BiblePassageParams) (string, error) {
	return internal.GetBibleReference(apiKey, bibleId, reference, params)
}
