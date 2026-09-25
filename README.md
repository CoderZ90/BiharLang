# FLOW
```
 Source Code (.bh file)
    ↓
 [LEXER] → Breaks text into meaningful pieces (tokens)
    ↓
 [TOKENS] → ["rakha", "x", "=", "10"]
    ↓
 [PARSER] → Organizes tokens into a tree structure
    ↓
 [AST] → Tree representing program structure
    ↓
 [EVALUATOR] → Walks the tree and executes it
    ↓
 Output
```

# PHASE

## Phase 1: Foundation 
 - Understand tokens and lexing conceptually
 - Set up project structure
 - Build a simple lexer that handles basic tokens

## Phase 2: Parser Basics
    - Understand AST concept
    - Build parser for simple expressions (numbers, arithmetic)

## Phase 3: Evaluator
    - Evaluate simple expressions
    - Add variables and environment

## Phase 4: Control Flow
    - Add conditions (jodi/na ta)
    - Add loops (jabtak)

## Phase 5: Functions & Beyond
    - Functions (kaam/lauta)
    - Jugaad feature
    - Error handling with personality

# LEXERs

```js
  rakha x = 10 + 20
```

- RAKHA - keyword for variable declaration
- IDENTIFIER(x) - variable name
- ASSIGN - assignment operator (=)
- NUMBER(10) - numeric literal
- PLUS - arithmetic operator (+)
- NUMBER(20) - numeric literal
- EOF - end of file (marks the end)

# 

| Bihari     | English / Programming meaning           | Example                     |
| ---------- | --------------------------------------- | --------------------------- |
| `rakha`    | `let` / `var` — declare variable        | `rakha x = 10`              |
| `bol`      | `print`                                 | `bol "hello"`               |
| `jodi`     | `if`                                    | `jodi x > 10 {}`            |
| `na ta`    | `else`                                  | `na ta {}`                  |
| `jabtak`   | `while`                                 | `jabtak x < 10 {}`          |
| `kaam`     | `function`                              | `kaam add(a, b) {}`         |
| `lauta`    | `return`                                | `lauta x + y`               |
| `sach`     | `true`                                  | `rakha x = sach`            |
| `jhoot`    | `false`                                 | `rakha x = jhoot`           |
| `kuchnahi` | `null`                                  | `rakha x = kuchnahi`        |
| `bas`      | `break`                                 | `bas`                       |
| `aage`     | `continue`                              | `aage`                      |
| `aur`      | `AND`                                   | `x > 5 aur y < 10`          |
| `ya`       | `OR`                                    | `x > 5 ya y > 20`           |
| `ulta`     | `NOT`                                   | `ulta sach`                 |
| `jugaad`   | `try / attempt / recoverable operation` | `jugaad { ... }`            |
| `pakka`    | `success / confirmed`                   | `jugaad {} pakka {}`        |
| `na ta`    | `failure / otherwise`                   | `jugaad {} na ta {}`        |
| `laa`      | `import`                                | `laa "math"`                |
| `dikha`    | display/show                            | `dikha user`                |
| `chala`    | run/execute                             | CLI: `bihari chala main.bh` |
| `baitho`   | enter/start REPL                        | `bihari baitho`             |
| `saaf`     | format/clean                            | `bihari saaf main.bh`       |
| `madad`    | help                                    | `bihari madad`              |


## Completed
✅ TokenType definitions
✅ Lexer with readChar/readPosition state machine
✅ Identifier & keyword lookup
✅ Number parsing
✅ String parsing  
✅ Whitespace handling
✅ Full test suite PASSING