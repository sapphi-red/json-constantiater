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

## options for array/map types
```go
// value:",noescape"
type A []string

// key:",noescape"
// value:",noescape"
type M map[string]string
```

## Benchmarks

|Name in graph|Description|
|---|---|
|EncodingJson|[doc](https://golang.org/pkg/encoding/json/)|
|JsonIter|[GitHub](https://github.com/json-iterator/go) ConfigCompatibleWithStandardLibrary|
|JsonIterFastest|[GitHub](https://github.com/json-iterator/go) ConfigFastest|
|GoJay|[GitHub](https://github.com/francoispqt/gojay)|
|GoJson|[GitHub](https://github.com/goccy/go-json)|
|ConstantiateNonOptimized|This one. Not using Extended options.|
|Constantiate|This one. Using Extended options.|

### SmallStruct
![Small](https://user-images.githubusercontent.com/49056869/92892761-049c3e00-f454-11ea-9cc6-bab5eccc713e.png)

### MediumStruct
![Medium](https://user-images.githubusercontent.com/49056869/92892769-05cd6b00-f454-11ea-9402-f07dc1cbbdca.png)

### LargeStruct
![Large](https://user-images.githubusercontent.com/49056869/92892776-06fe9800-f454-11ea-86be-878693039a1b.png)
