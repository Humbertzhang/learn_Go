```Go
name, ok := elements["Un"]
fmt.Println(name, ok)
Accessing an element of a map can return two values instead of just one.
The first value is the result of the lookup, the second tells us whether or not the lookup was successful.
In Go we often see code like this:
if name, ok := elements["Un"]; ok {
  fmt.Println(name, ok)
}
First we try to get the value from the map, then if it's successful we run the code inside of the block.

```
