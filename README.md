# Summary
This repo will be a verbose "journal" of my progress following the (Learn Go with test)[https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/] book.

There is no defined goals or outcome apart from gaining a greater understanding of Go and really learning the fundamentals of TDD and testing to verify that your entended actions return the expected result.
Personally I do not believe TDD is the best way all programming should be written, where as having *good* tests is much more important than coverage etc.


## Who cares?
Hopefully not you, nor should anyone.

This is a way for me to try and be less of a noob and progress my career into something I really value.


## Other learning
Go has many fantastic resources:
1. A Tour Of Go - Fantastic starting point with an interactive playgroud. (link)[https://go.dev/tour/]
2. Go By Example - A hands-on introduction to Go using annotated example programs. (link)[https://gobyexample.com/]
3. Effective Go - The holy grail blog post by the Go team written in 2009, gives tips for writing clear, idiomatic Go code. (link)[https://go.dev/doc/effective_go]


## Interesting resources
1. (Writing An Interpreter In Go)[https://interpreterbook.com/]
   - How to build an interpreter for a C-like programming language from scratch
   - What a lexer, a parser and an Abstract Syntax Tree (AST) are and how to build your own
   - What closures are and how and why they work
   - What the Pratt parsing technique and a recursive descent parser is
   - What others talk about when they talk about built-in data structures
   - What REPL stands for and how to build one

2. (Writing A Compiler In Go)[https://compilerbook.com/]
We'll learn a lot about computers, how they work, what machine code and opcodes are, what the stack is and how to work with stack pointers and frame pointers, what it means to define a calling convention, and much more.
   - We define our own bytecode instructions, specifying their operands and their encoding. Along the way, we also build a mini-disassembler for them.
   - We write a compiler that takes in a Monkey AST and turns it into bytecode by emitting instructions
   - At the same time we build a stack-based virtual machine that executes the bytecode in its main loop

Also:
   - build a symbol table and a constant pool
   - do stack arithmetic
   - generate jump instructions
   - build frames into our VM to execute functions with local bindings and arguments!
   - add built-in functions to the VM
   - get real closures working in the virtual machine and learn why closure-compilation is so tricky

