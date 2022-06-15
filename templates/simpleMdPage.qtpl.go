// Code generated by qtc from "simpleMdPage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line ../templates/simpleMdPage.qtpl:1
package templates

//line ../templates/simpleMdPage.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line ../templates/simpleMdPage.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line ../templates/simpleMdPage.qtpl:2
type SimpleMdPage struct {
	PageTitle       []byte
	RenderedContent []byte
}

//line ../templates/simpleMdPage.qtpl:8
func (p *SimpleMdPage) StreamTitle(qw422016 *qt422016.Writer) {
//line ../templates/simpleMdPage.qtpl:8
	qw422016.E().Z(p.PageTitle)
//line ../templates/simpleMdPage.qtpl:8
	qw422016.N().S(` - xynos space`)
//line ../templates/simpleMdPage.qtpl:8
}

//line ../templates/simpleMdPage.qtpl:8
func (p *SimpleMdPage) WriteTitle(qq422016 qtio422016.Writer) {
//line ../templates/simpleMdPage.qtpl:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../templates/simpleMdPage.qtpl:8
	p.StreamTitle(qw422016)
//line ../templates/simpleMdPage.qtpl:8
	qt422016.ReleaseWriter(qw422016)
//line ../templates/simpleMdPage.qtpl:8
}

//line ../templates/simpleMdPage.qtpl:8
func (p *SimpleMdPage) Title() string {
//line ../templates/simpleMdPage.qtpl:8
	qb422016 := qt422016.AcquireByteBuffer()
//line ../templates/simpleMdPage.qtpl:8
	p.WriteTitle(qb422016)
//line ../templates/simpleMdPage.qtpl:8
	qs422016 := string(qb422016.B)
//line ../templates/simpleMdPage.qtpl:8
	qt422016.ReleaseByteBuffer(qb422016)
//line ../templates/simpleMdPage.qtpl:8
	return qs422016
//line ../templates/simpleMdPage.qtpl:8
}

//line ../templates/simpleMdPage.qtpl:9
func (p *SimpleMdPage) StreamDescription(qw422016 *qt422016.Writer) {
//line ../templates/simpleMdPage.qtpl:9
}

//line ../templates/simpleMdPage.qtpl:9
func (p *SimpleMdPage) WriteDescription(qq422016 qtio422016.Writer) {
//line ../templates/simpleMdPage.qtpl:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../templates/simpleMdPage.qtpl:9
	p.StreamDescription(qw422016)
//line ../templates/simpleMdPage.qtpl:9
	qt422016.ReleaseWriter(qw422016)
//line ../templates/simpleMdPage.qtpl:9
}

//line ../templates/simpleMdPage.qtpl:9
func (p *SimpleMdPage) Description() string {
//line ../templates/simpleMdPage.qtpl:9
	qb422016 := qt422016.AcquireByteBuffer()
//line ../templates/simpleMdPage.qtpl:9
	p.WriteDescription(qb422016)
//line ../templates/simpleMdPage.qtpl:9
	qs422016 := string(qb422016.B)
//line ../templates/simpleMdPage.qtpl:9
	qt422016.ReleaseByteBuffer(qb422016)
//line ../templates/simpleMdPage.qtpl:9
	return qs422016
//line ../templates/simpleMdPage.qtpl:9
}

//line ../templates/simpleMdPage.qtpl:10
func (p *SimpleMdPage) StreamHead(qw422016 *qt422016.Writer) {
//line ../templates/simpleMdPage.qtpl:10
}

//line ../templates/simpleMdPage.qtpl:10
func (p *SimpleMdPage) WriteHead(qq422016 qtio422016.Writer) {
//line ../templates/simpleMdPage.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../templates/simpleMdPage.qtpl:10
	p.StreamHead(qw422016)
//line ../templates/simpleMdPage.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line ../templates/simpleMdPage.qtpl:10
}

//line ../templates/simpleMdPage.qtpl:10
func (p *SimpleMdPage) Head() string {
//line ../templates/simpleMdPage.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line ../templates/simpleMdPage.qtpl:10
	p.WriteHead(qb422016)
//line ../templates/simpleMdPage.qtpl:10
	qs422016 := string(qb422016.B)
//line ../templates/simpleMdPage.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line ../templates/simpleMdPage.qtpl:10
	return qs422016
//line ../templates/simpleMdPage.qtpl:10
}

//line ../templates/simpleMdPage.qtpl:13
func (p *SimpleMdPage) StreamBody(qw422016 *qt422016.Writer) {
//line ../templates/simpleMdPage.qtpl:13
	qw422016.N().S(`
	<article class="mx-auto container max-w-7xl">
      <div class="pt-8 py-8">
      	<h1 class="p-0">`)
//line ../templates/simpleMdPage.qtpl:16
	qw422016.E().Z(p.PageTitle)
//line ../templates/simpleMdPage.qtpl:16
	qw422016.N().S(`</h1>
      </div>
	  `)
//line ../templates/simpleMdPage.qtpl:18
	qw422016.N().Z(p.RenderedContent)
//line ../templates/simpleMdPage.qtpl:18
	qw422016.N().S(`
	</article>
`)
//line ../templates/simpleMdPage.qtpl:20
}

//line ../templates/simpleMdPage.qtpl:20
func (p *SimpleMdPage) WriteBody(qq422016 qtio422016.Writer) {
//line ../templates/simpleMdPage.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line ../templates/simpleMdPage.qtpl:20
	p.StreamBody(qw422016)
//line ../templates/simpleMdPage.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line ../templates/simpleMdPage.qtpl:20
}

//line ../templates/simpleMdPage.qtpl:20
func (p *SimpleMdPage) Body() string {
//line ../templates/simpleMdPage.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line ../templates/simpleMdPage.qtpl:20
	p.WriteBody(qb422016)
//line ../templates/simpleMdPage.qtpl:20
	qs422016 := string(qb422016.B)
//line ../templates/simpleMdPage.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line ../templates/simpleMdPage.qtpl:20
	return qs422016
//line ../templates/simpleMdPage.qtpl:20
}