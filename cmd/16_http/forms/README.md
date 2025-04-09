# Go Form Handling and Validation Guide

This guide demonstrates different approaches to handle HTML forms and implement form validation in Go web applications.

## Basic Form Handling

### Form Example
The simplest way to handle forms in Go using inline templates:

```go
func InlineFormHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := `
        <html>
            <body>
                <form action="/inline-submit" method="POST">
                    <label>Name:</label>
                    <input type="text" name="name">
                    <input type="submit" value="Submit">
                </form>
            </body>
        </html>
    `
    t := template.Must(template.New("inline-form").Parse(tmpl))
    t.Execute(w, nil)
}
```

#### Getting Form Values

Retrieve form values using r.FormValue() :
```go
name := r.FormValue("username")
email := r.FormValue("email")
```

#### Form Validation

1. Define a struct to hold form data and validation errors:
```go
type ValidationFormData struct {
    Username string
    Password string
    Error    string
}
```
2. Implement validation logic:
```go
if strings.TrimSpace(data.Username) == "" {
    data.Error = "Username is required"
    return
}
if len(data.Username) < 5 {
    data.Error = "Username must be at least 5 characters long"
    return
}
```
3. Advanced Validation with Regular Expressions
```go
// Email validation
var rxEmail = regexp.MustCompile(".+@.+\\..+")
match := rxEmail.MatchString(email)

// Password validation (minimum 8 characters with at least one letter and number)
var rxPassword = regexp.MustCompile(`^[A-Za-z\d]{8,}$`)
```

## Best Practices

1. Separate Concerns
   
   - Keep form handling logic separate from validation logic
   - Use structs to organize form data and validation errors
2. Input Sanitization
   
   - Always trim whitespace from inputs
   - Validate input length and format
   - Use regular expressions for complex validations
3. Error Handling
   
   - Return clear error messages to users
   - Use maps for multiple validation errors
   - Handle both client-side and server-side validation
4. Security
   
   - Always validate on the server side
   - Use HTTPS for form submissions
   - Implement CSRF protection
   - Sanitize inputs before using in database queries
5. User Experience
   
   - Preserve valid form data when validation fails
   - Show clear error messages
   - Use appropriate HTTP status codes
   - Implement proper redirects after successful submission