package main

import (
	"fmt"
	"os"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/csv"
	// "github.com/apache/arrow/go/v15/arrow/table"
)

type Measurement struct {
	Station string
	Temp    float64
}

type Stats struct {
	Min, Mean, Max float64
}

// func main() {

// 	ch := make(chan arrow.Record, 20)
// 	go func() {
// 		defer close(ch)
// 		file, err := os.Open("data/weather_stations.csv")
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer file.Close()

// 		// Infer the types and schema from the header line and first line of data
// 		// rdr := csv.NewInferringReader(
// 		// 	file,
// 		// 	csv.WithComma(';'),
// 		// 	csv.WithChunk(5000),
// 		// 	csv.WithHeader(false),
// 		// )

// 		schema := arrow.NewSchema(
// 			[]arrow.Field{
// 				{Name: "station", Type: arrow.BinaryTypes.String},
// 				{Name: "temp", Type: arrow.PrimitiveTypes.Float64},
// 			}, nil,
// 		)

// 		rdr := csv.NewReader(file, schema, csv.WithComma(';'), csv.WithChunk(5000), csv.WithHeader(false))

// 		for rdr.Next() {
// 			rec := rdr.Record()
// 			// print and format the record
// 			fmt.Println(rec)
// 			rec.Retain()
// 			ch <- rec
// 		}

// 		if rdr.Err() != nil {
// 			panic(rdr.Err())
// 		}
// 	}()

// }

func main() {

	file, err := os.Open("data/weather_stations.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "station", Type: arrow.BinaryTypes.String},
			{Name: "temp", Type: arrow.PrimitiveTypes.Float64},
		}, nil,
	)

	rdr := csv.NewReader(
		file,
		schema,
		csv.WithComma(';'),
		csv.WithChunk(5000),
		csv.WithHeader(false),
	)

	for rdr.Next() {
		rec := rdr.Record()
		fmt.Println(rec)
		rec.Retain()
	}

	// var tbl arrow.Table
	// for rdr.Next() {
	// 	rec := rdr.Record()
	// 	if tbl == nil {
	// 		tbl = rec.NewTable()
	// 	} else {
	// 		tbl.Append(rec)
	// 	}
	// 	rec.Release()
	// }

	// if err := rdr.Err(); err != nil {
	// 	panic(err)
	// }

	// groupedTbl := compute.TableGroupBy(tbl, []compute.Aggregate{
	// 	{
	// 		Name:     "avg_temp",
	// 		Function: "hash_mean",
	// 		Target: &compute.AggregateTarget{
	// 			Column: "temp",
	// 		},
	// 	},
	// }, []string{"station"})

	// defer groupedTbl.Release()

	// fmt.Println(groupedTbl)
}
