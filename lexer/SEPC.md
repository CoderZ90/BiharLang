# 🗣️ Bihari

> **Programming theek-thaak tareeka se.**

**Bihari** is a small, dynamically typed, interpreted programming language inspired by Bihari/Hinglish speech, local slang, and the concept of **jugaad**.

The language is intentionally funny and expressive on the surface, but underneath it is a real programming language project designed to explore:

* Lexing
* Parsing
* Abstract Syntax Trees
* Interpreters
* Runtime values
* Environments / scope
* Error handling
* Functions
* Control flow
* Language design
* Eventually bytecode / virtual machines

The interpreter is written entirely in **Go**.

---

# 🚧 Project Status

**Early design / pre-alpha**

The language specification is intentionally evolving.

The goal is to first create a small, coherent language and then expand it carefully.

---

# 🎯 Philosophy

Bihari is **not** supposed to be:

> "Python but all the keywords are translated into Bihari."

The language should have its own identity.

The central concept is:

## Jugaad

In Bihari, **jugaad is a real programming concept**, not merely a funny keyword.

The idea is:

> "Try something. If it works, great. If it doesn't, handle it somehow."

This gives the language a unique mechanism for dealing with uncertain or failure-prone operations.

The language should feel conversational:

```bh
rakha naam = "Babu"

jodi naam == "Babu" {
    bol "ka ho bhai!"
}
```

But it should still have well-defined semantics underneath.

---

# ✨ Example

A small Bihari program:

```bh
bol "Ka ho duniya!"

rakha x = 10
rakha y = 20

bol x + y

jodi x < y {
    bol "x chhota ba"
} na ta {
    bol "x bada ba"
}
```

Output:

```text
Ka ho duniya!
30
x chhota ba
```

---

# 📁 File Extension

Bihari source files use:

```text
.bh
```

Example:

```text
main.bh
hello.bh
jugaad.bh
```

---

# 🚀 CLI

The main executable is:

```bash
bihari
```

Run a program:

```bash
bihari chala main.bh
```

Start the REPL:

```bash
bihari baitho
```

Format a file:

```bash
bihari saaf main.bh
```

Show help:

```bash
bihari madad
```

Possible future commands:

```bash
bihari bana
bihari test
bihari package
bihari install
```

---

# 🧱 Core Syntax

## Variables

Variables are declared using:

```bh
rakha
```

Example:

```bh
rakha naam = "Babu"
rakha umar = 20
rakha paisa = 500
```

Variables can be reassigned:

```bh
rakha x = 10

x = 20
```

Bihari is dynamically typed.

Therefore this is valid:

```bh
rakha x = 10

x = "ka ho"
```

---

# 📦 Data Types

Initial version:

```text
number
string
boolean
null
array
function
```

Potential future types:

```text
object
error
jugaad
```

---

# 🔢 Numbers

Numbers can be used directly:

```bh
rakha x = 10
rakha y = 20

bol x + y
```

Supported arithmetic:

```text
+
-
*
/
%
```

Example:

```bh
bol 10 + 5
bol 10 - 5
bol 10 * 5
bol 10 / 5
bol 10 % 3
```

---

# 📝 Strings

Strings use double quotes:

```bh
rakha naam = "Babu"

bol naam
```

String concatenation:

```bh
rakha naam = "Babu"

bol "Ka ho " + naam
```

---

# ✅ Booleans

Bihari uses:

```bh
sach
jhoot
```

Example:

```bh
rakha active = sach
rakha banned = jhoot
```

---

# 🚫 Null

The null value is:

```bh
kuchnahi
```

Example:

```bh
rakha result = kuchnahi
```

---

# 🖨️ Output

The main output function is:

```bh
bol
```

Example:

```bh
bol "Hello World"
bol 10
bol sach
```

Multiple values may eventually be supported:

```bh
bol "Naam:", naam
```

---

# 🔀 Conditions

The equivalent of `if` is:

```bh
jodi
```

Example:

```bh
jodi x > 10 {
    bol "x bada ba"
}
```

