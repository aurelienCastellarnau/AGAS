** Package helpers **

*** debuggingTools.go ***

Bring some development wrappers, printers and utilitaries

Struct DebugBox {}

(*DebugBox).GetFunctionName(interface{})
    - return the package.name string of the function passed as parameter
(*DebugBox).GetStack(interface{})
    - return a formatted stack print
(*DebugBox).Debug
    - return an instance of DebugBox
    This way, functions can be called after a DebugBox instanciation:
    box := &helpers.DebugBox{}
    or chained after a call to Debug
    helpers.Debug().GetStack()