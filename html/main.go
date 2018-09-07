package main

import (
	"fmt"
	"html"
)

func main() {
	const s1 = `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Println(html.EscapeString(s1))

	const s2 = `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
	fmt.Println(html.UnescapeString(s2))
}
