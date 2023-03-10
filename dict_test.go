package uni

import (
	"testing"

	"github.com/ikawaha/kagome-dict/dict"
)

const (
	UniDictEntrySize = 879222
	testDictPath     = "./uni.dict"
)

func Test_DictShrink(t *testing.T) {
	d := DictShrink()
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := 0, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test_LoadShrink(t *testing.T) {
	d, err := dict.LoadShrink(testDictPath)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := 0, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test_New(t *testing.T) {
	d := Dict()
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := UniDictEntrySize, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test_LoadDictFile(t *testing.T) {
	d, err := dict.LoadDictFile(testDictPath)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if want, got := UniDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := UniDictEntrySize, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test_Singleton(t *testing.T) {
	a := Dict()
	b := Dict()
	if a != b {
		t.Errorf("got %p and %p, expected singleton", a, b)
	}
}

func Test_ContentsMeta(t *testing.T) {
	d := Dict()
	if want, got := 0, d.ContentsMeta[dict.POSStartIndex]; want != int(got) {
		t.Errorf("pos start index: want %d, got %d", want, got)
	}
	if want, got := 4, d.ContentsMeta[dict.POSHierarchy]; want != int(got) {
		t.Errorf("pos end index: want %d, got %d", want, got)
	}
	if want, got := 4, d.ContentsMeta[dict.InflectionalType]; want != int(got) {
		t.Errorf("base form index: want %d, got %d", want, got)
	}
	if want, got := 5, d.ContentsMeta[dict.InflectionalForm]; want != int(got) {
		t.Errorf("base form index: want %d, got %d", want, got)
	}
	if want, got := 0, d.ContentsMeta[dict.ReadingIndex]; want != int(got) { // undefined
		t.Errorf("reading index: want %d, got %d", want, got)
	}
	if want, got := 10, d.ContentsMeta[dict.BaseFormIndex]; want != int(got) {
		t.Errorf("base form index: want %d, got %d", want, got)
	}
	if want, got := 9, d.ContentsMeta[dict.PronunciationIndex]; want != int(got) {
		t.Errorf("pronunciation index: want %d, got %d", want, got)
	}
}

/*
func Test_InflectionalType(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("?????????????????????")
	if want, got := 6, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].InflectionalType(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//?????? ??????,??????,*,*,?????????-??????,?????????-??????,?????????,?????????,??????,??????,?????????,?????????,???,*,*,*,*
	if got, ok := tokens[3].InflectionalType(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "?????????-??????"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func Test_InflectionalForm(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("?????????????????????")
	if want, got := 6, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].InflectionalForm(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//?????? ??????,??????,*,*,?????????-??????,?????????-??????,?????????,?????????,??????,??????,?????????,?????????,???,*,*,*,*
	if got, ok := tokens[3].InflectionalForm(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "?????????-??????"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func Test_BaseForm(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("?????????????????????")
	if want, got := 6, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].BaseForm(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//??????->?????????
	if got, ok := tokens[3].BaseForm(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "?????????"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func Test_Reading(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("???????????????")
	if want, got := 5, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].Reading(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//??????->???????????? // undefined
	if got, ok := tokens[1].Reading(); ok {
		t.Errorf("want !ok, but ok, got %s", got)
	}
}

func Test_Pronunciation(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("???????????????")
	if want, got := 5, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].Pronunciation(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//??????->????????????
	if got, ok := tokens[1].Pronunciation(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "????????????"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

func Test_POS(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("???????????????")
	if want, got := 5, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got := tokens[0].POS(); len(got) > 0 {
		t.Errorf("want empty, got %+v", got)
	}
	//??????
	if want, got := []string{"??????", "???????????????", "*", "*"}, tokens[3].POS(); !reflect.DeepEqual(want, got) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func Test_FeatureIndex(t *testing.T) {
	testdata := []struct {
		index   FeatureIndex
		feature string
	}{
		{index: CType, feature: "?????????-??????"},
		{index: CForm, feature: "?????????-??????"},
		{index: LForm, feature: "?????????"},
		{index: Lemma, feature: "?????????"},
		{index: Orth, feature: "??????"},
		{index: Pron, feature: "??????"},
		{index: OrthBase, feature: "?????????"},
		{index: PronBase, feature: "?????????"},
		{index: Goshu, feature: "???"},
		{index: IType, feature: "*"},
		{index: IForm, feature: "*"},
		{index: FType, feature: "*"},
		{index: FForm, feature: "*"},
		// aliases
		{index: InflectionalType, feature: "?????????-??????"},
		{index: InflectionalForm, feature: "?????????-??????"},
		{index: BaseForm, feature: "?????????"},
		{index: Pronunciation, feature: "??????"},
	}
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}

	tokens := tnz.Tokenize("?????????????????????")
	if want, got := 6, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//?????? ??????,??????,*,*,?????????-??????,?????????-??????,?????????,?????????,??????,??????,?????????,?????????,???,*,*,*,*
	token := tokens[3]
	for _, v := range testdata {
		if got, ok := token.FeatureAt(v.index); !ok {
			t.Errorf("index %v: want ok, but !ok", v.index)
		} else if want := v.feature; want != got {
			t.Fatalf("index %v: want %s, got %s", v.index, want, got)
		}
	}
}
*/
