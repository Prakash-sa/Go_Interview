package interfaces

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
)

// Storer is a minimal interface representing a blob store.
type Storer interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Put(ctx context.Context, key string, val []byte) error
}

// S3Store is a stub implementing Storer.
type S3Store struct{}

func (s *S3Store) Get(ctx context.Context, k string) ([]byte, error) { return []byte(k), nil }
func (s *S3Store) Put(ctx context.Context, k string, v []byte) error { return nil }

// Service depends on a Storer interface, enabling decoupled testing.
type Service struct{ store Storer }

func NewService(store Storer) *Service { return &Service{store: store} }

func (s *Service) Handle(ctx context.Context, k string) error {
	b, err := s.store.Get(ctx, k)
	if err != nil {
		return err
	}
	return s.store.Put(ctx, k, bytes.ToUpper(b))
}

// CopyUpper demonstrates standard io.Reader/io.Writer usage.
func CopyUpper(dst io.Writer, src io.Reader) error {
	r := bufio.NewReader(src)
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if 'a' <= b && b <= 'z' {
			b -= 32
		}
		if _, err := dst.Write([]byte{b}); err != nil {
			return err
		}
	}
}

// TypeSwitchInfo inspects dynamic types using a type switch.
func TypeSwitchInfo(anyVal any) string {
	switch v := anyVal.(type) {
	case fmt.Stringer:
		return "stringer:" + v.String()
	case int:
		return "int"
	default:
		return fmt.Sprintf("unknown:%T", v)
	}
}

// TypedNilPitfall shows how typed nils keep the interface non-nil.
type MyErr struct{}

func (MyErr) Error() string { return "boom" }

func TypedNilPitfall() (plainNil bool, typedNil bool) {
	var e error
	plainNil = e == nil

	var err error = (*MyErr)(nil)
	typedNil = err == nil // false: interface has dynamic type *MyErr
	return plainNil, typedNil
}

// WriterExample illustrates type assertion.
func WriterExample() (synced bool) {
	var w io.Writer = os.Stdout
	if f, ok := w.(*os.File); ok {
		_ = f.Sync()
		return true
	}
	return false
}
