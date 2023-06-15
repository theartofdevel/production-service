package tracing

import (
	"bytes"
	"context"
	"encoding/json"
	"path"
	"reflect"

	"github.com/iancoleman/strcase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Attributed interface {
	Attributes() []attribute.KeyValue
}

func TraceIVal(ctx context.Context, name string, val interface{}) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	av, ok := attributeValue(reflect.ValueOf(val))
	if ok {
		span.SetAttributes(attribute.KeyValue{
			Key:   attribute.Key(name),
			Value: av,
		})
	}
}

func TraceVal(ctx context.Context, name string, val string) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	span.SetAttributes(attribute.String(name, val))
}

func TraceObjWP(ctx context.Context, prefix string, obj interface{}) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	traceStruct(span, prefix, obj)
}

func TraceObj(ctx context.Context, obj interface{}) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	traceStruct(span, "", obj)
}

func Error(ctx context.Context, err error) {
	if err == nil {
		return
	}
	if span := trace.SpanFromContext(ctx); span.IsRecording() {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}

func traceStruct(span trace.Span, prefix string, obj interface{}) {
	if attributed, ok := obj.(Attributed); ok {
		span.SetAttributes(attributed.Attributes()...)
	} else {
		span.SetAttributes(AttributesFrom(prefix, obj)...)
	}
}

const (
	tag = "trace"
	dot = '.'
)

func AttributesFrom(prefix string, obj interface{}) []attribute.KeyValue {
	if prefix != "" {
		prefix += "_"
	}

	return attributesFrom(prefix, obj)
}

func attributesFrom(prefix string, obj interface{}) []attribute.KeyValue {
	if obj == nil {
		return nil
	}

	rv := reflect.ValueOf(obj)

	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}

	if isUnsupported(rv) {
		return nil
	}

	rt := rv.Type()
	fn := rt.NumField()
	attrs := make([]attribute.KeyValue, 0, fn)
	kv := attribute.KeyValue{}
	pref, buf := prefixAndBuffer(rt)

	for i := 0; i < fn; i++ {
		f := rt.Field(i)
		if !f.IsExported() {
			continue
		}

		name := f.Name

		if tag, ok := f.Tag.Lookup(tag); ok {
			if tag == "-" {
				continue
			}

			name = tag
		}

		if av, ok := attributeValue(rv.Field(i)); ok {
			buf.Reset()
			buf.Write(pref)
			buf.WriteString(strcase.ToSnake(name))

			kv.Key = attribute.Key(prefix + buf.String())
			kv.Value = av

			attrs = append(attrs, kv)
		}
	}

	return attrs
}

func prefixAndBuffer(rt reflect.Type) (prefix []byte, buf bytes.Buffer) {
	if sf, ok := rt.FieldByName("_"); ok {
		if tag, ok := sf.Tag.Lookup(tag); ok {
			buf.WriteString(tag)
			buf.WriteByte(dot)
		}
	}

	if buf.Len() == 0 {
		buf.WriteString(path.Base(rt.PkgPath()))
		buf.WriteByte(dot)
		buf.WriteString(strcase.ToSnake(rt.Name()))
		buf.WriteByte(dot)
	}

	prefix = buf.Bytes()

	return prefix, buf
}

func attributeValue(v reflect.Value) (av attribute.Value, ok bool) {
SwitchKind:
	switch v.Kind() {
	case reflect.Struct, reflect.Map, reflect.Interface, reflect.Slice, reflect.Array:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			av = attribute.StringValue(err.Error())
		} else {
			av = attribute.StringValue(string(b))
		}
	case reflect.String:
		av = attribute.StringValue(v.String())
	case reflect.Int:
		av = attribute.IntValue(int(v.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		av = attribute.Int64Value(int64(v.Uint()))
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		av = attribute.Int64Value(v.Int())
	case reflect.Float32, reflect.Float64:
		av = attribute.Float64Value(v.Float())
	case reflect.Bool:
		av = attribute.BoolValue(v.Bool())
	case reflect.Pointer:
		v = v.Elem()
		goto SwitchKind
	default:
		// unsupported types
		ok = true
	}

	return av, !ok
}

func isUnsupported(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Struct, reflect.Map:
		return false
	}

	return true
}
