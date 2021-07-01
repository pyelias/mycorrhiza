// Code generated by qtc from "admin.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/admin.qtpl:1
package views

//line views/admin.qtpl:1
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/admin.qtpl:2
import "github.com/bouncepaw/mycorrhiza/user"

//line views/admin.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/admin.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/admin.qtpl:4
func StreamAdminPanelHTML(qw422016 *qt422016.Writer) {
//line views/admin.qtpl:4
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<h1>Administrative functions</h1>
	<section>
		<h2>Safe things</h2>
		<ul>
			<li><a href="/about">About this wiki</a></li>
			<li><a href="/user-list">User list</a></li>
			<li><a href="/update-header-links">Update header links</a></li>
			<li><a href="/admin/users/">Manage users</a></li>
		</ul>
	</section>
	<section>
		<h2>Dangerous things</h2>
		<form action="/admin/shutdown" method="POST" style="float:left">
			<fieldset>
				<legend>Shutdown wiki</legend>
				<input type="submit" class="btn">
			</fieldset>
		</form>
		<form action="/reindex" method="GET" style="float:left">
			<fieldset>
				<legend>Reindex hyphae</legend>
				<input type="submit" class="btn">
			</fieldset>
		</form>
	</section>
</main>
</div>
`)
//line views/admin.qtpl:34
}

//line views/admin.qtpl:34
func WriteAdminPanelHTML(qq422016 qtio422016.Writer) {
//line views/admin.qtpl:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/admin.qtpl:34
	StreamAdminPanelHTML(qw422016)
//line views/admin.qtpl:34
	qt422016.ReleaseWriter(qw422016)
//line views/admin.qtpl:34
}

//line views/admin.qtpl:34
func AdminPanelHTML() string {
//line views/admin.qtpl:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/admin.qtpl:34
	WriteAdminPanelHTML(qb422016)
//line views/admin.qtpl:34
	qs422016 := string(qb422016.B)
//line views/admin.qtpl:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/admin.qtpl:34
	return qs422016
//line views/admin.qtpl:34
}

//line views/admin.qtpl:36
func StreamAdminUsersPanelHTML(qw422016 *qt422016.Writer, userList []*user.User) {
//line views/admin.qtpl:36
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<h1>Manage users</h1>

	<form action="/admin/reindex-users" method="post">
		<button class="btn" type="submit">Reindex users</button>
	</form>

	<h2>Users list</h2>

	<table class="users-table">
	<thead>
		<tr>
			<th>Name</th>
			<th>Group</th>
			<th>Registered at</th>
			<th aria-label="Actions"></th>
		</tr>
	</thead>
	<tbody>
	`)
//line views/admin.qtpl:57
	for _, u := range userList {
//line views/admin.qtpl:57
		qw422016.N().S(`
		<tr>
			<td class="table-cell--fill">
				<a href="/hypha/`)
//line views/admin.qtpl:60
		qw422016.N().U(cfg.UserHypha)
//line views/admin.qtpl:60
		qw422016.N().S(`/`)
//line views/admin.qtpl:60
		qw422016.N().U(u.Name)
//line views/admin.qtpl:60
		qw422016.N().S(`">`)
//line views/admin.qtpl:60
		qw422016.E().S(u.Name)
//line views/admin.qtpl:60
		qw422016.N().S(`</a>
			</td>
			<td>`)
//line views/admin.qtpl:62
		qw422016.E().S(u.Group)
//line views/admin.qtpl:62
		qw422016.N().S(`</td>
			<td>
				`)
//line views/admin.qtpl:64
		if u.RegisteredAt.IsZero() {
//line views/admin.qtpl:64
			qw422016.N().S(`
					unknown
				`)
//line views/admin.qtpl:66
		} else {
//line views/admin.qtpl:66
			qw422016.N().S(`
					`)
//line views/admin.qtpl:67
			qw422016.E().S(u.RegisteredAt.UTC().Format("2006-01-02 15:04"))
//line views/admin.qtpl:67
			qw422016.N().S(`
				`)
//line views/admin.qtpl:68
		}
//line views/admin.qtpl:68
		qw422016.N().S(`
			</td>
			<td>
				<a href="/admin/users/`)
//line views/admin.qtpl:71
		qw422016.N().U(u.Name)
//line views/admin.qtpl:71
		qw422016.N().S(`/edit">Edit</a>
			</td>
		</tr>
	`)
//line views/admin.qtpl:74
	}
