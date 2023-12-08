package greetings 

import "testing"
import "regexp"

func TestHelloName(t *testing.T) {
  name := []string { "Badman" }; 
  want := regexp.MustCompile(`\b` + name[0] + `\b`); 

  msg, err := Hello(name); 

  if !want.MatchString(msg["Badman"]) || err != nil {
    t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err); 
  }
}


func TestHelloEmpty(t *testing.T) {
  msg, err := Hello([]string { "" }); 
  if msg[""] != "" || err == nil {
    t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err); 
  }
}
