package benchmark

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	gojay "github.com/francoispqt/gojay"
	gojson "github.com/goccy/go-json"
)

func Benchmark_Encode_SmallStruct_EncodingJson(b *testing.B) {
	s := NewSmallPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(s); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_SmallStruct_JsonIter(b *testing.B) {
	s := NewSmallPayload()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(s); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_SmallStruct_JsonIterFastest(b *testing.B) {
	s := NewSmallPayload()
	var json = jsoniter.ConfigFastest
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(s); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_SmallStruct_GoJay(b *testing.B) {
	s := NewSmallPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojay.MarshalJSONObject(s); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_SmallStruct_GoJson(b *testing.B) {
	s := NewSmallPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojson.Marshal(s); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_SmallStruct_ConstantiateNonOptimized(b *testing.B) {
	s := NewSmallPayloadNonOptimized()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = s.NewJsonMarshal()
	}
}

func Benchmark_Encode_SmallStruct_Constantiate(b *testing.B) {
	s := NewSmallPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = s.NewJsonMarshal()
	}
}

func Benchmark_Encode_MediumStruct_EncodingJson(b *testing.B) {
	m := NewMediumPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(m); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_MediumStruct_JsonIter(b *testing.B) {
	m := NewMediumPayload()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(m); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_MediumStruct_JsonIterFastest(b *testing.B) {
	m := NewMediumPayload()
	var json = jsoniter.ConfigFastest
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(m); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_MediumStruct_GoJay(b *testing.B) {
	m := NewMediumPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojay.MarshalJSONObject(m); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_MediumStruct_GoJson(b *testing.B) {
	m := NewMediumPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojson.Marshal(m); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_MediumStruct_ConstantiateNonOptimized(b *testing.B) {
	m := NewMediumPayloadNonOptimized()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = m.NewJsonMarshal()
	}
}

func Benchmark_Encode_MediumStruct_Constantiate(b *testing.B) {
	m := NewMediumPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = m.NewJsonMarshal()
	}
}

func Benchmark_Encode_LargeStruct_EncodingJson(b *testing.B) {
	l := NewLargePayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(l); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_LargeStruct_JsonIter(b *testing.B) {
	l := NewLargePayload()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(l); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_LargeStruct_JsonIterFastest(b *testing.B) {
	l := NewLargePayload()
	var json = jsoniter.ConfigFastest
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(l); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_LargeStruct_GoJay(b *testing.B) {
	l := NewLargePayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojay.MarshalJSONObject(l); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_LargeStruct_GoJson(b *testing.B) {
	l := NewLargePayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojson.Marshal(l); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_LargeStruct_ConstantiateNonOptimized(b *testing.B) {
	l := NewLargePayloadNonOptimized()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = l.NewJsonMarshal()
	}
}

func Benchmark_Encode_LargeStruct_Constantiate(b *testing.B) {
	l := NewLargePayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = l.NewJsonMarshal()
	}
}