//line views/admin.qtpl:74
	qw422016.N().S(`
	</tbody>
	</table>
</main>
</div>
`)
//line views/admin.qtpl:79
}

//line views/admin.qtpl:79
func WriteAdminUsersPanelHTML(qq422016 qtio422016.Writer, userList []*user.User) {
//line views/admin.qtpl:79
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/admin.qtpl:79
	StreamAdminUsersPanelHTML(qw422016, userList)
//line views/admin.qtpl:79
	qt422016.ReleaseWriter(qw422016)
//line views/admin.qtpl:79
}

//line views/admin.qtpl:79
func AdminUsersPanelHTML(userList []*user.User) string {
//line views/admin.qtpl:79
	qb422016 := qt422016.AcquireByteBuffer()
//line views/admin.qtpl:79
	WriteAdminUsersPanelHTML(qb422016, userList)
//line views/admin.qtpl:79
	qs422016 := string(qb422016.B)
//line views/admin.qtpl:79
	qt422016.ReleaseByteBuffer(qb422016)
//line views/admin.qtpl:79
	return qs422016
//line views/admin.qtpl:79
}

//line views/admin.qtpl:81
func StreamAdminUsersUserHTML(qw422016 *qt422016.Writer, u *user.User) {
//line views/admin.qtpl:81
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<h1>`)
//line views/admin.qtpl:84
	qw422016.E().S(u.Name)
//line views/admin.qtpl:84
	qw422016.N().S(`</h1>

	<form action="" method="post">
		<div class="form-field">
			<label for="group">Group:</label>
			<select id="group" name="group">
				<option`)
//line views/admin.qtpl:90
	if u.Group == "anon" {
//line views/admin.qtpl:90
		qw422016.N().S(` selected`)
//line views/admin.qtpl:90
	}
//line views/admin.qtpl:90
	qw422016.N().S(`>anon</option>
				<option`)
//line views/admin.qtpl:91
	if u.Group == "editor" {
//line views/admin.qtpl:91
		qw422016.N().S(` selected`)
//line views/admin.qtpl:91
	}
//line views/admin.qtpl:91
	qw422016.N().S(`>editor</option>
				<option`)
//line views/admin.qtpl:92
	if u.Group == "trusted" {
//line views/admin.qtpl:92
		qw422016.N().S(` selected`)
//line views/admin.qtpl:92
	}
//line views/admin.qtpl:92
	qw422016.N().S(`>trusted</option>
				<option`)
//line views/admin.qtpl:93
	if u.Group == "moderator" {
//line views/admin.qtpl:93
		qw422016.N().S(` selected`)
//line views/admin.qtpl:93
	}
//line views/admin.qtpl:93
	qw422016.N().S(`>moderator</option>
				<option`)
//line views/admin.qtpl:94
	if u.Group == "admin" {
//line views/admin.qtpl:94
		qw422016.N().S(` selected`)
//line views/admin.qtpl:94
	}
//line views/admin.qtpl:94
	qw422016.N().S(`>admin</option>
			</select>
		</div>

		<button class="btn" type="submit">Update</button>
		<a class="btn btn_weak" href="/admin/users/">Cancel</a>
	</form>
</main>
</div>
`)
//line views/admin.qtpl:103
}

//line views/admin.qtpl:103
func WriteAdminUsersUserHTML(qq422016 qtio422016.Writer, u *user.User) {
//line views/admin.qtpl:103
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/admin.qtpl:103
	StreamAdminUsersUserHTML(qw422016, u)
//line views/admin.qtpl:103
	qt422016.ReleaseWriter(qw422016)
//line views/admin.qtpl:103
}

//line views/admin.qtpl:103
func AdminUsersUserHTML(u *user.User) string {
//line views/admin.qtpl:103
	qb422016 := qt422016.AcquireByteBuffer()
//line views/admin.qtpl:103
	WriteAdminUsersUserHTML(qb422016, u)
//line views/admin.qtpl:103
	qs422016 := string(qb422016.B)
//line views/admin.qtpl:103
	qt422016.ReleaseByteBuffer(qb422016)
//line views/admin.qtpl:103
	return qs422016
//line views/admin.qtpl:103
}
