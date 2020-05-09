#### Notes from my Go learning
- Only needed packages must be imported. Unused imports will result in a compilation error
- `import` declarations must follow the `package` declaration.
- `slices` are equivalent to `python`'s `list`s.
- When using `os.Args`, the slice's first element is the actual `go` binary. The passed in arguments start from `Args[1]`.
- Comments are `JS`-esque
- It's good convention to describe the imported packages after they are imported.
- Only one type of loop: `for`, with its three syntax variations covering all loop cases. Pretty neat.
- Parentheses are never used around the three components of a `for` loop. Same applies to the `if` statement.
- `range` is equivalent to `python`'s `enumerate`. An index and the value are produced.
- `go` doesn't permit the use of unused local variables. Memory efficient! In such cases where one is needed, a blank identifier, `_` should be used.
- The syntax, `s := ""`, which defines and assigns a variable can only be used within a function, and not for package-level variables.
- `map` is equivalen to `python`'s `dict`. The key may be of any type whose values can be compared using the `==` operator. `string`s are the most common. Technically, since the datatypes are declared on start, the equivalence isn't with a Python dictionary, rather, it's equivalent to `python`'s `collections.defaultdict`. Ordering, unlike Python's is unordered.
- `map[string]int` means a dictionary with keys of the `string` datatype and values of `int` datatype. The inbuilt command, `make` is used to actually create the dictionary. i.e., `make(map[string]int)`. Neeto!