The equivalent of `else` is:

```bh
na ta
```

Example:

```bh
jodi x > 10 {
    bol "bada ba"
} na ta {
    bol "chhota ba"
}
```

Possible future `else if` syntax:

```bh
jodi x > 100 {
    bol "bahut bada"
} na ta jodi x > 50 {
    bol "thik-thak bada"
} na ta {
    bol "chhota"
}
```

---

# 🔁 Loops

The initial loop construct is:

```bh
jabtak
```

Example:

```bh
rakha x = 0

jabtak x < 10 {
    bol x
    x = x + 1
}
```

This behaves similarly to a `while` loop.

---

# 🛑 Loop Control

Potential keywords:

```text
bas
aage
```

Possible meanings:

```bh
bas
```

→ break

```bh
aage
```

→ continue

These are provisional and can be changed during language design.

---

# 🛠️ Functions

Functions are declared using:

```bh
kaam
```

Example:

```bh
kaam jor(a, b) {
    lauta a + b
}
```

Call it:

```bh
rakha result = jor(10, 20)

bol result
```

Output:

```text
30
```

---

# 🔙 Return

Functions return values using:

```bh
lauta
```

Example:

```bh
kaam square(x) {
    lauta x * x
}
```

---

# 🧠 Scope

Bihari should eventually support lexical/block scope.

Example:

```bh
rakha x = 10

jodi sach {
    rakha y = 20
    bol y
}

bol x
```

`y` should not be accessible outside its block.

The exact scoping rules should be defined before implementation.

---

# 🧨 JUGAAD

## The Signature Feature of Bihari

Jugaad is the feature that should make Bihari more than a reskinned joke language.

The conceptual idea:

```text
Try something uncertain.
If it succeeds → get the result.
If it fails → get a failure value that can be handled.
```

Example:

```bh
rakha result = jugaad {
    risky_operation()
}
```

Instead of necessarily terminating the entire program, the operation can produce a special result.

Conceptually:

```text
Jugaad
├── Success(value)
└── Failure(error)
```

---

# 🛠️ Jugaad Handling

A possible syntax:

```bh
jugaad {
    rakha x = risky_operation()

    bol x
} pakka {
    bol "ho gaya bhai"
} na ta {
    bol "nahi hua bhai"
}
```

Meaning:

```text
jugaad
    ↓
try the operation

pakka
    ↓
success branch

na ta
    ↓
failure branch
```

The exact syntax is still experimental.

---

# 💥 Jugaad Values

A failed operation should ideally not behave exactly like a normal runtime exception.

Conceptually:

```text
Success(42)
```

or:

```text
Failure(
    message = "division by zero"
)
```

This allows programs to inspect failures.

Example:

```bh
rakha result = jugaad {
    divide(10, 0)
}

jodi result fail ba {
    bol "arre bhai, division nahi ho raha"
}
```

The exact syntax for checking a failed Jugaad is still to be designed.

---

# 🤡 Runtime Errors

Errors should have personality while remaining technically useful.

Example:

```text
ka ho bhai 💀

Error:
Expected number
Got string

Line 7:
    rakha x = "lauki" + 10

Jugaad karo ya code theek karo.
```

Errors should contain:

1. A funny message
2. The actual technical problem
3. Source line
4. Relevant code
5. Eventually a source location / stack trace

Humor should never replace useful debugging information.

---

# 🧮 Operators

Initial operators:

### Arithmetic

```text
+
-
*
/
%
```

### Comparison

```text
==
!=
<
>
<=
>=
```

### Logical

Potential syntax:

```text
aur
ya
ulta
```

Example:

```bh
jodi age > 18 aur citizen == sach {
    bol "allowed"
}
```

These keywords are provisional.

---

# 🧵 Comments

Single-line comments:

```bh
// ee comment ba
```

Example:

```bh
rakha x = 10 // temporary jugaad
```

Potential future multiline comments:

```bh
/*
    bada wala comment
*/
```

---

# 📚 Arrays

Potential syntax:

```bh
rakha dost = ["Babu", "Raju", "Pintu"]
```

