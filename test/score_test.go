package test

import (
	"github.com/qinsheng99/go-py/sdk"
	"testing"
)

var client sdk.CalculateEvaluate

func TestEvaluate(t *testing.T) {
	err, res := client.Evaluate(&sdk.Calculate{
		PredPath: "xihe-obj/competitions/昇思AI挑战赛-多类别图像分类/submit_result/s9qfqri3zpc8j2x7_1/result_example_5120-2022-8-8-15-3-16.txt",
		TruePath: "xihe-obj/competitions/昇思AI挑战赛-多类别图像分类/result/label.txt",
		Cls:      256,
		Pos:      1,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}

func TestCalculate(t *testing.T) {
	err := client.Calculate(&sdk.Calculate{
		UserResult: "xihe-obj/competitions/昇思AI挑战赛-艺术家画作风格迁移/submit_result/victor_1/result",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestMain(t *testing.M) {
	client = sdk.NewCalculateEvaluate("http://192.168.1.218:8000")
}
