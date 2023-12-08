package greetings; 

import "fmt"; 
import "errors"; 
import "math/rand"

func Hello(names []string) (map[string]string, error) {
  // Makes a map (HashMap) of strings 
  // map[key-type]value-type
  messages := make(map[string] string); 
  
  for _, name := range names {
    // fmt.Printf("%d", i); 
    if name == "" { return nil, errors.New("empty name"); }

    message := fmt.Sprintf(randomFormat(), name); 
    messages[name] = message;
  }

  // := is a shorthand for declaring and initializing a variable on one line
  return messages, nil; 
}

func randomFormat() string {
  formats := []string { 
    "Hi, %v. Did you know you're a bitch?", 
    "Yo, Bitch. I don't care what you're name is. you're a bitch. Okay %v", 
    "Wait %v. Who do you think you are. Actually I know. You're a bitch", 
    "Alright %v, I think you know what I have to say about you... Bitch", 
    "Welcome %v, you seem great... PSYCHE YOURE A BITCHHHHH",
  }; 

  return formats[rand.Intn(len(formats))]; 
}
