package stripmd

import (
	"fmt"
	"testing"
)

func TestStripMarkdown(t *testing.T) {
	// Same tests as github.com/stiang/remove-markdown
	in := `## This is a heading ##

This is an _emphasized paragraph_ with [a link](http://www.disney.com/). Here's an _[emphasized link](https://write.as)_.

### This is another heading

In ` + "`Getting Started` we **set up** `something`" + ` __foo__.

* Some list
* With items
  * Even indented`

	out := `This is a heading 

This is an emphasized paragraph with a link. Here's an emphasized link.

This is another heading

In Getting Started we set up something foo.

Some list
With items
  Even indented`

	if res := Strip(in); res != out {
		t.Errorf("Original:\n\n%s\n\nGot:\n\n%s", in, res)
	}

	// More extensive tests
	in = `# Markdown is simple
It lets you _italicize words_ and **bold them**, too. You can ~~cross things out~~ and add __emphasis__ to your writing.

But you can also link to stuff, like [Write.as](https://write.as)!

## Organize text with headers
Create sections in your text just like this.

### Use lists
You might already write lists like this:

* Get groceries
* Go for a walk

And sometimes you need to do things in a certain order:

1. Put on clothes
2. Put on shoes
3. Go for a walk

### Highlight text
You can quote interesting people:

> Live long and prosper.

You can even share ` + "`code stuff`."

	out = `Markdown is simple
It lets you italicize words and bold them, too. You can cross things out and add emphasis to your writing.

But you can also link to stuff, like Write.as!

Organize text with headers
Create sections in your text just like this.

Use lists
You might already write lists like this:

Get groceries
Go for a walk

And sometimes you need to do things in a certain order:

Put on clothes
Put on shoes
Go for a walk

Highlight text
You can quote interesting people:

  Live long and prosper.

You can even share code stuff.`

	if res := Strip(in); res != out {
		t.Errorf("Original:\n\n%s\n\nGot:\n\n%s", in, res)
	}

	in = "![] (https://write.as/favicon.ico)"
	out = ""
	if res := Strip(in); res != out {
		t.Errorf("Original:\n\n%s\n\nGot:\n\n%s", in, res)
	}

	in = "![Some image] (https://write.as/favicon.ico)"
	out = "Some image"
	if res := Strip(in); res != out {
		t.Errorf("Original:\n\n%s\n\nGot:\n\n%s", in, res)
	}

	in = "![Some image](https://write.as/favicon.ico)"
	out = "Some image"
	if res := Strip(in); res != out {
		t.Errorf("Original:\n\n%s\n\nGot:\n\n%s", in, res)
	}
	// Test to skip images
	in = "![Some image](https://write.as/favicon.ico)"
	out = ""
	if res := StripOptions(in, Options{SkipImages: true}); res != out {
		t.Errorf("Original:\n\n%s\n\nGot:\n\n%s", in, res)
	}

	// Test for skipping single underscores
	in = "_this is single emphasis_"
	out = "_this is single emphasis_"
	if res := StripOptions(in, Options{SkipUnderscores: true}); res != out {
		t.Errorf("Original:\n\n%s\n\nGot:\n\n%s", in, res)
	}
	if res := Strip(in); res == out {
		t.Errorf("Original:\n\n%s\n\nGot:\n\n%s", in, res)
	}

	// Test for skipping double underscores
	in = "__this is double emphasis__"
	out = "__this is double emphasis__"
	if res := StripOptions(in, Options{SkipUnderscores: true}); res != out {
		t.Errorf("Original:\n\n%s\n\nGot:\n\n%s", in, res)
	}
	if res := Strip(in); res == out {
		t.Errorf("Original:\n\n%s\n\nGot:\n\n%s", in, res)
	}

}

func ExampleStrip() {
	fmt.Println(Strip(`# Hello, world!

This is [a Go library](https://github.com/writeas/go-strip-markdown) for stripping **Markdown** from _any_ text.`))

	// Output:
	// Hello, world!
	//
	// This is a Go library for stripping Markdown from any text.
}
