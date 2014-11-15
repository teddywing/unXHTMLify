unXHTMLify
==========

Removes slashes from self-closing HTML elements.

	<br />

becomes

	<br>

Currently the implementation is very simplistic and liberal. Any `/>` will get replaced to remove the slash.

Takes input from STDIN and outputs to STDOUT.


## Installing

	$ go get github.com/teddywing/unXHTMLify
	$ cd $GOPATH/src/github.com/teddywing/unXHTMLify
	$ go install


## Running
Pipe an HTML file into `unXHTMLify` to convert the file:

	$ cat sample.html | unXHTMLify > sample-converted.html


## License
Licensed under the MIT license. See the included LICENSE file.
