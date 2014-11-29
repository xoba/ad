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

program: { }
| program statement { c := yylex.(*context); c.statements = append(c.statements, $2.node); } 
;
 
statement: IDENT equals exp { $$.node = NewStatement($1.node,$3.node); }
| func equals exp {  $$.node = NewStatement($1.node,$3.node); }
;

equals: '='
| ':' '='
;

exp: simple { $$ = $1; } 
| '(' exp ')' { $$ = $2; }
|  '-' exp { $$.node = Negate($2.node);  }
|  exp '*' exp  {  $$.node = Function("multiply",$1.node,$3.node);  }
|  exp '/' exp  {  $$.node = Function("divide",$1.node,$3.node);  }
|  exp '+' exp  {  $$.node = Function("add",$1.node,$3.node);  }
|  exp '-' exp  {  $$.node = Function("subtract",$1.node,$3.node);  }
|  exp '^' exp  {  $$.node = Function("pow",$1.node,$3.node);  }
;

simple: NUM { $$ = $1; } 
| IDENT { $$ = $1; } 
| IDENT '[' NUM ']' { $$.node = IndexedIdentifier($1.node,$3.node); }
| func { $$ = $1; } 
;

func:  IDENT '(' args ')' { $$.node = FunctionArgs($1.node.S,$3.node); }
;

args: exp { $$.node =  NewArgList($1.node); }
|  args ',' exp { $$.node = $1.node.AddChild($3.node); }
