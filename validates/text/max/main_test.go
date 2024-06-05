package max

import (
	"testing"

	"github.com/rzjprogramador/Libs_Rzj/Go_Libs/validates/feedback"
)

func Test_MaxText(t *testing.T) {
	numTree := 3
	numtwo := 2
	textOk := "foo"
	sutOK, _ := MaxText(textOk, numTree)
	textFail := "bar"
	_, sutFail := MaxText(textFail, numtwo)
	feedbackErrorMin := feedback.MaxText(numtwo)

	if sutOK != textOk {
		t.Error("Ops...Falhou o test MinText para MaxText")
	}

	if sutFail.Error() != feedbackErrorMin {
		t.Error("Ops...Falhou o test MaxText para Min - Não deu o feedback de Error")
	}
}
