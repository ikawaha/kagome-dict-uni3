Uni3: A Dictionary of Kagome v2
===

A dictionary package of [kagome v2](http://github.com/ikawaha/kagome/tree/v2). 
This software includes a binary and/or source version of data from

* unidic-cwj-3.1.1-full

which can be obtained from

* https://clrd.ninjal.ac.jp/unidic_archive/cwj/3.1.1/unidic-cwj-3.1.1.zip

 # Feature Fields
 
Features are information given to a word, such as follows:

```
公園	名詞,普通名詞,一般,*,*,*,コウエン,公園,公園,コーエン,公園,コーエン,漢,*,*,*,*,*,*,体,コウエン,コウエン,コウエン,コウエン,0,C2,*,3269956204634624,11896
に	助詞,格助詞,*,*,*,*,ニ,に,に,ニ,に,ニ,和,*,*,*,*,*,*,格助,ニ,ニ,ニ,ニ,*,名詞%F1,*,7745518285496832,28178
行っ	動詞,非自立可能,*,*,五段-カ行,連用形-促音便,イク,行く,行っ,イッ,行く,イク,和,*,*,*,*,*,*,用,イッ,イク,イッ,イク,0,C2,*,470874511778436,1713
た	助動詞,*,*,*,助動詞-タ,終止形-一般,タ,た,た,タ,た,タ,和,*,*,*,*,*,*,助動,タ,タ,タ,タ,*,動詞%F2@1,形容詞%F4@-2,*,5948916285711019,21642
```
 
|No.|feature| name |description|
|:----|:---|:---|:---|
| 0   | POS hierarchy | 品詞大分類 |The POS name and each level in its hierarchical structure. |
| 1   | POS hierarchy 1 | 品詞中分類 | |
| 2   | POS hierarchy 2 | 品詞小分類 |     |
| 3   | POS hierarchy 3 | 品詞細分類|     |
| 4   | cType | 活用型 | Inflection type indicates a category that is an inflected form, e.g. 五段-カ行. |
| 5   | cForm | 活用形 | Inflected form, e.g. 連用形-促音便. |
| 6   | lForm | 語彙素読み | A possible reading for a lemma entry. Readings are given in katakana, e.g. コウエン, イク. |
| 7   | lemma | 語彙素表記 | A lemma, e.g. 公園, 行く.|
| 8   | orth | 書字形| A conjugation form for an orthBase. e.g. 行く -> 行っ. |
| 9   | pron | 発音形| A possible pronunciations for an orthToken. Pronunciations are given in katakana, .e.g. コーエン, イッ. |
| 10  | orthBase | 書字形基本形|  A form of dictionary headword, e.g. 行っ -> 行く.  |
| 11  | pronBase | 発音形基本形 | A possible pronunciations for an entry headword. Pronunciations are given in katakana, .e.g. イク. | 
| 12  | goshu | 語種| A possible pronunciations for an entry. Pronunciations are given in katakana, .e.g. コーエン. |
| 13  | iType | 語頭変化型 | |
| 14  | iForm | 語頭変化形 | |
| 15  | fType | 語末変化型| |
| 16  | fForm | 語末変化型| |
| 17  | iConType| | |
| 18  | fConType| | |
| 19  | type| | |
| 20  | kana| | |
| 21  | kanaBase| | |
| 22  | form| | |
| 23  | formBase| | |
| 24  | aType| | |
| 25  | aConType| | |
| 26  | aModType| | |
| 27  | lid| | |
| 28  | lemma_id| | |
 
# Licence
  
MIT