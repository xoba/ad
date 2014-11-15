%{ 
  package parser 
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
| '(' exp ')' { $$ = $2; }
| IDENT '(' exp ')' { $$.node = Function($1.node.S,$3.node); }
| IDENT '(' exp ',' exp ')' { }
|  '-' exp %prec '*' { $$.node = Negate($2.node);  }
|  exp '+' exp  {  $$.node = Binary('+',$1.node,$3.node);  }
|  exp '-' exp  {  $$.node = Binary('-',$1.node,$3.node);  }
|  exp '*' exp  {  $$.node = Binary('*',$1.node,$3.node);  }
|  exp '/' exp  {  $$.node = Binary('/',$1.node,$3.node);  }
|  exp '^' exp  {  $$.node = Binary('^',$1.node,$3.node);  }
;
