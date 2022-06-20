package main

import (
	"fmt"
	flesch_go "github.com/NICEXAI/flesch-go"
)

func main() {
	tContent := `A Gram-stain-negative, facultatively aerobic bacterial strain,
designated DB1506(T), of the family Acetobacteraceae, was isolated from
an air-conditioning system in the Republic of Korea. Colonies were pink-
to rosy-coloured and cells were nonmotile cocci with catalase- and
oxidase-positive activities. Growth of strain DB1506(T) was observed at
20-37 degrees C (optimum, 30 degrees C), pH 5.5-8.5 (pH 7.0) and in the
presence of 0-0.5 \\% (w/v) NaCl (0 \\%). Strain DB1506(T) contained
summed feature 8 (comprising C-18(:1)omega 7c and/or C-18:(1)omega 6c)
and C-18(:1) 2-OH as major fatty acids and ubiquinone-10 as the sole
isoprenoid quinone. Phosphatidylglycerol, phosphatidylethanolamine,
unidentified phospholipids, unidentified aminolipids and unidentified
polar lipids were detected as major polar lipids. The G+C content of the
genomic DNA calculated from the whole genome sequence was 72.5 mol\\%.
Strain DB1506(T) was most closely related to Paracraurococcus ruber
NS89(T), Dankookia rubra WS-10(T) and Siccirubricoccus deserti SYSU
D8009(T) with 16S rRNA gene sequence similarities of 96.01, 95.88 and
95.44\\%, respectively, but strain DB1506(T) formed a clearly distinct
phylogenic lineage from them within the family Acetobacteraceae. On the
basis of phenotypic, chemotaxonomic and molecular properties, strain
DB1506(T) represents a novel species of a new genus within the family
Acetobacteraceae, for which the name Roseicella frigidaeris gen nov., sp
nov. is proposed. The type strain is DB1506(T) (=KACC 19791(T)=JCM
32945(T)).`

	document, err := flesch_go.ParseString(tContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get score
	fmt.Println(document.Score())

	// get kincaid
	fmt.Println(document.Kincaid())

	fmt.Println(document.Syllables(), document.SentencesCount())

	// get word count
	fmt.Println(document.WordCount())

	for _, sentence := range document.Sentences {
		fmt.Println("***************************************")
		fmt.Println(sentence)
	}
}
