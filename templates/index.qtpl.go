// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/index.qtpl:1
package templates

//line templates/index.qtpl:1
import (
	"github.com/thexyno/xynoblog/db"
)

//line templates/index.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/index.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/index.qtpl:7
type IndexPage struct {
	Posts []db.Post
}

//line templates/index.qtpl:13
func (p *IndexPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/index.qtpl:13
	qw422016.N().S(`
	This is the main page
`)
//line templates/index.qtpl:15
}

//line templates/index.qtpl:15
func (p *IndexPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/index.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/index.qtpl:15
	p.StreamTitle(qw422016)
//line templates/index.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line templates/index.qtpl:15
}

//line templates/index.qtpl:15
func (p *IndexPage) Title() string {
//line templates/index.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/index.qtpl:15
	p.WriteTitle(qb422016)
//line templates/index.qtpl:15
	qs422016 := string(qb422016.B)
//line templates/index.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/index.qtpl:15
	return qs422016
//line templates/index.qtpl:15
}

//line templates/index.qtpl:18
func (p *IndexPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/index.qtpl:18
	qw422016.N().S(`
	<div class="mx-auto container">
	<h1 class="py-8 text-2xl font-semibold">Posts:</h1>
      <div class="px-8 flex flex-col">
      `)
//line templates/index.qtpl:22
	if len(p.Posts) == 0 {
//line templates/index.qtpl:22
		qw422016.N().S(`
	  	No posts.
	  `)
//line templates/index.qtpl:24
	} else {
//line templates/index.qtpl:24
		qw422016.N().S(`
	  		`)
//line templates/index.qtpl:25
		streamemitPosts(qw422016, p.Posts)
//line templates/index.qtpl:25
		qw422016.N().S(`
	  `)
//line templates/index.qtpl:26
	}
//line templates/index.qtpl:26
	qw422016.N().S(`
	  </div>
	</div>
`)
//line templates/index.qtpl:29
}

//line templates/index.qtpl:29
func (p *IndexPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/index.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/index.qtpl:29
	p.StreamBody(qw422016)
//line templates/index.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line templates/index.qtpl:29
}

//line templates/index.qtpl:29
func (p *IndexPage) Body() string {
//line templates/index.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/index.qtpl:29
	p.WriteBody(qb422016)
//line templates/index.qtpl:29
	qs422016 := string(qb422016.B)
//line templates/index.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/index.qtpl:29
	return qs422016
//line templates/index.qtpl:29
}

//line templates/index.qtpl:31
func streamemitPosts(qw422016 *qt422016.Writer, posts []db.Post) {
//line templates/index.qtpl:31
	qw422016.N().S(`
   `)
//line templates/index.qtpl:32
	for _, v := range posts {
//line templates/index.qtpl:32
		qw422016.N().S(`
     <a href="/post/`)
//line templates/index.qtpl:33
		qw422016.E().S(v.Id)
//line templates/index.qtpl:33
		qw422016.N().S(`">
      <span class="font-semibold text-lg">`)
//line templates/index.qtpl:34
		qw422016.E().S(v.Title)
//line templates/index.qtpl:34
		qw422016.N().S(`</span>
      <span class="text-xs font-thin text-dark3 dark:text-light3 ">(`)
//line templates/index.qtpl:35
		qw422016.E().S(v.Created.Format("2006-01-02"))
//line templates/index.qtpl:35
		qw422016.N().S(`)</span>
     </a>
   `)
//line templates/index.qtpl:37
	}
//line templates/index.qtpl:37
	qw422016.N().S(`
`)
//line templates/index.qtpl:38
}

//line templates/index.qtpl:38
func writeemitPosts(qq422016 qtio422016.Writer, posts []db.Post) {
//line templates/index.qtpl:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/index.qtpl:38
	streamemitPosts(qw422016, posts)
//line templates/index.qtpl:38
	qt422016.ReleaseWriter(qw422016)
//line templates/index.qtpl:38
}

//line templates/index.qtpl:38
func emitPosts(posts []db.Post) string {
//line templates/index.qtpl:38
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/index.qtpl:38
	writeemitPosts(qb422016, posts)
//line templates/index.qtpl:38
	qs422016 := string(qb422016.B)
//line templates/index.qtpl:38
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/index.qtpl:38
	return qs422016
//line templates/index.qtpl:38
}
