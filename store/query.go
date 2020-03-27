package store

import (
	"context"

	"github.com/influxdata/influxdb-client-go"
	"github.com/sirupsen/logrus"
)

// InsertData insert data
func (i *Influx) InsertData(metrics *[]influxdb.Metric, bucket string, org string) (count int, err error) {
	count, err = i.GetDB().Write(context.Background(), bucket, org, *metrics...)
	if err != nil {
		logrus.Fatal(err)
	}

	return
}

// QueryData query data
func (i *Influx) QueryData(queryString string, org string) (res *influxdb.QueryCSVResult, err error) {

	res, err = i.GetDB().QueryCSV(context.Background(),
		queryString,
		org)
	if err != nil {
		logrus.Error(err)
		return
	}

	return
}
