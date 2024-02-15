package view

import (
	"math"
	"time"

	"github.com/pterm/pterm"
)

type Context struct {
	Cmd    string
	Target string
	From   string

	Protocol  string
	Port      int
	Resolver  string
	Trace     bool
	QueryType string

	Limit   int // Number of probes to use
	Packets int // Number of packets to send

	ToJSON    bool // Determines whether the output should be in JSON format.
	ToLatency bool // Determines whether the output should be only the stats of a measurement
	CI        bool // Determine whether the output should be in a format that is easy to parse by a CI tool
	Full      bool // Full output
	Share     bool // Display share message
	Infinite  bool // Infinite flag

	APIMinInterval time.Duration // Minimum interval between API calls

	Area            *pterm.AreaPrinter
	Hostname        string
	MStartedAt      time.Time // Time when the measurement started
	CompletedStats  []MeasurementStats
	InProgressStats []MeasurementStats
	CallCount       int      // Number of measurements created
	MaxHistory      int      // Maximum number of measurements to keep in history
	History         *Rbuffer // History of measurements
}

type HTTPOpts struct {
	Path     string
	Query    string
	Host     string
	Method   string
	Protocol string
	Port     int
	Resolver string
	Headers  []string
}

type MeasurementStats struct {
	Sent  int     // Number of packets sent
	Rcv   int     // Number of packets received
	Lost  int     // Number of packets lost
	Loss  float64 // Percentage of packets lost
	Last  float64 // Last RTT
	Min   float64 // Minimum RTT
	Avg   float64 // Average RTT
	Max   float64 // Maximum RTT
	Mdev  float64 // Mean deviation of RTT
	Time  float64 // Total time of measurement, in milliseconds
	Tsum  float64 // Total sum of RTT
	Tsum2 float64 // Total sum of RTT squared
}

func NewMeasurementStats() MeasurementStats {
	return MeasurementStats{Last: -1, Min: math.MaxFloat64, Avg: -1, Max: -1}
}

type Rbuffer struct {
	Index int
	Slice []string
}

func (q *Rbuffer) Push(id string) {
	q.Slice[q.Index] = id
	q.Index = (q.Index + 1) % len(q.Slice)
}

func (q *Rbuffer) ToString(sep string) string {
	s := ""
	i := q.Index
	isFirst := true
	for {
		if q.Slice[i] != "" {
			if isFirst {
				isFirst = false
				s += q.Slice[i]
			} else {
				s += sep + q.Slice[i]
			}
		}
		i = (i + 1) % len(q.Slice)
		if i == q.Index {
			break
		}
	}
	return s
}

func NewRbuffer(size int) *Rbuffer {
	return &Rbuffer{
		Index: 0,
		Slice: make([]string, size),
	}
}