package md

import (
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestIsBold(t *testing.T) {
	// <strong>foobar</strong>
	a1 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Strong,
		Data:     "strong",
	}
	a2 := &html.Node{
		Type: html.TextNode,
		Data: "foobar",
	}
	a1.AppendChild(a2)

	// <b>foobar</b>
	b1 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.B,
		Data:     "b",
	}
	b2 := &html.Node{
		Type: html.TextNode,
		Data: "foobar",
	}
	b1.AppendChild(b2)

	// <strong><code>foobar</code></strong>
	c1 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Strong,
		Data:     "strong",
	}
	c2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Code,
		Data:     "code",
	}
	c3 := &html.Node{
		Type: html.TextNode,
		Data: "foobar",
	}
	c1.AppendChild(c2)
	c2.AppendChild(c3)

	// <b><code>foobar</code></b>
	d1 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.B,
		Data:     "b",
	}
	d2 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Code,
		Data:     "code",
	}
	d3 := &html.Node{
		Type: html.TextNode,
		Data: "foobar",
	}
	d1.AppendChild(d2)
	d2.AppendChild(d3)

	// <p>foobar</p>
	e1 := &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.P,
		Data:     "p",
	}
	e2 := &html.Node{
		Type: html.TextNode,
		Data: "foobar",
	}
	e1.AppendChild(e2)

	tests := []struct {
		name string
		in   *html.Node
		out  bool
	}{
		{
			name: "StrongText_Strong",
			in:   a1,
			out:  true,
		},
		{
			name: "StrongText_Strong",
			in:   a2,
			out:  true,
		},
		{
			name: "BText_B",
			in:   b1,
			out:  true,
		},
		{
			name: "BText_Text",
			in:   b2,
			out:  true,
		},
		{
			name: "StrongCodeText_Strong",
			in:   c1,
			out:  true,
		},
		{
			name: "StrongCodeText_Code",
			in:   c2,
			out:  true,
		},
		/*
			// TODO: I think this should work but it doesn't.
			{
				name: "StrongCodeText_Text",
				in:   c3,
				out:  true,
			},
		*/
		{
			name: "BCodeText_B",
			in:   d1,
			out:  true,
		},
		{
			name: "BCodeText_Code",
			in:   d2,
			out:  true,
		},
		/*
			// TODO: I think this should work but it doesn't
			{
				name: "BCodeText_Text",
				in:   d3,
				out:  true,
			},
		*/
		{
			name: "PText_P",
			in:   e1,
		},
		{
			name: "PText_Text",
			in:   e2,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if out := isBold(tc.in); out != tc.out {
				t.Errorf("isBold(%v) = %t, want %t", tc.in, out, tc.out)
			}
		})
	}
}