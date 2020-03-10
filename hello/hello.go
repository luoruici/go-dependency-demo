package hello

import "rsc.io/quote/v3"

// Hello returns a greeting.
func Hello() string {
	return quote.HelloV3()
}
