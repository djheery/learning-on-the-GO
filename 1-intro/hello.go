package main

import "fmt"
import "rsc.io/quote"

func main() {
  fmt.Println("Hello, World!"); 
  quote_me_daddy(); 
}

func quote_me_daddy() {
  fmt.Println(quote.Go()); 
}


