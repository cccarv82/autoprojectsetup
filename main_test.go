package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {

  // captura a saída padrão
  var buf bytes.Buffer
  old := os.Stdout // armazena os handles atuais
  r, w, _ := os.Pipe()
  os.Stdout = w

  main()
  
  // restaura os handles
  w.Close()
  os.Stdout = old
  out, _ := io.ReadAll(r)

  got := string(out)
  want := "Hello World\n"

  if got != want {
    t.Errorf("saída incorreta: obtido '%s', esperado '%s'", got, want)
  }
}
