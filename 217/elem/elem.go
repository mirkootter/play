package elem

import (
	"github.com/shurcooL/play/217/vec"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

//func Text(s string) *vec.HTML {
//	return &vec.HTML{
//		Type: html.TextNode,
//		Data: s,
//	}
//}

func Abbr(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.Abbr,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}

func Code(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.Code,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}

func Img(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.Img,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}

func Div(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.Div,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}

func Span(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.Span,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}

func P(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.P,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}

func A(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.A,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}

func H1(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.H1,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}

func H2(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.H2,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}

func Ul(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.Ul,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}

func Li(markup ...vec.MarkupOrComponentOrHTML) *vec.HTML {
	h := &vec.HTML{
		Type:     html.ElementNode,
		DataAtom: atom.Li,
	}
	for _, m := range markup {
		vec.Apply(h, m)
	}
	return h
}
