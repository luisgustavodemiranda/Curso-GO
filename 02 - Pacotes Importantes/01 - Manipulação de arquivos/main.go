package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func estudarArquivo() error {
	// Um diretório temporário evita sobrescrever arquivos do estudante.
	pasta, err := os.MkdirTemp("", "curso-go-arquivos-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(pasta)

	arquivo, err := os.Create(pasta + "/mensagem.txt")
	if err != nil {
		return err
	}
	defer arquivo.Close()

	if _, err := arquivo.WriteString("Estudando arquivos em Go!\n"); err != nil {
		return err
	}
	// Volta ao início para ler o mesmo arquivo.
	if _, err := arquivo.Seek(0, io.SeekStart); err != nil {
		return err
	}
	conteudo, err := io.ReadAll(arquivo)
	if err != nil {
		return err
	}
	fmt.Print(string(conteudo))
	return nil
}

func main() {
	if err := estudarArquivo(); err != nil {
		log.Fatal(err)
	}
}
