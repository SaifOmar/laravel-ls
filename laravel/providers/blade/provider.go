package blade

import (
	// "path"

	"strings"

	"github.com/laravel-ls/laravel-ls/file"
	// "github.com/laravel-ls/laravel-ls/laravel"
	"github.com/laravel-ls/laravel-ls/laravel/providers/blade/queries"
	"github.com/laravel-ls/laravel-ls/lsp/protocol"
	"github.com/laravel-ls/laravel-ls/provider"
	"github.com/laravel-ls/laravel-ls/treesitter/php"
)

type Provider struct {
	rootPath string
	bd       map[string]string
}

func NewProvider() *Provider {
	return &Provider{
		bd: queries.BladeDirectives,
	}
}

func (p *Provider) Init(ctx provider.InitContext) {
	p.rootPath = ctx.RootPath
}

func (p *Provider) Register(manager *provider.Manager) {
	manager.Register(file.TypeBlade, p)
}

// we will have a list of blade directives that should be avaialble for completion when use types @
// it should be the @ symbol and then the directive name then /n and @end then the directive name

func (p *Provider) ResolveCompletion(ctx provider.CompletionContext) {
	if ctx.File.Type != file.TypeBlade {
		return
	}

	node := queries.GetBladeNode(ctx.File).At(ctx.Position)
	if node == nil {
		return
	}

	text := php.GetStringContent(node, ctx.File.Src)
	if text == "" {
		return
	}

	// Suggest matching directives
	for directive, desc := range p.bd {
		directiveName := strings.TrimPrefix(directive, "@")
		// Match with @ (e.g., "@if")
		if strings.HasPrefix(text, "@") && strings.HasPrefix(directive, text) {
			ctx.Publish(protocol.CompletionItem{
				Label:  directive,
				Detail: "Blade: " + desc,
				Kind:   protocol.CompletionItemKindKeyword,
			})
		}
		// Match without @ (e.g., "if" -> "@if")
		if !strings.HasPrefix(text, "@") && strings.HasPrefix(directiveName, text) {
			ctx.Publish(protocol.CompletionItem{
				Label:  directive,
				Detail: "Blade: " + desc,
				Kind:   protocol.CompletionItemKindKeyword,
			})
		}
	}
}

// resolve view() calls to view files.
// func (p *Provider) ResolveDefinition(ctx provider.DefinitionContext) {
// 	node := queries.ViewNames(ctx.File).At(ctx.Position)
//
// 	if node != nil {
// 		name := php.GetStringContent(node, ctx.File.Src)
//
// 		if len(name) < 1 {
// 			return
// 		}
//
// 		viewFile := laravel.ViewFromName(name)
//
// 		ctx.Logger.Debug(viewFile)
//
// 		fullPath, found := p.fs.find(p.rootPath, viewFile.Filename())
//
// 		ctx.Logger.Debugf("%s %v", fullPath, found)
//
// 		if found {
// 			ctx.Publish(protocol.Location{
// 				URI: path.Join(p.rootPath, fullPath),
// 			})
// 		}
// 	}
// }

// func (p *Provider) ResolveCompletion(ctx provider.CompletionContext) {
// node := queries.ViewNames(ctx.File).At(ctx.Position)
//
// if node != nil {
// 	text := php.GetStringContent(node, ctx.File.Src)
//
// 	results, err := p.fs.search(p.rootPath, text)
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("failed to search view files")
// 		return
// 	}
//
// 	for _, result := range results {
// 		ctx.Publish(protocol.CompletionItem{
// 			Label:  result.Name(),
// 			Detail: result.Path(),
// 			Kind:   protocol.CompletionItemKindFile,
// 		})
// 	}
// }
// }

// func (p *Provider) Diagnostic(ctx provider.DiagnosticContext) {
// // Find all view calls in the file.
// for _, capture := range queries.ViewNames(ctx.File) {
// 	name := php.GetStringContent(&capture.Node, ctx.File.Src)
//
// 	// Report diagnostic if view does not exist.
// 	if !p.fs.exists(p.rootPath, name) {
// 		ctx.Publish(provider.Diagnostic{
// 			Range:    capture.Node.Range(),
// 			Severity: protocol.SeverityError,
// 			Message:  "View not found",
// 		})
// 	}
// }
// }

// func (p *Provider) Hover(ctx provider.HoverContext) {
// node := queries.ViewNames(ctx.File).At(ctx.Position)
//
// if node != nil {
// 	name := php.GetStringContent(node, ctx.File.Src)
// 	if len(name) < 1 {
// 		return
// 	}
//
// 	if view, found := p.fs.findView(p.rootPath, name); found {
// 		ctx.Publish(provider.Hover{
// 			Content: view.Path(),
// 		})
// 	}
// }
// }
