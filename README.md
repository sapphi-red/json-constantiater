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

### `omitnano`
This options omit nanoseconds of time.Time.

```go
type A struct {
  t string
}
type B struct {
  t string `json:",omitnano"`
}

func main() {
  a := A {
    t: time.Now()
  }
  fmt.Println(string(a.NewJsonMarshal())) // {"t":"2020-09-11T20:51:06.5260311+09:00"}

  b := B {
    t: time.Now()
  }
  fmt.Println(string(b.NewJsonMarshal())) // {"t":"2020-09-11T20:51:06+09:00"}
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
![Small](https://user-images.githubusercontent.com/49056869/92896352-424e9600-f457-11ea-90aa-90134b2597a5.png)

### MediumStruct
![Medium](https://user-images.githubusercontent.com/49056869/92896716-948fb700-f457-11ea-8a92-489064c13351.png)

### LargeStruct
![Large](https://user-images.githubusercontent.com/49056869/92896709-93f72080-f457-11ea-8a10-e28e6bbd1b89.png)
