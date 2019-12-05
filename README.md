# Go-Hunter_Power-processor

This program written in Go takes a program written in Hunter–Power (Grammer:https://drive.google.com/file/d/1lXhzJVH70KI92RnhjtiwogLTwe18HOS1/view?usp=sharing), and outputs the tokens and lexemes into a new file. 
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


Given the following program written in this language: <br/>
	$x <= "hi" : <br/>
	#d <= 12 : <br/>
	#d <= 12 : <br/>
	%r <= 3.44 : <br/>
	%k <= ( #d + #d ) ^ 1.4 : <br/>
	WRITE "this is the program" : <br/>
	WRITE %k <br/>
 
The output file should look like: <br/>
	ID[STRING]: x <br/>
	ASSIGN <br/>
	STRING: hi <br/>
	COLON <br/>
	D[INT]: d <br/>
	ASSIGN <br/>
	INT_CONST: 12 <br/> 
	OLON <br/>
	ID[REAL]: r <br/>
	ASSIGN <br/>
	REAL_CONST: 3.44 <br/> 
	COLON <br/>
	ID[REAL]: k <br/>
	ASSIGN <br/>
	LPAREN <br/>
	ID[INT]: d <br/>
	PLUS ID[INT]: d <br/>
	RPAREN <br/>
	POWER <br/>
	REAL_CONST: 1.4 <br/>
	COLON <br/>
	WRITE <br/>
	STRING: this is the program <br/>
	COLON <br/>
	WRITE <br/>
	ID[REAL]: k<br/> 
