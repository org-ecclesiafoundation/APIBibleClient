# APIBibleClient

## In Brief
This is a Go Language API Bible Client for the https://scripture.api.bible site

Please refer to the below link for the Go Package Documentation

[![Go Reference](https://pkg.go.dev/badge/www.ecclesiafoundation.org/apibibleclient.svg)](https://pkg.go.dev/www.ecclesiafoundation.org/apibibleclient)

Link to [Licenses](#Licenses)

## Functionality
As of 2023-09-17, the functionality mirrors that which is found in the API documentation found here: https://scripture.api.bible/livedocs In order to use this, you must have an API key which you can get from the folks at https://scripture.api.bible

## Testing the Library
All tests can be run with the `make test` command. Test suites can also be individually run. Please consult the Makefile for more details.

## TODO

### Prior to releasing v1.0:
- Review all doctest examples and documentation to ensure they make sense and look good on pkg.go.dev
- Write a set of high-level blurbs in the README to help guide folks to the right stuff in the package documentation.

### Prior to releasing v1.1:
- Add bible verse lexer/parser
- Add uberclient.GetBibleReference(apiKey, bibleId, reference ...)

### Maintenance
- Keep up on the scripture.api.bible live documentation
- Keep up if they decide to turn the experimental functionality into actual functionality by periodically testing the library and looking for failures.
- Set up at least a monthly cadence of running the tests to see if the API has changed, and at least a quarterly review of the API documentation to see if there are any new versions or endpoints.

## Licenses

### Examples and Doctests
MIT License

Copyright 2023 Ecclesia Foundation

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

### Library
Copyright (c) Ecclesia Foundation

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

ref. LICENSE.txt
