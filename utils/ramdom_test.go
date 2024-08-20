package utils

import (
	"strings"
	"testing"
)


func TestRandomStrig(t *testing.T) {
	
	for i := 1; i <= 10; i++ {
		result := RandomStrig(i)
		if len(result) != i {
			t.Errorf("Esperado comprimento %d, mas obteve %d", i, len(result))
		}

	
		for _, c := range result {
			if !strings.ContainsRune(alphabet, c) {
				t.Errorf("Caracter %c não está no alfabeto permitido", c)
			}
		}
	}
}


func TestRandomEmail(t *testing.T) {
	result := RandomEmail()

	
	if !strings.HasSuffix(result, "@email.com") {
		t.Errorf("E-mail deve terminar com @email.com, mas obteve %s", result)
	}

	
	localPart := strings.TrimSuffix(result, "@email.com")
	if len(localPart) != 6 {
		t.Errorf("Parte local do e-mail deve ter comprimento 6, mas obteve %d", len(localPart))
	}

	
	for _, c := range localPart {
		if !strings.ContainsRune(alphabet, c) {
			t.Errorf("Caracter %c não está no alfabeto permitido", c)
		}
	}
}
