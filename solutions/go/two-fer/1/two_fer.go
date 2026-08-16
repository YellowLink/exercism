// Package twofer determines your dialogue when giving 
// a cookie to another person standing in line at the
// bakery.
package twofer

import ("fmt")

// ShareWith returns a string representing your dialogue when sharing your cookie.
// It takes a string argument representing the person's name.
// If the name is unknown the string will be empty.
func ShareWith(name string) string {
	msg := ""
	if name == "" {
		msg += "One for you, one for me."
	} else {
		msg += fmt.Sprintf("One for %s, one for me.", name)
	}
	return msg
}
