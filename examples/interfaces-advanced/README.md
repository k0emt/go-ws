# Custom cat with Go

Demonstrate working with interfaces by implementing a customer writer.

## Usage

`go run cat.go spicy-bacon.txt`

Output will show the number of the line, the length, the contents, and the biggest word on the line.

The need for the counter in the struct isn't obvious until you have a large enough file for the writer to be called repeatedly.
During testing, it was observed that a file with more than 6000 lines of Bacon Ipsum was needed.

If you build the code, then you will need to do `./cat spicy-bacon.txt`
