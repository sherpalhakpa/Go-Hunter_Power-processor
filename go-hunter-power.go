package main

import (
"log"
"os"
"fmt"
"strings"
"io/ioutil"
"strconv"
)

func main(){
  //clear contents of file if file already exists
 outputSecond := os.Args[1][:strings.Index(os.Args[1],".")] + ".out"
 outputFile := os.Args[1][:strings.Index(os.Args[1],".")] + ".xtra"
 out, err := os.OpenFile(outputSecond,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 if err != nil {
  log.Fatal("Cannot create file", err)
  }
 file, err := os.OpenFile(outputFile,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 if err != nil {
  log.Fatal("Cannot create file", err)
  }
 out.Truncate(0)
 out.Seek(0,0)
 file.Truncate(0)
 file.Seek(0,0)
 lexemes := lexemeGenerator()
 j := 0 //for tokens produced
 line := 1 //for lines
 for i:= range lexemes{
  allLexeme := tokenize(lexemes[i],&line)
  PrintLexim(allLexeme,&j)
  }
 fmt.Println(j, "Tokens Produced")
}
//---------------------------------------------------------------
//function to generate a slice of lexemes from the
//..code text File
//returns a slice of string containing each lexeme as an string
//no value parameters
//Uses the commandline Args to access text file

func lexemeGenerator() []string {
 readFile := os.Args[1] //testfile name from commandline command
 b,err := ioutil.ReadFile(string(readFile))
 if err != nil{
 fmt.Println(err)
 }
 fmt.Println("Processing Input File " + string(readFile))
 str:= string(b) //convert read byte to string
 lexemes := []string{}
 //first split by line
 linesOfCode := []string(strings.Split(str,"\n"))
 for i:= range linesOfCode{
  temp := []string {}
  //can't split "space" inside quote
  if strings.ContainsAny(linesOfCode[i],"\"") {
    start := strings.Index(linesOfCode[i],"\"")
    end := strings.LastIndex(linesOfCode[i],"\"")
    firstPart:= strings.Split(linesOfCode[i][:start-1]," ")
    quotePart:= linesOfCode[i][start:end+1]
    lastPart:= linesOfCode[i][len(linesOfCode[i])-1]
    lexemes = append(lexemes, firstPart...)
    lexemes = append(lexemes, quotePart)
    lexemes = append(lexemes, string(lastPart))
    } else { //everything else just split by "space"
      temp = strings.Split(linesOfCode[i]," ")
      for j := range temp{
       lexemes = append(lexemes,temp[j])
       }
      }
    temp = temp[:0] //erase temp content
  }
 outputFile := os.Args[1][:strings.Index(os.Args[1],".")] + ".out"
 outputFileTwo := os.Args[1][:strings.Index(os.Args[1],".")] + ".xtra"
 fmt.Println("Results in Output File", outputFile)
 fmt.Println("Structure Details in Output File", outputFileTwo)
 return lexemes
}
//------------------------------------------------------------------
//function to tokenize each lexim
//given a string checks conditions and creates a type
//..lexeme struct related to the lexeme
//lexme struct contains all details about lexim including its token type
func tokenize(lexeme string, line * int) Lexeme {
  var lexemeDetails Lexeme
  if (strings.HasPrefix(lexeme, "$")){
    lexemeDetails.Lexeme = string(lexeme[1:])
    lexemeDetails.Type = "STRING"
    lexemeDetails.Value = ""
    lexemeDetails.Token = "ID"
  } else if (strings.HasPrefix(lexeme, "#")){
    lexemeDetails.Lexeme = string(lexeme[1:])
    lexemeDetails.Type = "INT"
    lexemeDetails.Value = ""
    lexemeDetails.Token = "ID"
  } else if (strings.HasPrefix(lexeme, "%")){
    lexemeDetails.Lexeme = string(lexeme[1:])
    lexemeDetails.Type = "REAL"
    lexemeDetails.Value = ""
    lexemeDetails.Token = "ID"
  } else if (strings.Compare(lexeme,"BEGIN")==0){
    lexemeDetails.Lexeme = "BEGIN"
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "BEGIN"
  } else if (strings.Compare(lexeme,"END")==0){
    lexemeDetails.Lexeme = "END"
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "END"
  } else if (strings.Compare(lexeme,"WRITE")==0){
    lexemeDetails.Lexeme = "WRITE"
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "WRITE"
  } else if (strings.Compare(lexeme,"WRITE")==0){
    lexemeDetails.Lexeme = "WRITE"
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "WRITE"
  } else if (strings.HasPrefix(lexeme,"\"")&& strings.HasSuffix(lexeme,"\"")){
    lexemeDetails.Lexeme = lexeme[1:len(lexeme)-1]
    lexemeDetails.Type = ""
    lexemeDetails.Value = lexeme[1:len(lexeme)-1]
    lexemeDetails.Token = "STRING"
  } else if (isNumber(lexeme)!= "NO"){
    lexemeDetails.Lexeme = lexeme
    lexemeDetails.Type = ""
    lexemeDetails.Value = lexeme
    lexemeDetails.Token = isNumber(lexeme)
  } else if (strings.Compare(lexeme,":")==0){
    lexemeDetails.Lexeme = lexeme
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "COLON"
    *line++
  } else if (strings.Compare(lexeme,"<=")==0){
    lexemeDetails.Lexeme = lexeme
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "ASSIGN"
  } else if (strings.Compare(lexeme,"*")==0){
    lexemeDetails.Lexeme = lexeme
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "TIMES"
  } else if (strings.Compare(lexeme,"/")==0){
    lexemeDetails.Lexeme = lexeme
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "DIVISION"
  } else if (strings.Compare(lexeme,"+")==0){
    lexemeDetails.Lexeme = lexeme
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "PLUS"
  } else if (strings.Compare(lexeme,"-")==0){
    lexemeDetails.Lexeme = lexeme
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "MINUS"
  } else if (strings.Compare(lexeme,"^")==0){
    lexemeDetails.Lexeme = lexeme
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "POWER"
  } else if (strings.Compare(lexeme,"(")==0){
    lexemeDetails.Lexeme = lexeme
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "LPAREN"
  } else if (strings.Compare(lexeme,")")==0){
    lexemeDetails.Lexeme = lexeme
    lexemeDetails.Type = ""
    lexemeDetails.Value = ""
    lexemeDetails.Token = "RPAREN"
  } else if (strings.Compare(lexeme,"")==0){

  } else {
    fmt.Println("Lexical Error at Line", *line)
  }
return lexemeDetails
}

//func to indetify if a string is number
//takes a string as Input
//and returns the type of Number
//INT_CONST or REAL_CONST
func isNumber(number string) string{
 _, err := strconv.ParseFloat(number, 64)
 _, err1 := strconv.Atoi(number)
 if (err1 == nil && !strings.ContainsAny(number, ".")){
  return "INT_CONST"
  }else if err == nil {
    return "REAL_CONST"
    }else {
      return "NO"
    }
}

//Given a struct Lexeme
//Prints the content of the lexeme
//takes the Lexeme struct and a counters as parameters
//counter for number of tokens
func PrintLexim(lexeme Lexeme,j *int){
  outputSecond := os.Args[1][:strings.Index(os.Args[1],".")] + ".out"
  outputFile := os.Args[1][:strings.Index(os.Args[1],".")] + ".xtra"
  out, err := os.OpenFile(outputSecond,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal("Cannot create file", err)
   }
  file, err := os.OpenFile(outputFile,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
   if err != nil {
       log.Fatal("Cannot create file", err)
  }
  if (!emptyLexeme(lexeme)){
  *j++
  file.WriteString("Token no ")
  file.WriteString(fmt.Sprintf("%d",*j))
  file.WriteString(":\n")
  file.WriteString("    Lexeme:   "+lexeme.Lexeme+"\n")
  file.WriteString("    Type:     "+lexeme.Type+"\n")
  file.WriteString("    Value:    "+lexeme.Value+"\n")
  file.WriteString("    Token:    "+lexeme.Token+"\n")
  file.WriteString("\n")
  if (lexeme.Type != ""){
    out.WriteString(lexeme.Token)
    out.WriteString("["+lexeme.Type+"]: "+lexeme.Lexeme+"\n")
  } else if (lexeme.Token == "REAL_CONST" || lexeme.Token == "INT_CONST" || lexeme.Token == "STRING") {
    out.WriteString(lexeme.Token+ ": " + lexeme.Lexeme + "\n")
  } else {
    out.WriteString(lexeme.Token + "\n")
  }
  }
}
//---------------------------------------------------------------
//Given a Lexeme struct
//indetifies if lexeme is valid
//i.e. notImpty
func emptyLexeme(lexeme Lexeme) bool {
  if (lexeme.Lexeme == ""&&lexeme.Type==""&&lexeme.Value==""&&lexeme.Token==""){
      return true
    }
  return false
}

//---------------------------------------------------------------
//user defined struct to hold lexeme information
type Lexeme struct  {
  Lexeme string
  Type string
  Value string
  Token string
}
