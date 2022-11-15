// Code generated by qtc from "basepage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line ../templates/basepage.qtpl:1
package templates

//line ../templates/basepage.qtpl:1
import (
	"github.com/thexyno/xynoblog/statics"
	"strings"
	"time"
)

//line ../templates/basepage.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line ../templates/basepage.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line ../templates/basepage.qtpl:9
type Page interface {
//line ../templates/basepage.qtpl:9
	Title() string
//line ../templates/basepage.qtpl:9
	StreamTitle(qw422016 *qt422016.Writer)
//line ../templates/basepage.qtpl:9
	WriteTitle(qq422016 qtio422016.Writer)
//line ../templates/basepage.qtpl:9
	Body() string
//line ../templates/basepage.qtpl:9
	StreamBody(qw422016 *qt422016.Writer)
//line ../templates/basepage.qtpl:9
	WriteBody(qq422016 qtio422016.Writer)
//line ../templates/basepage.qtpl:9
	Description() string
//line ../templates/basepage.qtpl:9
	StreamDescription(qw422016 *qt422016.Writer)
//line ../templates/basepage.qtpl:9
	WriteDescription(qq422016 qtio422016.Writer)
//line ../templates/basepage.qtpl:9
	Head() string
//line ../templates/basepage.qtpl:9
	StreamHead(qw422016 *qt422016.Writer)
//line ../templates/basepage.qtpl:9
	WriteHead(qq422016 qtio422016.Writer)
//line ../templates/basepage.qtpl:9
}

//line ../templates/basepage.qtpl:19
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line ../templates/basepage.qtpl:19
	qw422016.N().S(`
<!doctype html>
<html>
    <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <!--<meta name="description" content="`)
//line ../templates/basepage.qtpl:25
	p.StreamDescription(qw422016)
//line ../templates/basepage.qtpl:25
	qw422016.N().S(`">-->
      <link href="/css/`)
//line ../templates/basepage.qtpl:26
	qw422016.E().S(strings.TrimSpace(statics.CSSFile))
//line ../templates/basepage.qtpl:26
	qw422016.N().S(`" rel="stylesheet">
      <title>`)
//line ../templates/basepage.qtpl:27
	p.StreamTitle(qw422016)
//line ../templates/basepage.qtpl:27
	qw422016.N().S(`</title>
      `)
//line ../templates/basepage.qtpl:28
	p.StreamHead(qw422016)
//line ../templates/basepage.qtpl:28
	qw422016.N().S(`
    </head>
    <body class="mx-2">
      <header class="top-0 z-40 w-full flex-none max-w-8xl mx-auto py-4 px-8 relative flex items-center">
        <a class="pr-8 mr-3 text-2xl flex-none text-neutral_orange visited:text-neutral_orange hover:text-bright_orange font-semibold overflow-hidden md:w-auto" href="/">xynos space</a>
        <a class="mr-3 text-xl flex-none font-semibold overflow-hidden md:w-auto" href="/posts">Blog</a>
      </header>
      `)
//line ../templates/basepage.qtpl:35
	p.StreamBody(qw422016)
//line ../templates/basepage.qtpl:35
	qw422016.N().S(`
      <footer class="flex flex-col items-center justify-center bottom-0 pt-12 backdrop-blur">
        <p>Copyright (C) `)
//line ../templates/basepage.qtpl:37
	qw422016.N().D(time.Now().Year())
//line ../templates/basepage.qtpl:37
	qw422016.N().S(` xyno (Philipp Hochkamp)</p>
        <br>
        <p>
          <a href="/impressum-de">Impressum / Datenschutzerklärung</a>
        </p>
      </footer>
    </body>
</html>
`)
//line ../templates/basepage.qtpl:45
}

//line ../templates/basepage.qtpl:45
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line ../templates/basepage.qtpl:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../templates/basepage.qtpl:45
	StreamPageTemplate(qw422016, p)
//line ../templates/basepage.qtpl:45
	qt422016.ReleaseWriter(qw422016)
//line ../templates/basepage.qtpl:45
}

//line ../templates/basepage.qtpl:45
func PageTemplate(p Page) string {
//line ../templates/basepage.qtpl:45
	qb422016 := qt422016.AcquireByteBuffer()
//line ../templates/basepage.qtpl:45
	WritePageTemplate(qb422016, p)
//line ../templates/basepage.qtpl:45
	qs422016 := string(qb422016.B)
//line ../templates/basepage.qtpl:45
	qt422016.ReleaseByteBuffer(qb422016)
//line ../templates/basepage.qtpl:45
	return qs422016
//line ../templates/basepage.qtpl:45
}
