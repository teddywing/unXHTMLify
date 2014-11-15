unXHTMLify
==========

Removes slashes from self-closing HTML elements.

	<br />

becomes

	<br>

Currently the implementation is very simplistic and liberal. Any `/>` will get replaced to remove the slash.

Takes input from STDIN and outputs to STDOUT.


## License
Licensed under the MIT license. See the included LICENSE file.
