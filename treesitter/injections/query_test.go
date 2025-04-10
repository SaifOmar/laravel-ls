package injections_test

import (
	"testing"

	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/injections"
	"github.com/laravel-ls/laravel-ls/treesitter/language"

	"github.com/stretchr/testify/assert"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func TestQuery(t *testing.T) {
	src := []byte(`<div>
<a href="target">link</a>
<?php echo "hello world" ?>
</div>`)

	lang := language.Get(language.PHP)

	parser := ts.NewParser()
	assert.NoError(t, parser.SetLanguage(lang.TSObject()))
	tree := parser.Parse(src, nil)

	injectionQuery, err := treesitter.ReadQueryFromFile(lang, "injections")
	assert.Nil(t, err)

	query, err := lang.Query(injectionQuery)
	assert.Nil(t, err)

	expected := []injections.Capture{
		{
			Language: "html",
			Combined: true,
			Range: ts.Range{
				StartByte: 0,
				EndByte:   32,
				StartPoint: ts.Point{
					Row:    0,
					Column: 0,
				},
				EndPoint: ts.Point{
					Row:    2,
					Column: 0,
				},
			},
		},
		{
			Language: "html",
			Combined: true,
			Range: ts.Range{
				StartByte: 60,
				EndByte:   66,
				StartPoint: ts.Point{
					Row:    3,
					Column: 0,
				},
				EndPoint: ts.Point{
					Row:    3,
					Column: 6,
				},
			},
		},
	}

	actual := injections.Query(query, tree.RootNode(), src)
	assert.Equal(t, expected, actual)
}

func TestQuery_Blade(t *testing.T) {
	src := []byte(`<div>
	<a href="target">{{ $link }}</a>
	@include('some.file')
	</div>`)

	lang := language.Get(language.Blade)

	parser := ts.NewParser()
	assert.NoError(t, parser.SetLanguage(lang.TSObject()))
	tree := parser.Parse(src, nil)

	injectionQuery, err := treesitter.ReadQueryFromFile(lang, "injections")
	assert.Nil(t, err)

	query, err := lang.Query(injectionQuery)
	assert.Nil(t, err)

	expected := []injections.Capture{
		{
			Language: "php",
			Combined: true,
			Range: ts.Range{
				StartByte: 0,
				EndByte:   24,
				StartPoint: ts.Point{
					Row:    0,
					Column: 0,
				},
				EndPoint: ts.Point{
					Row:    1,
					Column: 18,
				},
			},
		},
		{
			Language: "php_only",
			Combined: false,
			Range: ts.Range{
				StartByte: 27,
				EndByte:   33,
				StartPoint: ts.Point{
					Row:    1,
					Column: 21,
				},
				EndPoint: ts.Point{
					Row:    1,
					Column: 27,
				},
			},
		},
		{
			Language: "php",
			Combined: true,
			Range: ts.Range{
				StartByte: 35,
				EndByte:   41,
				StartPoint: ts.Point{
					Row:    1,
					Column: 29,
				},
				EndPoint: ts.Point{
					Row:    2,
					Column: 1,
				},
			},
		},
		{
			Language: "php_only",
			Combined: false,
			Range: ts.Range{
				StartByte: 50,
				EndByte:   61,
				StartPoint: ts.Point{
					Row:    2,
					Column: 10,
				},
				EndPoint: ts.Point{
					Row:    2,
					Column: 21,
				},
			},
		},
		{
			Language: "php",
			Combined: true,
			Range: ts.Range{
				StartByte: 64,
				EndByte:   70,
				StartPoint: ts.Point{
					Row:    3,
					Column: 1,
				},
				EndPoint: ts.Point{
					Row:    3,
					Column: 7,
				},
			},
		},
	}

	actual := injections.Query(query, tree.RootNode(), src)
	assert.Equal(t, expected, actual)
}
