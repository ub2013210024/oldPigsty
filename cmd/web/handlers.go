// Test

package main

import "net/http"

// Create our handler functions
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my webpage"))
}
