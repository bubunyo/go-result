# go-result `Experimental`

This is an attempt to implement the rust `Result<T, Error>` Type in go.
This is purely experimental. Do not use it in production.

## Use Case 

if you define a  method ReadFile that returns a Result? eg
```go
func ReadFile(input string) result.Result[[]string] {
	file, err := ioutil.ReadFile(input)
	if err != nil {
		return result.Error[[]string](err)
	}
	fileContent := string(file)
	return result.Ok(strings.Split(fileContent, "\n"))
}

```

You can read the contents of a file using

```go 
for l := range ReadFile("./rick-roll-1").Result() {
  fmt.Println(l)
}

/* 
Output:  

Never gonna give you up
Never gonna let you down
Never gonna run around and desert you
*/
```

You can also chain results like this

```go
song := ReadFile("./rick-roll-1").
  Then(func(p1 []string) result.Resolver[[]string] {
    return result.Ok(append(p1, ReadFile("./rick-roll-2").Result()...))
	}).
	Result()


  /* 
Output:

Never gonna give you up
Never gonna let you down
Never gonna run around and desert you

Never gonna make you cry
Never gonna say goodbye
Never gonna tell a lie and hurt you
*/
``` 


