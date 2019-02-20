package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	sqsSvc = sqs.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	s3Svc  = s3.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
)

func main() {
	var buff bytes.Buffer
	for {
		input := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(os.Getenv("SQS_URL")),
			WaitTimeSeconds:     aws.Int64(10),
			MaxNumberOfMessages: aws.Int64(10),
		}
		resp, err := sqsSvc.ReceiveMessage(input)
		if err != nil {
			log.Fatalln(err)
		}
		if len(resp.Messages) == 0 {
			break
		}
		for _, v := range resp.Messages {
			var ms MeguroSqs
			if err := json.Unmarshal([]byte(*v.Body), &ms); err != nil {
				log.Fatalln(err)
			}
			meguro := convertMeguro(ms)
			d, err := json.Marshal(meguro)
			if err != nil {
				log.Fatalln(err)
			}
			buff.Write(d)
			buff.WriteString("\n")
			delinput := &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(os.Getenv("SQS_URL")),
				ReceiptHandle: v.ReceiptHandle,
			}
			if _, err := sqsSvc.DeleteMessage(delinput); err != nil {
				log.Fatalln(err)
			}
		}
	}

	param := &s3.PutObjectInput{
		Bucket:       aws.String(os.Getenv("S3_BUCKET")),
		Body:         bytes.NewReader(buff.Bytes()),
		Key:          aws.String("data.json"),
		StorageClass: aws.String(s3.StorageClassStandardIa),
	}
	if _, err := s3Svc.PutObject(param); err != nil {
		log.Fatalln(err)
	}
}

type MeguroSqs struct {
	CompactnessMean struct {
		N string `json:"N"`
	} `json:"compactness_mean"`
	ConcavitySe struct {
		N string `json:"N"`
	} `json:"concavity_se"`
	SmoothnessSe struct {
		N string `json:"N"`
	} `json:"smoothness_se"`
	AreaSe struct {
		N string `json:"N"`
	} `json:"area_se"`
	SymmetryMean struct {
		N string `json:"N"`
	} `json:"symmetry_mean"`
	SymmetrySe struct {
		N string `json:"N"`
	} `json:"symmetry_se"`
	AreaMean struct {
		N string `json:"N"`
	} `json:"area_mean"`
	ConcavePointsMean struct {
		N string `json:"N"`
	} `json:"concave_points_mean"`
	PerimeterSe struct {
		N string `json:"N"`
	} `json:"perimeter_se"`
	PerimeterWorst struct {
		N string `json:"N"`
	} `json:"perimeter_worst"`
	RadiusSe struct {
		N string `json:"N"`
	} `json:"radius_se"`
	SmoothnessWorst struct {
		N string `json:"N"`
	} `json:"smoothness_worst"`
	ConcavePointsWorst struct {
		N string `json:"N"`
	} `json:"concave_points_worst"`
	ConcavityWorst struct {
		N string `json:"N"`
	} `json:"concavity_worst"`
	ID struct {
		N string `json:"N"`
	} `json:"id"`
	ConcavityMean struct {
		N string `json:"N"`
	} `json:"concavity_mean"`
	SmoothnessMean struct {
		N string `json:"N"`
	} `json:"smoothness_mean"`
	PerimeterMean struct {
		N string `json:"N"`
	} `json:"perimeter_mean"`
	CompactnessSe struct {
		N string `json:"N"`
	} `json:"compactness_se"`
	Diagnosis struct {
		S string `json:"S"`
	} `json:"diagnosis"`
	FractalDimensionMean struct {
		N string `json:"N"`
	} `json:"fractal_dimension_mean"`
	CompactnessWorst struct {
		N string `json:"N"`
	} `json:"compactness_worst"`
	ConcavePointsSe struct {
		N string `json:"N"`
	} `json:"concave_points_se"`
	AreaWorst struct {
		N string `json:"N"`
	} `json:"area_worst"`
	RadiusMean struct {
		N string `json:"N"`
	} `json:"radius_mean"`
	TextureSe struct {
		N string `json:"N"`
	} `json:"texture_se"`
	TextureMean struct {
		N string `json:"N"`
	} `json:"texture_mean"`
	RadiusWorst struct {
		N string `json:"N"`
	} `json:"radius_worst"`
	TextureWorst struct {
		N string `json:"N"`
	} `json:"texture_worst"`
	FractalDimensionSe struct {
		N string `json:"N"`
	} `json:"fractal_dimension_se"`
	FractalDimensionWorst struct {
		N string `json:"N"`
	} `json:"fractal_dimension_worst"`
	SymmetryWorst struct {
		N string `json:"N"`
	} `json:"symmetry_worst"`
}

