package main; 

import "fmt" 
import "example.com/greetings"
import "log"

func main() {
  
  log.SetPrefix("Greetings: "); 
  log.SetFlags(0); 
  
  names := []string { "Daniel", "Anna", "Conor", "Ryan" };
  message, err := greetings.Hello(names); 
  
  if err != nil { log.Fatal(err) }; 

  for k, v := range message {
    fmt.Printf("Name: %v\n", k);  
    fmt.Printf("Message: %v\n\n", v);
  }

}
