package model

import (
	"fmt"

	"github.com/spf13/viper"
)

type Metric struct {
	Id          int     `json:"id"`
	IdLxd       int     `json:"id_lxd"`
	CpuUsage    float64 `json:"cpu_usage"`
	MemoryUsage int     `json:"memory_usage"`
	Counter     int     `json:"counter"`
}

func (m *Metric) CreateMetric(ldxId int) (*int, error) {
	err := Db.QueryRow("INSERT INTO metrics(id_lxd) VALUES($1) RETURNING id", ldxId).Scan(&m.Id)

	if err != nil {
		return nil, err
	}

	return &m.Id, nil
}

func (m *Metric) GetMetric(id int) (*Metric, error) {
	queryString := fmt.Sprintf("select * from metrics where id=%d", id)

	row := Db.QueryRow(queryString)

	metric := Metric{}

	if err := row.Scan(&metric.Id, &metric.IdLxd, &metric.CpuUsage, &metric.MemoryUsage, &metric.Counter); err != nil {
		return nil, err
	}

	return &metric, nil
}

func (m *Metric) GetMetrics() ([]Metric, error) {
	rows, err := Db.Query("select * from metrics order by cpu_usage asc, memory_usage asc")

	if err != nil {
		return nil, err
	}

	metrics := []Metric{}

	for rows.Next() {
		var metric Metric
		if err := rows.Scan(&metric.Id, &metric.IdLxd, &metric.CpuUsage, &metric.MemoryUsage, &metric.Counter); err != nil {
			return nil, err
		}

		metrics = append(metrics, metric)
	}

	return metrics, nil
}

func (m *Metric) GetMetricsBelowThreshold() ([]Metric, error) {
	queryString := fmt.Sprintf("SELECT * FROM metrics m where (SELECT COUNT(*) from lxcs where id_lxd=m.id_lxd) < %s order by m.cpu_usage asc, m.memory_usage asc;", viper.GetString("LXC_THRESHOLD"))
	rows, err := Db.Query(queryString)

	if err != nil {
		return nil, err
	}

	metrics := []Metric{}

	for rows.Next() {
		var metric Metric
		if err := rows.Scan(&metric.Id, &metric.IdLxd, &metric.CpuUsage, &metric.MemoryUsage, &metric.Counter); err != nil {
			return nil, err
		}

		metrics = append(metrics, metric)
	}

	return metrics, nil
}

func (m *Metric) GetMetricByLXDId(lxdId int) (*Metric, error) {
	queryString := fmt.Sprintf("select * from metrics where id_lxd=%d", lxdId)

	row := Db.QueryRow(queryString)

	metric := Metric{}

	if err := row.Scan(&metric.Id, &metric.IdLxd, &metric.CpuUsage, &metric.MemoryUsage, &metric.Counter); err != nil {
		return nil, err
	}

	return &metric, nil
}

func (m *Metric) UpdateMetric(metric Metric) error {
	queryString := fmt.Sprintf("Update metrics set cpu_usage='%f', memory_usage='%d', counter='%d' where id_lxd='%d'", metric.CpuUsage, metric.MemoryUsage, metric.Counter, metric.IdLxd)

	_, err := Db.Exec(queryString)

	if err != nil {
		return err
	}

	return nil
}