Access:

```bh
bol dost[0]
```

Possible operations:

```bh
dost.push("Sonu")
dost.pop()
```

Array support should be added after the basic interpreter works.

---

# 🗺️ Language Roadmap

## v0.1 — Tiny Interpreter

Implement only:

```text
✓ Lexer
✓ Tokens
✓ Parser
✓ AST
✓ Numbers
✓ Strings
✓ Booleans
✓ Null
✓ Variables
✓ Assignment
✓ Arithmetic
✓ Comparison
✓ bol
✓ jodi
✓ na ta
✓ jabtak
```

Goal:

```bh
rakha x = 10
rakha y = 20

jodi x < y {
    bol x + y
}
```

must work.

---

## v0.2 — Functions

Add:

```text
✓ kaam
✓ parameters
✓ function calls
✓ lauta
✓ lexical scope
```

Example:

```bh
kaam jor(a, b) {
    lauta a + b
}

bol jor(10, 20)
```

---

## v0.3 — Jugaad

This is where Bihari gets its unique identity.

Implement:

```text
✓ Jugaad AST nodes
✓ Success values
✓ Failure values
✓ Jugaad execution
✓ Failure handling
✓ Useful runtime errors
```

---

## v0.4 — Collections

Add:

```text
✓ arrays
✓ indexing
✓ array operations
```

Potentially:

```text
✓ objects/maps
```

---

## v0.5 — Modules

Potential syntax:

```bh
laa "math"
```

or:

```bh
import "math"
```

This is intentionally undecided.

---

# 🧑‍💻 Interpreter Architecture

The implementation should follow a traditional interpreter pipeline:

```text
                 Bihari Source
                       │
                       ▼
                    Lexer
                       │
                       ▼
                    Tokens
                       │
                       ▼
                    Parser
                       │
                       ▼
                     AST
                       │
                       ▼
                  Evaluator
                       │
                       ▼
                 Runtime Values
                       │
                       ▼
                    Output
```

---

# 🔤 Lexer

The lexer converts source code into tokens.

Example:

```bh
rakha x = 10 + 20
```

becomes approximately:

```text
RAKHA
IDENTIFIER(x)
ASSIGN
NUMBER(10)
PLUS
NUMBER(20)
EOF
```

The lexer should be implemented manually in Go initially.

Do not depend on a parser generator for the first version.

---

# 🌳 AST

The parser should construct an Abstract Syntax Tree.

Example:

```bh
rakha x = 10 + 20
```

could become:

```text
VariableDeclaration
├── Name: x
└── Value
    └── BinaryExpression
        ├── Left: 10
        ├── Operator: +
        └── Right: 20
```

The AST should represent the meaning of the program rather than preserving every piece of syntax.

---

# 🧠 Evaluator

The evaluator walks the AST and executes it.

Example:

```text
BinaryExpression
    10
    +
    20
```

becomes:

```text
30
```

Variables should be stored in an environment.

Conceptually:

```text
Environment
┌──────────┬───────┐
│ x        │ 30    │
│ naam     │ Babu  │
│ active   │ true  │
└──────────┴───────┘
```

---

# 📦 Runtime Objects

Potential internal Go representation:

```text
Number
String
Boolean
Null
Array
Function
JugaadSuccess
JugaadFailure
```

The exact Go implementation should be designed during development.

---

# 📂 Suggested Go Project Structure

```text
bihari/
│
├── cmd/
│   └── bihari/
│       └── main.go
│
├── token/
│   └── token.go
│
├── lexer/
│   ├── lexer.go
│   └── lexer_test.go
│
├── ast/
│   └── ast.go
│
├── parser/
│   ├── parser.go
│   └── parser_test.go
│
├── object/
│   └── object.go
│
├── evaluator/
│   ├── evaluator.go
│   └── evaluator_test.go
│
├── repl/
│   └── repl.go
│
├── examples/
│   ├── hello.bh
│   ├── variables.bh
│   ├── conditions.bh
│   ├── functions.bh
│   └── jugaad.bh
│
├── go.mod
└── README.md
```

