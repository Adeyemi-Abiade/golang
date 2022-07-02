### Intro
The go programming was created by google in the year 2009, to be able to solve
issue like multi-threading in different types of app. 
The go language, or golang, as it is called is a compiled language which can be run on different os. i.e one compiled 'exec' file will run the same in any machine regardless of the os.

### Syntax
The go language is statically-typed but also allows the use of dynamic-type variables whose type can be inferred by the go compiler. The go compiler throws an error when a variable is
declared without using it. 
#### fmt (package)
The fmt  / format package in go is used to format input-output (i/o). It contains one such function which is very useful, this function id the 'Printf' function. It allows the interpolation of variables with the string output using some special operators. Some of these special operators are:
"%v" --> for outputting the value of a variable
"%T" --> for outputting the type pf the variable. N.B -- this must be an uppercase "T"

Another method is the scan method which accepts the variable for usre input as an argument.
This actually accepts a pointer to that variable so it can get the user inout and store it 
in that memory location.