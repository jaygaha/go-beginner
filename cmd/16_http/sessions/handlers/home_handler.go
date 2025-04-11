package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
	<html>
		<head>
			<title>Sessions & Cookies</title>
		</head>
		<body>
			<h1>Cookies</h1>
			<ol>
				<li><a href="/cookie/set">Set Cookie</a></li>
				<li><a href="/cookie/get">Get Cookie</a></li>
				<li><a href="/cookie/delete">Delete Cookie</a></li>
			</ol>
			<br />
			<hr>
			<br />
			<h1>Sessions</h1>
			<ol>
				<li><a href="/auth/login">Login</a></li>
				<li><a href="/user/profile">Profile</a></li>
			</ol>
		</body>
	</html>
`

	t := template.Must(template.New("home").Parse(tmpl))
	err := t.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
