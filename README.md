# APIBibleClient
Go Language API Bible Client for the https://scripture.api.bible site

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
#### GET /v1/bibles/{bibleId}/books/{bookId}/chapters
#### GET /v1/bibles/{bibleId}/chapters/{chapterId}
## TODO
### Figure out the licensing and annotate files accordingly (Linux?)
### Keep up on the scripture.api.bible live documentation
### Books (Fetch Books for a Bible)
### Chapters (Fetch Chapters for a Bible)
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