type Meguro struct {
	CompactnessMean       float64 `json:"compactness_mean"`
	ConcavitySe           float64 `json:"concavity_se"`
	SmoothnessSe          float64 `json:"smoothness_se"`
	AreaSe                float64 `json:"area_se"`
	SymmetryMean          float64 `json:"symmetry_mean"`
	SymmetrySe            float64 `json:"symmetry_se"`
	AreaMean              float64 `json:"area_mean"`
	ConcavePointsMean     float64 `json:"concave_points_mean"`
	PerimeterSe           float64 `json:"perimeter_se"`
	PerimeterWorst        float64 `json:"perimeter_worst"`
	RadiusSe              float64 `json:"radius_se"`
	SmoothnessWorst       float64 `json:"smoothness_worst"`
	ConcavePointsWorst    float64 `json:"concave_points_worst"`
	ConcavityWorst        float64 `json:"concavity_worst"`
	ID                    int64   `json:"id"`
	ConcavityMean         float64 `json:"concavity_mean"`
	SmoothnessMean        float64 `json:"smoothness_mean"`
	PerimeterMean         float64 `json:"perimeter_mean"`
	CompactnessSe         float64 `json:"compactness_se"`
	Diagnosis             string  `json:"diagnosis"`
	FractalDimensionMean  float64 `json:"fractal_dimension_mean"`
	CompactnessWorst      float64 `json:"compactness_worst"`
	ConcavePointsSe       float64 `json:"concave_points_se"`
	AreaWorst             float64 `json:"area_worst"`
	RadiusMean            float64 `json:"radius_mean"`
	TextureSe             float64 `json:"texture_se"`
	TextureMean           float64 `json:"texture_mean"`
	RadiusWorst           float64 `json:"radius_worst"`
	TextureWorst          float64 `json:"texture_worst"`
	FractalDimensionSe    float64 `json:"fractal_dimension_se"`
	FractalDimensionWorst float64 `json:"fractal_dimension_worst"`
	SymmetryWorst         float64 `json:"symmetry_worst"`
}

func s2f(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func s2i(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func convertMeguro(ms MeguroSqs) Meguro {
	return Meguro{
		CompactnessMean: s2f(ms.CompactnessMean.N), ConcavitySe: s2f(ms.ConcavitySe.N),
		SmoothnessSe: s2f(ms.SmoothnessSe.N), AreaSe: s2f(ms.AreaSe.N),
		SymmetryMean: s2f(ms.SymmetryMean.N), SymmetrySe: s2f(ms.SymmetrySe.N),
		AreaMean: s2f(ms.AreaMean.N), ConcavePointsMean: s2f(ms.ConcavePointsMean.N),
		PerimeterSe: s2f(ms.PerimeterSe.N), PerimeterWorst: s2f(ms.PerimeterWorst.N),
		RadiusSe: s2f(ms.RadiusSe.N), SmoothnessWorst: s2f(ms.SmoothnessWorst.N),
		ConcavePointsWorst: s2f(ms.ConcavePointsWorst.N), ConcavityWorst: s2f(ms.ConcavityWorst.N),
		ID: s2i(ms.ID.N), ConcavityMean: s2f(ms.ConcavityMean.N),
		SmoothnessMean: s2f(ms.SmoothnessMean.N), PerimeterMean: s2f(ms.PerimeterMean.N),
		CompactnessSe: s2f(ms.CompactnessSe.N), Diagnosis: ms.Diagnosis.S,
		FractalDimensionMean: s2f(ms.FractalDimensionMean.N), CompactnessWorst: s2f(ms.CompactnessWorst.N),
		ConcavePointsSe: s2f(ms.ConcavePointsSe.N), AreaWorst: s2f(ms.AreaWorst.N),
		RadiusMean: s2f(ms.RadiusMean.N), TextureSe: s2f(ms.TextureSe.N),
		TextureMean: s2f(ms.TextureMean.N), RadiusWorst: s2f(ms.RadiusWorst.N),
		TextureWorst: s2f(ms.TextureWorst.N), FractalDimensionSe: s2f(ms.FractalDimensionSe.N),
		FractalDimensionWorst: s2f(ms.FractalDimensionWorst.N), SymmetryWorst: s2f(ms.SymmetryWorst.N),
	}
}
