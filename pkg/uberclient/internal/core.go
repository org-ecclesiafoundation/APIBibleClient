//  Copyright (c) Ecclesia Foundation
//  This Source Code Form is subject to the terms of the Mozilla Public
//  License, v. 2.0. If a copy of the MPL was not distributed with this
//  file, You can obtain one at https://mozilla.org/MPL/2.0/.

package internal

import (
	"www.ecclesiafoundation.org/apibibleclient/pkg/client"
	"www.ecclesiafoundation.org/apibibleclient/pkg/client/params"
)

func GetBibleReference(apiKey string, bibleId string, reference string, params *params.BiblePassageParams) (string, error) {
	refForClient, refParseErr := ParseRef(reference)
	if refParseErr != nil {
		return "", refParseErr
	}
	return client.GetBiblePassage(apiKey, bibleId, refForClient, params)
}
