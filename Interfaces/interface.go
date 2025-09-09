package interface

// decoupling via behavior
// Define minimal interfaces; implicit satisfaction
// Notes: Keep interfaces small (“interface segregation”); accept interfaces, return concrete types.

type Storer interface {
    Get(ctx context.Context, key string) ([]byte, error)
    Put(ctx context.Context, key string, val []byte) error
}

type S3Store struct{ /* fields */ }
func (s *S3Store) Get(ctx context.Context, k string) ([]byte, error) { /* ... */ return nil, nil }
func (s *S3Store) Put(ctx context.Context, k string, v []byte) error { /* ... */ return nil }

type Service struct{ store Storer }

func NewService(store Storer) *Service { return &Service{store: store} }

func (s *Service) Handle(ctx context.Context, k string) error {
    b, err := s.store.Get(ctx, k)
    if err != nil { return err }
    // mutate and write back
    return s.store.Put(ctx, k, bytes.ToUpper(b))
}

// Using standard interfaces (io.Reader/Writer)

func CopyUpper(dst io.Writer, src io.Reader) error {
    r := bufio.NewReader(src)
    for {
        b, err := r.ReadByte()
        if err == io.EOF { return nil }
        if err != nil { return err }
        if 'a' <= b && b <= 'z' { b -= 32 }
        if _, err := dst.Write([]byte{b}); err != nil { return err }
    }
}

// Type assertions & type switches
var w io.Writer = os.Stdout

if f, ok := w.(*os.File); ok {
    _ = f.Sync()
}

switch v := anyVal.(type) {
case fmt.Stringer:
    fmt.Println("stringer:", v.String())
case int:
    fmt.Println("int:", v)
default:
    fmt.Printf("unknown: %T\n", v)
}

// nil interface vs typed nil pitfalls

var e error          // nil interface value (type, value both nil)
var _ = e == nil     // true

var *MyErr = nil
var err error = (*MyErr)(nil)
fmt.Println(err == nil) // false: interface has (type=*MyErr, value=nil)

// Notes: An interface is nil only if both its dynamic type and value are nil. 
// This trips error handling—return nil concrete error, not a typed-nil in an error interface.

