# json-constantiater

## Usage
```shell
$ go get -u github.com/sapphi-red/json-constantiater
$ json-constantiater -input p.go -output p_json.go
```

When `output` argument is not passed, is will behave as if `-output {{input name}}_constantiated.go` is passed.

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
	n int `json:",small"`
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

### `unsigned`
If this option were set, it assumes the number is 0 or over 0.
It will panic when the number is below 0.

```go
type A struct {
	n int `json:",unsigned"`
}

func main() {
	a := A {
		n: 2
	}
	fmt.Println(string(a.NewJsonMarshal())) // {"n":2}

	b := A {
		n: -5
	}
	fmt.Println(string(b.NewJsonMarshal())) // panic!
}
```

### `nonnil`
If this option were set, it assumes the pointer is never nil.
It will panic when the value is nil.
This overrides `omitempty` option.

```go
type A struct {
	n *string `json:",nonnil"`
}

func main() {
	str := "str"
	a := A {
		n: &str
	}
	fmt.Println(string(a.NewJsonMarshal())) // {"n":"str"}

	b := A {
		n: nil
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
![Small](https://user-images.githubusercontent.com/49056869/97721180-77e53680-1b0c-11eb-8d47-35fdd73d33dd.png)

### MediumStruct
![Medium](https://user-images.githubusercontent.com/49056869/97721178-77e53680-1b0c-11eb-8f78-9f3c92ecde3b.png)

### LargeStruct
![Large](https://user-images.githubusercontent.com/49056869/97721175-76b40980-1b0c-11eb-82b3-8e63ffb6ce19.png)
