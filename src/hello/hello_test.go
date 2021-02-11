package main
import "testing"

func TestHello(t *testing.T){
	t.Run("saying hello with arguments", func(t *testing.T){
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want{
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("saying `Hello, world` in case of absence of arguments", func(t *testing.T){
		got := Hello("")
		want := "Hello, world"

		if got != want{
			t.Errorf("got %q want %q", got, want)
		}
	})
}
