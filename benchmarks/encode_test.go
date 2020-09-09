package benchmark

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	gojay "github.com/francoispqt/gojay"
	gojson "github.com/goccy/go-json"
)

func Benchmark_Encode_SmallStruct_EncodingJson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(NewSmallPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_SmallStruct_JsonIter(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(NewSmallPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_SmallStruct_JsonIterFastest(b *testing.B) {
	var json = jsoniter.ConfigFastest
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(NewSmallPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_SmallStruct_GoJay(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojay.MarshalJSONObject(NewSmallPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_SmallStruct_GoJson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojson.Marshal(NewSmallPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_SmallStruct_ConstantiateNonOptimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewSmallPayloadNonOptimized().NewJsonMarshal()
	}
}

func Benchmark_Encode_SmallStruct_Constantiate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewSmallPayload().NewJsonMarshal()
	}
}

func Benchmark_Encode_MediumStruct_EncodingJson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(NewMediumPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_MediumStruct_JsonIter(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(NewMediumPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_MediumStruct_JsonIterFastest(b *testing.B) {
	var json = jsoniter.ConfigFastest
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(NewMediumPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_MediumStruct_GoJay(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojay.MarshalJSONObject(NewMediumPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_MediumStruct_GoJson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojson.Marshal(NewMediumPayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_MediumStruct_ConstantiateNonOptimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewMediumPayloadNonOptimized().NewJsonMarshal()
	}
}

func Benchmark_Encode_MediumStruct_Constantiate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewMediumPayload().NewJsonMarshal()
	}
}

func Benchmark_Encode_LargeStruct_EncodingJson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(NewLargePayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_LargeStruct_JsonIter(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(NewLargePayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_LargeStruct_JsonIterFastest(b *testing.B) {
	var json = jsoniter.ConfigFastest
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(NewLargePayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_LargeStruct_GoJay(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojay.MarshalJSONObject(NewLargePayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_LargeStruct_GoJson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := gojson.Marshal(NewLargePayload()); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_Encode_LargeStruct_ConstantiateNonOptimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewLargePayloadNonOptimized().NewJsonMarshal()
	}
}

func Benchmark_Encode_LargeStruct_Constantiate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewLargePayload().NewJsonMarshal()
	}
}
