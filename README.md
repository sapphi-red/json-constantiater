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
|JsonIterFastest|[GitHub](https://github.com/json-iterator/go) ConfigCompatibleWithStandardLibrary|
|Jettison|[GitHub](https://github.com/wI2L/jettison) NoHTMLEscaping,NoUTF8Coercion,UnsortedMap,NoCompact|
|EasyJson|[GitHub](https://github.com/mailru/easyjson)|
|GoJay|[GitHub](https://github.com/francoispqt/gojay)|
|SegmentJson|[GitHub](https://github.com/segmentio/encoding/blob/master/json/README.md)|
|GoJson|[GitHub](https://github.com/goccy/go-json)|
|ConstantiateNonOptimized|This one. Not using Extended options.|
|Constantiate|This one. Using Extended options.|

### SmallStruct
![Small](https://user-images.githubusercontent.com/49056869/92930893-a421f680-f47d-11ea-9bf9-4b8657a9b803.png)

### MediumStruct
![Medium](https://user-images.githubusercontent.com/49056869/92930900-a5532380-f47d-11ea-9114-744fd514a704.png)

### LargeStruct
![Large](https://user-images.githubusercontent.com/49056869/92930896-a5532380-f47d-11ea-8aef-07db44175bc4.png)
