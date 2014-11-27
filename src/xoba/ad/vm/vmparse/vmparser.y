%{ 
  package vmparse
%} 
 
%union {
  node *Node
} 
 
%token NUM
%token IDENT

%right '='
%left '+' '-'
%left '*' '/' '^'

%% 

statement: IDENT ':' '=' exp { c := yylex.(*context); c.lhs,c.rhs = $1.node, $4.node; }
;

exp: NUM { $$ = $1; } 
| IDENT { $$ = $1; } 
| IDENT '[' NUM ']' { $$.node = IndexedIdentifier($1.node,$3.node); }
| '(' exp ')' { $$ = $2; }
| IDENT '(' exp ')' { $$.node = Function($1.node.S,$3.node); }
| IDENT '(' exp ',' exp ')' { $$.node = Function($1.node.S,$3.node,$5.node); }
|  '-' exp %prec '*' { $$.node = Negate($2.node);  }
|  exp '+' exp  {  $$.node = Function("add",$1.node,$3.node);  }
|  exp '-' exp  {  $$.node = Function("subtract",$1.node,$3.node);  }
|  exp '*' exp  {  $$.node = Function("multiply",$1.node,$3.node);  }
|  exp '/' exp  {  $$.node = Function("divide",$1.node,$3.node);  }
|  exp '^' exp  {  $$.node = Function("pow",$1.node,$3.node);  }
;
