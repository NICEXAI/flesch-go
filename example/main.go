package main

import (
	"fmt"
	flesch_go "github.com/NICEXAI/flesch-go"
)

func main() {
	tContent := `Axillary lymph node metastasis (ALNM) is commonly the earliest detectable clinical manifestation of breast cancer when distant metastasis emerges. This study aimed to explore the influencing factors of ALNM and develop models that can predict its occurrence preoperatively. Cases of sonographically visible clinical stage T1-2N0M0 breast cancers treated with breast and axillary surgery at West China Hospital were retrospectively reviewed. Univariate and multivariate logistic regression analyses were performed to evaluate associations between ALNM and variables. Decision tree analyses were performed to construct predictive models using the C5.0 packages. Of the 1671 tumors, 541 (32.9%) showed axillary lymph node positivity on final surgical histopathologic analysis. In multivariate logistic regression analysis, tumor size (P<.001), infiltration of subcutaneous adipose tissue (P<.001), infiltration of the interstitial adipose tissue (P=.031), and tumor quadrant locations (P<.001) were significantly correlated with ALNM. Furthermore, the accuracy in the decision tree model was 69.52%, and the false-negative rate (FNR) was 74.18%. By using the error-cost matrix algorithm, the FNR significantly decreased to 14.75%, particularly for nodes 5, 8, and 13 (FNR: 11.4%, 9.09%, and 14.29% in the training set and 18.1%,14.71%, and 20% in the test set, respectively). In summary, our study demonstrated that tumor lesion boundary, tumor size, and tumor quadrant locations were the most important factors affecting ALNM in cT1-2N0M0 stage breast cancer. The decision tree built using these variables reached a slightly higher FNR than sentinel lymph node dissection in predicting ALNM in some selected patients.`

	document, err := flesch_go.ParseString(tContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get score
	fmt.Println(document.Score())

	// get word count
	fmt.Println(document.WordCount())
}
