# APIBibleClient

Go Language API Bible Client for the https://scripture.api.bible site

## License for Library

Copyright (c) Ecclesia Foundation

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

## License for Examples and Doctests

MIT License

Copyright 2023 Ecclesia Foundation

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

## Functionality
As of 2023-09-17, the functionality mirrors that which is found in the API documentation found here: https://scripture.api.bible/livedocs In order to use this, you must have an API key which you can get from the folks at https://scripture.api.bible
## TODO
### Prior to releasing v1.0 the following needs to be done:
- Make doctest examples of utils
- Make doctest examples of params
- Review all doctest examples to ensure they look good on pkg.go.dev
### Maintenance
- Keep up on the scripture.api.bible live documentation
- Keep up if they decide to turn the experimental functionality into actual functionality
- better document the params and clean up the examples at bare minimum.
- give examples of each util in action by itself.
- write a set of high-level blurbs in the README to help guide folks to the right stuff in the package documentation.
- set up at least a monthly cadence of running the tests to see if the API has changed, and at least a quarterly review of the API documentation to see if there are any new versions or endpoints.
That will take it from something I'd be ashamed to put a v1.0.0 tag on, to something that mightn't be optimal, but at least deserves the tag.
