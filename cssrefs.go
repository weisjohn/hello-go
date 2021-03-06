package main

import (
	"fmt"
	"strings"

	"github.com/weisjohn/cssrefs"
)

func main() {
	// example HTML reader
	reader := strings.NewReader(`

        @import url("fineprint.css") print;
        @import "../foo.css";

        @font-face {
          font-family: 'Glyphicons Halflings';
          src: url('../fonts/bootstrap/glyphicons-halflings-regular.eot');
          src: url('../fonts/bootstrap/glyphicons-halflings-regular.eot?#iefix') format('embedded-opentype'), 
            url('../fonts/bootstrap/glyphicons-halflings-regular.woff') format('woff'), 
            url('../fonts/bootstrap/glyphicons-halflings-regular.ttf') format('truetype'), 
            url('../fonts/bootstrap/glyphicons-halflings-regular.svg#glyphicons_halflingsregular') format('svg'); 
        }

        body {
            background: url('../img/light_honeycomb_@2X.png');
            background-image: url('../img/light_honeycomb.png');
            background-size: 270px 289px;
            background-repeat: repeat; 
        }
    `)

	// get the refs from the implementation
	refs := cssrefs.All(reader)

	for _, req := range refs {
		fmt.Println(req.Token, ":", req.URI)
	}
}
