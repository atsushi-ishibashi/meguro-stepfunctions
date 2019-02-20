package main

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	ddbSvc = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
)

func main() {
	f, err := os.Open("./wdbc.data")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(f)
	reader.Comma = '\t'
	reader.LazyQuotes = true
	count := 0
	for count < 100 {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		rec := record[0]
		recs := strings.Split(rec, ",")
		input := &dynamodb.PutItemInput{
			Item: map[string]*dynamodb.AttributeValue{
				"id":                      {N: aws.String(recs[0])},
				"diagnosis":               {S: aws.String(recs[1])},
				"radius_mean":             {N: aws.String(recs[2])},
				"texture_mean":            {N: aws.String(recs[3])},
				"perimeter_mean":          {N: aws.String(recs[4])},
				"area_mean":               {N: aws.String(recs[5])},
				"smoothness_mean":         {N: aws.String(recs[6])},
				"compactness_mean":        {N: aws.String(recs[7])},
				"concavity_mean":          {N: aws.String(recs[8])},
				"concave_points_mean":     {N: aws.String(recs[9])},
				"symmetry_mean":           {N: aws.String(recs[10])},
				"fractal_dimension_mean":  {N: aws.String(recs[11])},
				"radius_se":               {N: aws.String(recs[12])},
				"texture_se":              {N: aws.String(recs[13])},
				"perimeter_se":            {N: aws.String(recs[14])},
				"area_se":                 {N: aws.String(recs[15])},
				"smoothness_se":           {N: aws.String(recs[16])},
				"compactness_se":          {N: aws.String(recs[17])},
				"concavity_se":            {N: aws.String(recs[18])},
				"concave_points_se":       {N: aws.String(recs[19])},
				"symmetry_se":             {N: aws.String(recs[20])},
				"fractal_dimension_se":    {N: aws.String(recs[21])},
				"radius_worst":            {N: aws.String(recs[22])},
				"texture_worst":           {N: aws.String(recs[23])},
				"perimeter_worst":         {N: aws.String(recs[24])},
				"area_worst":              {N: aws.String(recs[25])},
				"smoothness_worst":        {N: aws.String(recs[26])},
				"compactness_worst":       {N: aws.String(recs[27])},
				"concavity_worst":         {N: aws.String(recs[28])},
				"concave_points_worst":    {N: aws.String(recs[29])},
				"symmetry_worst":          {N: aws.String(recs[30])},
				"fractal_dimension_worst": {N: aws.String(recs[31])},
			},
			TableName: aws.String("<table name>"),
		}

		_, err = ddbSvc.PutItem(input)
		count++
	}
}
