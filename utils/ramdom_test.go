package utils

import (
	"testing"
	"strings"
)

// TestRandomStrig verifica se a função RandomStrig retorna uma string com o comprimento correto e apenas com caracteres permitidos.
func TestRandomStrig(t *testing.T) {
	// Testa a função com diferentes tamanhos de string
	for i := 1; i <= 10; i++ {
		result := RandomStrig(i)
		if len(result) != i {
			t.Errorf("Esperado comprimento %d, mas obteve %d", i, len(result))
		}

		// Verifica se todos os caracteres são do alfabeto permitido
		for _, c := range result {
			if !strings.ContainsRune(alphabet, c) {
				t.Errorf("Caracter %c não está no alfabeto permitido", c)
			}
		}
	}
}

// TestRandomEmail verifica se a função RandomEmail retorna uma string no formato correto de e-mail.
func TestRandomEmail(t *testing.T) {
	result := RandomEmail()

	// Verifica se o e-mail contém o formato correto
	if !strings.HasSuffix(result, "@email.com") {
		t.Errorf("E-mail deve terminar com @email.com, mas obteve %s", result)
	}

	// Verifica se a parte antes de @email.com tem o comprimento correto
	localPart := strings.TrimSuffix(result, "@email.com")
	if len(localPart) != 6 {
		t.Errorf("Parte local do e-mail deve ter comprimento 6, mas obteve %d", len(localPart))
	}

	// Verifica se todos os caracteres são do alfabeto permitido
	for _, c := range localPart {
		if !strings.ContainsRune(alphabet, c) {
			t.Errorf("Caracter %c não está no alfabeto permitido", c)
		}
	}
}
