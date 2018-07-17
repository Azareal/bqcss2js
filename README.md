# Box Query CSS to JavaScript
A little experiment to transpile my element query motivated oqcss format to JavaScript in Golang.

Bqcss aka Box Query CSS is like Eqcss, but several changes rules have been made to help simplify the format to both improve performance, reduce the complexity, and to reduce the cognitive burden the format imposes on the writer.

This is currently under construction and not functional yet, but feel free to keep an eye on this repo!

```css
/* only run this query when the width of .sidebar is at-least 50 pixels */
@element .sidebar and (min-width: 50px) {
	.row {
		background-color: red;
	}
}
/* this rule applies to .sidebar when the width of it is at-least 50 pixels, try to avoid using this as it's easy to create a cycle */
@element .sidebar and (min-width: 50px) {
	background-color: red;
}
```

As you can see, the format is fairly similar to EQCSS.js, but unlike there, every rule is scoped to $this, unless otherwise specified, this is to reduce the amount of boilerplate needed and to avoid incentivizing patterns which are harder to optimise and reason about.
