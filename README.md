# json-constantiater

```shell
$ go run constantiater.go -input p.go -output p_json.go
```

## Extended options
### `noescape`
If this option were set, it will skip escaping `"`/`\`/`\n`/`\r`/`\t`.
It may output invalid JSON when it is set to a field which might include those letters.

```go
type A struct {
  name string `json:"name,noescape"`
}

func main() {
  a := A {
    name: "john"
  }
  fmt.Println(string(a.NewJsonMarshal())) // {"name":"john"}

  b := A {
    name: "jo\"hn"
  }
  fmt.Println(string(b.NewJsonMarshal())) // {"name":"jo"hn"}
}
```

### `len*`
This option sets max string length of the field.

```go
type A struct {
  name string `json:"name,len4"`
}
type B struct {
  name string `json:"name"`
}

func main() {
  a1 := A {
    name: "john"
  }
  fmt.Println(string(a1.NewJsonMarshal())) // fastest

  a2 := A {
    name: "john titor"
  }
  fmt.Println(string(a2.NewJsonMarshal())) // slowest

  b := B {
    name: "john"
  }
  fmt.Println(string(b.NewJsonMarshal())) // normal
}
```
