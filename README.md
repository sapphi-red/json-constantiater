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
![Small](https://user-images.githubusercontent.com/49056869/92576330-6168ee80-f2c4-11ea-8636-40af0be47b59.png)

### MediumStruct
![Medium](https://user-images.githubusercontent.com/49056869/92576336-629a1b80-f2c4-11ea-978c-380cda97e407.png)

### LargeStruct
![Large](https://user-images.githubusercontent.com/49056869/92576335-62018500-f2c4-11ea-95aa-18c881ca669e.png)
