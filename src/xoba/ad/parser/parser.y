%{ 
package parser 
import "fmt"

%} 
 
%union {
  node Node
} 
 
%token NUM

%% 

program: {}
| program NUM { fmt.Println("parsed"); }
;