This structure may change as the language grows.

---

# 🧪 Example Programs

## Hello World

```bh
bol "Ka ho duniya!"
```

---

## Variables

```bh
rakha naam = "Babu"
rakha umar = 20
rakha paisa = 500

bol naam
bol umar
bol paisa
```

---

## Condition

```bh
rakha paisa = 100

jodi paisa >= 50 {
    bol "chai ho jayega"
} na ta {
    bol "paisa nahi ba bhai"
}
```

---

## Loop

```bh
rakha x = 0

jabtak x < 5 {
    bol x
    x = x + 1
}
```

---

## Function

```bh
kaam jor(a, b) {
    lauta a + b
}

rakha result = jor(10, 20)

bol result
```

---

## Jugaad

Experimental syntax:

```bh
jugaad {
    rakha result = risky_operation()
} pakka {
    bol "ho gaya bhai"
} na ta {
    bol "nahi hua 💀"
}
```

---

# 🧭 Design Principles

The language should follow these principles.

### 1. Funny surface, serious implementation

The language can be ridiculous.

The interpreter should not be.

---

### 2. Jokes should have technical meaning

Don't add features purely because the keyword sounds funny.

A funny concept should correspond to actual language semantics.

---

### 3. Keep the core small

The first version should be tiny.

Do not implement:

```text
classes
generics
async
concurrency
macros
modules
package manager
compiler
JIT
```

until the interpreter itself works.

---

### 4. Prefer understandable internals

The project is primarily a learning project.

Avoid unnecessary abstractions.

---

### 5. Errors must remain useful

Humor should enhance errors, not hide information.

Bad:

```text
bhai kuch toh phata hai 💀
```

Good:

```text
Error: cannot add string and number

Line 8:
    rakha x = "10" + 5

bhai ee addition na ho sakat ba 💀
```

---

# 🔮 Future Ideas

These are intentionally NOT part of the initial language.

Potential future experiments:

```text
• Jugaad expressions
• Error values
• Pattern matching
• Pipelines
• Maps/objects
• Modules
• Package manager
• Standard library
• File I/O
• HTTP requests
• JSON
• Concurrency
• Bytecode VM
• REPL improvements
• Debugger
• Formatter
• Language server
• VS Code extension
```

---

# 🧪 Possible "Cursed" Features

These should only be considered after the core language is stable.

### Magic runtime values

```bh
@last
@error
@args
@time
```

### Conversational expressions

```bh
jodi user logged_in ba {
    dashboard dikha
}
```

### Jugaad pipelines

```bh
jugaad
    fetch_user()
    se naam_nikaal()
    se uppercase_kar()
    dikha
```

These are experimental and should not be implemented until their semantics are clearly defined.

---

# 🛑 Important Development Rule

**Do not blindly add features because they sound funny.**

Before implementing a feature, answer:

```text
1. What does this syntax mean?
2. What AST node does it produce?
3. How does the evaluator execute it?
4. What runtime value does it produce?
5. What happens when it fails?
6. Why does this feature belong in Bihari?
```

If these questions cannot be answered clearly, the feature should not be implemented yet.

---

# 🤖 Development With AI

AI assistants may be used to help develop Bihari, but they should follow the existing language specification rather than inventing incompatible syntax.

When modifying the project:

1. Read the existing implementation first.
2. Preserve the language's existing semantics.
3. Do not introduce new syntax without documenting it.
4. Prefer small incremental changes.
5. Write tests for language features.
6. Explain parser/evaluator changes when requested.
7. Avoid rewriting the entire interpreter for small features.
8. Keep the implementation understandable to a beginner learning programming-language development.

---

# 📜 License

TBD.

---

# ❤️ Why Bihari Exists

Because programming languages don't always have to take themselves seriously.

But the **engineering behind them should be real.**

Bihari is an experiment in building a programming language from scratch while learning how lexers, parsers, ASTs, interpreters, runtime environments, and language semantics actually work.

> **Jugaad karo. Code chalao.**
>
> — Bihari
