# Closures

- Compiled functions no longer gets emitted by the compiler, `OpClosure` gets emitted instead.
- By making all functions closures, they all retain access to free variables.
- When entering a new function, we load all local variables and closures from previous functions onto the current one.