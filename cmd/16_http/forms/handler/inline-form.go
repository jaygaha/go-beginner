package handler

import (
	"net/http"
	"text/template"
)

type FormData struct {
	Name string
}

func InlineFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
		<html>
			<body>
				<form action="/inline-submit" method="POST">
					<label>Name:</label>
					<input type="text" name="name">
					<input type="submit" value="Submit">
				</form>
				<br />
				<a href="/form-validation">Validation</a>
			</body>
		</html>
	`

	t := template.Must(template.New("inline-form").Parse(tmpl))
	err := t.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func InlineFormSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get form value
	name := r.FormValue("name")

	// do something with the form value
	data := FormData{Name: name}

	tmpl := `
    <html>
        <body>
            <h1>Hello, {{.Name}}!</h1>
            <a href="/">Go back</a>
        </body>
    </html>
    `
	t := template.Must(template.New("result").Parse(tmpl))
	err := t.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
