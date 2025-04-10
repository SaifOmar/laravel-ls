package provider

import (
	"github.com/laravel-ls/laravel-ls/file"
	"github.com/laravel-ls/laravel-ls/project"
)

type Language struct {
	CompletionProviders  []CompletionProvider
	DiagnosticsProviders []DiagnosticProvider
	DefinitionProviders  []DefinitionProvider
	CodeActionProviders  []CodeActionProvider
	HoverProviders       []HoverProvider
}

type Manager struct {
	project   *project.Project
	providers []Provider
	languages map[file.Type]Language
}

func NewManager(providers ...Provider) *Manager {
	mgr := &Manager{
		languages: map[file.Type]Language{},
	}
	return mgr.Add(providers...)
}

func (m *Manager) Init(ctx InitContext) {
	var err error

	m.project, err = project.New(ctx.RootPath)
	if err != nil {
		ctx.Logger.WithError(err).Warn("failed to find binary")
	}

	for _, provider := range m.providers {
		provider.Init(ctx)
	}
}

func (m *Manager) Add(providers ...Provider) *Manager {
	for _, provider := range providers {
		provider.Register(m)
		m.providers = append(m.providers, provider)
	}
	return m
}

func (m *Manager) Register(typ file.Type, provider any) {
	lang, ok := m.languages[typ]
	if !ok {
		m.languages[typ] = Language{
			CompletionProviders:  []CompletionProvider{},
			DiagnosticsProviders: []DiagnosticProvider{},
			DefinitionProviders:  []DefinitionProvider{},
			HoverProviders:       []HoverProvider{},
		}
		lang = m.languages[typ]
	}

	if completion, ok := provider.(CompletionProvider); ok {
		lang.CompletionProviders = append(lang.CompletionProviders, completion)
	}

	if diagnostic, ok := provider.(DiagnosticProvider); ok {
		lang.DiagnosticsProviders = append(lang.DiagnosticsProviders, diagnostic)
	}

	if definition, ok := provider.(DefinitionProvider); ok {
		lang.DefinitionProviders = append(lang.DefinitionProviders, definition)
	}

	if hover, ok := provider.(HoverProvider); ok {
		lang.HoverProviders = append(lang.HoverProviders, hover)
	}

	if codeAction, ok := provider.(CodeActionProvider); ok {
		lang.CodeActionProviders = append(lang.CodeActionProviders, codeAction)
	}

	m.languages[typ] = lang
}

func (m *Manager) CodeAction(ctx CodeActionContext) {
	ctx.Project = m.project
	if providers, ok := m.languages[ctx.File.Type]; ok {
		for _, provider := range providers.CodeActionProviders {
			provider.ResolveCodeAction(ctx)
		}
	}
}

func (m *Manager) Completion(ctx CompletionContext) {
	ctx.Project = m.project
	if providers, ok := m.languages[ctx.File.Type]; ok {
		for _, provider := range providers.CompletionProviders {
			provider.ResolveCompletion(ctx)
		}
	}
}

func (m *Manager) Diagnostics(ctx DiagnosticContext) {
	ctx.Project = m.project
	if providers, ok := m.languages[ctx.File.Type]; ok {
		for _, provider := range providers.DiagnosticsProviders {
			provider.Diagnostic(ctx)
		}
	}
}

func (m *Manager) ResolveDefinition(ctx DefinitionContext) {
	ctx.Project = m.project
	if providers, ok := m.languages[ctx.File.Type]; ok {
		for _, provider := range providers.DefinitionProviders {
			provider.ResolveDefinition(ctx)
		}
	}
}

func (m *Manager) Hover(ctx HoverContext) {
	ctx.Project = m.project
	if providers, ok := m.languages[ctx.File.Type]; ok {
		for _, provider := range providers.HoverProviders {
			provider.Hover(ctx)
		}
	}
}
