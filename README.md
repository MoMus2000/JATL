## Just Another Toy Lang (JATL)

Another attempt at an interpreter after [Boa](https://github.com/MoMus2000/Boa) and [HODL](https://github.com/MoMus2000/HODLang)

Notes:
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

