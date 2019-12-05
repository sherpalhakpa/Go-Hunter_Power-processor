# Go-Hunter_Power-processor

This program written in Go takes a program written in Hunter – Power (Language grammer: https://drive.google.com/file/d/1lXhzJVH70KI92RnhjtiwogLTwe18HOS1/view?usp=sharing), and outputs the tokens and lexemes into a new file. 
---------------------------------------------------------------------
The program  runs like this: 
---------------------------------------------------------------------
prompt>./go-hunter-scanner input.txt 
Processing Input File input.txt 
14 Tokens produced Results in Output File input.out 
prompt>

The tokens in the grammar are: 
- STRING 
- REAL_CONST 
- INT_CONST 
- BEGIN 
- END  
- WRITE 
- ID 
- COLON 
- POINT 
- ASSIGN 
- TIMES 
- DIVISION 
- PLUS 
- MINUS 
- POWER 
- LPAREN 
- RPAREN 


Given the following program written in this language: 
$x <= "hi" : 
#d <= 12 : 
%r <= 3.44 : 
%k <= ( #d + #d ) ^ 1.4 : 
WRITE "this is the program" : 
WRITE %k 
 
The output file should look like: 
ID[STRING]: x 
ASSIGN 
STRING: hi 
COLON 
D[INT]: d 
ASSIGN 
INT_CONST: 12 
OLON 
ID[REAL]: r 
ASSIGN 
REAL_CONST: 3.44 
COLON 
ID[REAL]: k 
ASSIGN 
LPAREN 
ID[INT]: d 
PLUS ID[INT]: d 
RPAREN 
POWER 
REAL_CONST: 1.4 
COLON 
WRITE 
STRING: this is the program 
COLON 
WRITE 
ID[REAL]: k 
