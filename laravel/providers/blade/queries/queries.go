package queries

import (
	_ "embed"

	"github.com/laravel-ls/laravel-ls/parser"
	"github.com/laravel-ls/laravel-ls/treesitter"
	"github.com/laravel-ls/laravel-ls/treesitter/language"
)

const (
	// QueryCaptureBladeFileName  = "filename.blade.php"
	QueryCaptureDirectives     = "blade.punctuation"
	QueryCaptureDirectiveNames = "blade.directive.name"
)

// we should have here the list of blade directives that we are going to support

var BladeDirectives = map[string]string{
	"@if":           "if statement",
	"@for":          "for loop",
	"@csrf":         "CSRF token",
	"@foreach":      "foreach loop",
	"@forelse":      "forelse loop",
	"@else":         "else clause",
	"@elseif":       "elseif clause",
	"@section":      "section definition",
	"@extends":      "template extension",
	"@yield":        "yield content",
	"@include":      "include template",
	"@component":    "component definition",
	"@slot":         "slot content",
	"@continue":     "continue loop",
	"@break":        "break loop",
	"@isset":        "check if set",
	"@empty":        "check if empty",
	"@auth":         "auth check",
	"@guest":        "guest check",
	"@method":       "HTTP method",
	"@route":        "route helper",
	"@json":         "JSON output",
	"@php":          "PHP code block",
	"@endif":        "end of if statement",
	"@endfor":       "end of for loop",
	"@endforeach":   "end of foreach loop",
	"@endforelse":   "end of forelse loop",
	"@endsection":   "end of section",
	"@endextends":   "end of extends",
	"@endyield":     "end of yield content",
	"@endinclude":   "end of include template",
	"@endcomponent": "end of component definition",
	"@endslot":      "end of slot content",
	"@endcontinue":  "end of continue loop",
	"@endbreak":     "end of break loop",
	"@endisset":     "end of isset check",
	"@endempty":     "end of empty check",
	"@endauth":      "end of auth check",
	"@endguest":     "end of guest check",
	"@endmethod":    "end of HTTP method",
	"@endroute":     "end of route helper",
	"@endjson":      "end of JSON output",
	"@endphp":       "end of PHP code block",
	"@push":         "push content to a stack",
	"@endpush":      "end of push content to a stack",
	"@prepend":      "prepend content to a stack",
	"@endprepend":   "end of prepend content to a stack",
	"@stack":        "render stack content",
	"@error":        "display form validation errors",
	"@verbatim":     "output raw HTML without Blade compilation",
	"@endverbatim":  "end of raw HTML output",
	"@env":          "retrieve the value of an environment variable",
	"@unless":       "conditional clause that works like !if",
	"@endunless":    "end of unless statement",
	"@can":          "check if a user has permission",
	"@cannot":       "check if a user cannot perform an action",
	"@endcan":       "end of can check",
	"@endcannot":    "end of cannot check",
}

func findBladeFileName(file *parser.File, lang language.Identifier) treesitter.CaptureSlice {
	r, err := file.FindTags(lang, QueryCaptureDirectives)
	if err != nil {
		return treesitter.CaptureSlice{}
	}
	return r
}

func GetBladeNode(file *parser.File) treesitter.CaptureSlice {
	return append(findBladeFileName(file, language.PHP),
		findBladeFileName(file, language.PHPOnly)...)
}
