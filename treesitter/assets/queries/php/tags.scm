; --------------------------------------------------
;  Config
; --------------------------------------------------

; config() calls
(function_call_expression
  function: (name) @function (#eq? @function "config")
  arguments: (arguments
    . (argument [
       (string (string_content)?) @config.key
       (encapsed_string . (string_content) .) @config.key
       (encapsed_string "\"" . "\"") @config.key
    ])
  ))

; config()->type() calls
(member_call_expression
  object: (
    function_call_expression
    function: (name) @object.name (#eq? @object.name "config")
    arguments: (arguments "(" . ")"))
  name: (name) @function.name
  (#any-of? @function.name "get" "integer" "float" "string" "boolean" "array")
  arguments: (arguments
     . (argument [
        (string (string_content)?) @config.key
        (encapsed_string . (string_content) .) @config.key
        (encapsed_string "\"" . "\"") @config.key
     ])
  ))

; Config::type() calls.
(scoped_call_expression
  scope: [
    (qualified_name (name) @class)
    (name) @class
  ] (#eq? @class "Config")
  name: (name) @method 
  (#any-of? @method "get" "integer" "float" "string" "boolean" "array")
  arguments: (arguments
     . (argument [
        (string (string_content)?) @config.key
        (encapsed_string . (string_content) .) @config.key
        (encapsed_string "\"" . "\"") @config.key
     ])
  ))

; Config::getMany() calls.
(scoped_call_expression
  scope: [
    (qualified_name (name) @class)
    (name) @class
  ] (#eq? @class "Config")
  name: (name) @method (#eq? @method "getMany")
  arguments: (arguments
    . (argument
        (array_creation_expression
          (array_element_initializer
            [
              (string (string_content)?) @config.key
              (encapsed_string . (string_content) .) @config.key
              (encapsed_string "\"" . "\"") @config.key
            ]
    )))
  ))

; config()->getMany() calls.
(member_call_expression
  object: (
    function_call_expression
    function: (name) @object.name (#eq? @object.name "config")
    arguments: (arguments "(" . ")"))
  name: (name) @function.name (#eq? @function.name "getMany")
  arguments: (arguments
    . (argument
        (array_creation_expression
          (array_element_initializer
            [
             (string (string_content)?) @config.key
             (encapsed_string . (string_content) .) @config.key
             (encapsed_string "\"" . "\"") @config.key
             ]
    )))
  ))

; --------------------------------------------------
;  App
; --------------------------------------------------

; app('key')
(function_call_expression
  function: (name) @function (#eq? @function "app")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @app.service
        (encapsed_string . (string_content) .) @app.service
        (encapsed_string "\"" . "\"") @app.service
    ])
  ))


; app()->make('key')
; app()->bound('key')
; app()->isShared('key')
(member_call_expression
  object: (
    function_call_expression
    function: (name) @object.name (#eq? @object.name "app")
    arguments: (arguments "(" . ")"))
  name: (name) @function.name
  (#any-of? @function.name "make" "bound" "isShared")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @app.service
        (encapsed_string . (string_content) .) @app.service
        (encapsed_string "\"" . "\"") @app.service
    ])
  ))

; App::make('key')
; App::bound('key')
; App::isShared('key')
(scoped_call_expression
  scope: [
    (qualified_name (name) @class)
    (name) @class
  ] (#eq? @class "App")
  name: (name) @method
  (#any-of? @method "make" "bound" "isShared")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @app.service
        (encapsed_string . (string_content) .) @app.service
        (encapsed_string "\"" . "\"") @app.service
    ])
  ))

; --------------------------------------------------
;  View
; --------------------------------------------------

; view() calls
(function_call_expression
  function: (name) @function (#eq? @function "view")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @view.name
        (encapsed_string . (string_content) .) @view.name
        (encapsed_string "\"" . "\"") @view.name
    ])
  ))

; Route::view() calls.
(scoped_call_expression
  scope: [
    (qualified_name (name) @class)
    (name) @class
  ] (#eq? @class "Route")
  name: (name) @method (#eq? @method "view")
  arguments: (arguments
    (argument) ; First parameter is the route.
    . (argument [
        (string (string_content)?) @view.name
        (encapsed_string . (string_content) .) @view.name
        (encapsed_string "\"" . "\"") @view.name
    ])
  ))

; --------------------------------------------------
;  Environment
; --------------------------------------------------

; env() calls
(function_call_expression
  function: (name) @function (#eq? @function "env")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @env.key
        (encapsed_string . (string_content) .) @env.key
        (encapsed_string "\"" . "\"") @env.key
    ])
))

; Env::get() calls.
(scoped_call_expression
  scope: [
    (qualified_name (name) @class)
    (name) @class
  ] (#eq? @class "Env")
  name: (name) @method (#eq? @method "get")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @env.key
        (encapsed_string . (string_content) .) @env.key
        (encapsed_string "\"" . "\"") @env.key
    ])
))

; --------------------------------------------------
;  Assets
; --------------------------------------------------

; asset() calls
(function_call_expression
  function: (name) @function (#eq? @function "asset")
  arguments: (arguments
    . (argument [
        (string (string_content)?) @asset.filename
        (encapsed_string (string_content)?) @asset.filename
    ])
  ))


; ---------------------------------------------------
; Blade
; ---------------------------------------------------

; ---------------------------------------------------
; Blade Directives (@directive)
; ---------------------------------------------------

; Captures the '@' symbol and the following identifier/name,
; which constitutes the start of a Blade directive.
; NOTE: The effectiveness and specifics might depend on whether
; you're using a standard PHP grammar or one specifically
; enhanced or designed for Blade (like tree-sitter-blade).
; This query attempts a general approach.

(
  "@" @blade.punctuation  ; Capture the literal "@" token
  .                       ; Immediately followed by...
  [                       ; ...either
    (name)                ; a 'name' node (common in PHP grammar)
  ] @blade.directive.name ; Capture the node containing the directive's name (if, foreach, etc.)
)

; --- Alternative / More Specific (If using a Blade-aware grammar) ---
; If your grammar has specific nodes for Blade directives, queries like these would be more robust:
;
; (directive                ; Matches a generic directive node
;   "@" @blade.punctuation
;   name: (_) @blade.directive.name
; )
;
; or for specific directives:
;
; (if_directive             ; Matches a specific @if directive node
;   "@" @blade.punctuation
;   name: (_) @blade.directive.name ; Might capture 'if'
;   parameters: (_)? @blade.parameters
; ) @if_directive.structure
;
; (endif_directive          ; Matches a specific @endif directive node
;   "@" @blade.punctuation
; ) @endif_directive.structure

