## Just Another Toy Lang (JATL)

Another attempt at an interpreter after [Boa](https://github.com/MoMus2000/Boa) and [HODL](https://github.com/MoMus2000/HODLang)

Notes:

## General Notes: 
 - Expressions produce values, statements don't
 - Parsing Technique - Top Down - Recursive Descent - Pratt Parsing

 - Prefix Expr: -5, !false, !true
 - Infix  Expr: 5 / 10; 12 * 12;
 - Group  Expr: (5 * 5) + 10
 - Comp   Expr: 5 == 10; 5 < 10; 5 != 10;
 - Call   Expr: add(5, 10)
 - Ident  Expr: foo * bar / foo
 - FN     Expr: let add = fn(x, y) { return x + y };
 - If     Expr: let result = if (10 > 5) { true } else { false };


## Pratt Parsing:
 - A crucial part of this idea is that each token type can have two parsing functions associated with it, depending on the token’s position - infix or prefix.
 - Instead of associating parsing functions with grammar rules (defined in BNF or EBNF), Pratt associates these functions (which he calls “semantic code”) with single token types.
 - A prefix operator is an operator “in front of” its operand.
 - An infix operator sits between its operands, like this: 5 * 8
 - operator precedence. And alternative term for this is order of operations

