# APIBibleClient

Go Language API Bible Client for the https://scripture.api.bible site

## License for Library

Copyright (c) Ecclesia Foundation

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

## License for Examples

MIT License

Copyright 2023 Ecclesia Foundation

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
### Document Pain Points
- Document which API calls are broken on their side (put it in the README and in the comments above the functions that use them)
### Test & Document Functionality
- Document in the README code samples for each of the functionality and the responses one might expect.
- Take the stuff we put into the main file, simplify them to make the responses shorter and turn them into test cases to be run at the appropriate (Unit xor API Interface) level.
NOTE: The aforementioned two items might wind up executing exactly the same library code. Might this be preferable?
### Maintenance
- Put into the TODO section a list of the things we will need to check up on and act accordingly. At this point, it seems that is constituted by the sections API.
- Keep up on the scripture.api.bible live documentation
