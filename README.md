# APIBibleClient

Go Language API Bible Client for the https://scripture.api.bible site

> Copyright (c) Ecclesia Foundation
> This Source Code Form is subject to the terms of the Mozilla Public
> License, v. 2.0. If a copy of the MPL was not distributed with this
> file, You can obtain one at https://mozilla.org/MPL/2.0/.

## Functionality
### Bibles (Fetch Bibles and metadata)
#### GET /v1/bibles
#### GET /v1/bibles/{bibleId}
#### GET /v1/audio-bibles
#### GET /v1/audio-bibles/{audioBibleId}
### Books (Fetch Books for a Bible)
#### GET /v1/bibles/{bibleId}/books
#### GET /v1/bibles/{bibleId}/books/{bookId}
#### GET /v1/audio-bibles/{audioBibleId}/books
#### GET /v1/audio-bibles/{audioBibleId}/books/{bookId}
### Chapters (Fetch Chapters for a Bible)
#### GET /v1/bibles/{bibleId}/books/{bookId}/chapters
#### GET /v1/bibles/{bibleId}/chapters/{chapterId}
#### GET /v1/audio-bibles/{audioBibleId}/books/{bookId}/chapters
#### GET /v1/audio-bibles/{audioBibleId}/chapters/{chapterId}
### Sections (Fetch Sections for a Bible)
#### GET /v1/bibles/{bibleId}/books/{bookId}/sections
#### GET /v1/bibles/{bibleId}/chapters/{chapterId}/sections
#### GET /v1/bibles/{bibleId}/sections/{sectionId}
### Passages (Fetch a Passage for a Bible)
#### GET /v1/bibles/{bibleId}/passages/{passageId}
### Verses (Fetch Verses for a Bible)
#### GET /v1/bibles/{bibleId}/chapters/{chapterId}/verses
#### GET /v1/bibles/{bibleId}/verses/{verseId}
### Search (Search by keyword or reference)
#### GET /v1/bibles/{bibleId}/search
## TODO
Prior to releasing v1.0 the following needs to be done:
- Document which API calls are broken on their side (put it in the README and in the comments above the functions that use them)
- Take the stuff we put into the main file, simplify them to make the responses shorter and turn them into test cases to be run at the appropriate (Unit xor API Interface) level.
- Put into the TODO section a list of the things we will need to check up on and act accordingly. At this point, it seems that is constituted by the sections API.
- Keep up on the scripture.api.bible live documentation
