# notes

> By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc.

[interface-names](http://golang.org/doc/effective_go.html#interface-names)

> Since if and switch accept an initialization statement, it's common to see one used to set up a local variable.

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

[if](http://golang.org/doc/effective_go.html#if)

> An aside: The last example in the previous section demonstrates a detail of how the := short declaration form works. The declaration that calls os.Open reads, ... which looks as if it declares d and err. Notice, though, that err appears in both statements. This duplication is legal: err is declared by the first statement, but only re-assigned in the second. This means that the call to f.Stat uses the existing err variable declared above, and just gives it a new value.
In a := declaration a variable v may appear even if it has already been declared, provided:
 - this declaration is in the same scope as the existing declaration of v (if v is already declared in an outer scope, the declaration will create a new variable §),
 - the corresponding value in the initialization is assignable to v, and
 - there is at least one other variable in the declaration that is being declared anew.

[redelcartion](http://golang.org/doc/effective_go.html#redeclaration)

This means you can do lots of this:

```
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
```

blank identifier

```go
for _, value := range arr {
// ...
}
```

parallel assigment :

```
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

> There is no automatic fall through, but cases can be presented in comma-separated lists.

[switch](http://golang.org/doc/effective_go.html#switch)


This is the first example I've seen of type inspection / reflection:

> A switch can also be used to discover the dynamic type of an interface variable. Such a type switch uses the syntax of a type assertion with the keyword type inside the parentheses. If the switch declares a variable in the expression, the variable will have the corresponding type in each clause. It's also idiomatic to reuse the name in such cases, in effect declaring a new variable with the same name but a different type in each case.

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type)
```

[type switch](http://golang.org/doc/effective_go.html#type_switch)


