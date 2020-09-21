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
Set this option to omit nanoseconds of time.Time.

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

### `small`
If this option were set, it assumes the number is 0 or is between 0 and 999.
It will panic when the number is over 999 or under 0.

```go
type A struct {
  n string `json:",small"`
}

func main() {
  a := A {
    n: 2
  }
  fmt.Println(string(a.NewJsonMarshal())) // {"n":2}

  b := A {
    n: 2000
  }
  fmt.Println(string(b.NewJsonMarshal())) // panic!
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
|JsonIterFastest|[GitHub](https://github.com/json-iterator/go) ConfigFastest|
|Jettison|[GitHub](https://github.com/wI2L/jettison) NoHTMLEscaping,NoUTF8Coercion,UnsortedMap,NoCompact|
|EasyJson|[GitHub](https://github.com/mailru/easyjson)|
|GoJay|[GitHub](https://github.com/francoispqt/gojay)|
|SegmentJson|[GitHub](https://github.com/segmentio/encoding/blob/master/json/README.md)|
|GoJson|[GitHub](https://github.com/goccy/go-json)|
|ConstantiateNonOptimized|This one. Not using Extended options.|
|Constantiate|This one. Using Extended options.|

### SmallStruct
![Small](https://user-images.githubusercontent.com/49056869/92948633-a04d9e80-f494-11ea-83c6-6360ea5c35f9.png)

### MediumStruct
![Medium](https://user-images.githubusercontent.com/49056869/92948635-a2176200-f494-11ea-9a22-a84ef6ea86c8.png)

### LargeStruct
![Large](https://user-images.githubusercontent.com/49056869/92948634-a17ecb80-f494-11ea-9559-a384bd509c00.png)
