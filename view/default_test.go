package view

import (
	"io"
	"os"
	"testing"

	"github.com/jsdelivr/globalping-cli/globalping"
	"github.com/jsdelivr/globalping-cli/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Output_Default_HTTP_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurement := &globalping.Measurement{
		Results: []globalping.ProbeMeasurement{
			{
				Probe: globalping.ProbeDetails{
					Continent: "EU",
					Country:   "DE",
					City:      "Berlin",
					ASN:       123,
					Network:   "Network 1",
				},
				Result: globalping.ProbeResult{
					RawOutput:  "Headers 1\nBody 1",
					RawHeaders: "Headers 1",
					RawBody:    "Body 1",
				},
			},

			{
				Probe: globalping.ProbeDetails{
					Continent: "NA",
					Country:   "US",
					City:      "New York",
					State:     "NY",
					ASN:       567,
					Network:   "Network 2",
				},
				Result: globalping.ProbeResult{
					RawOutput:  "Headers 2\nBody 2",
					RawHeaders: "Headers 2",
					RawBody:    "Body 2",
				},
			},
		},
	}

	gbMock := mocks.NewMockClient(ctrl)
	gbMock.EXPECT().GetMeasurement(measurementID1).Times(1).Return(measurement, nil)

	r, w, err := os.Pipe()
	assert.NoError(t, err)
	defer r.Close()
	defer w.Close()

	m := &globalping.MeasurementCreate{
		Options: &globalping.MeasurementOptions{
			Request: &globalping.RequestOptions{
				Method: "GET",
			},
		},
	}

	viewer := NewViewer(&Context{
		Cmd: "http",
		CI:  true,
	}, NewPrinter(w), nil, gbMock)

	viewer.Output(measurementID1, m)
	w.Close()

	outContent, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Equal(t, `> EU, DE, Berlin, ASN:123, Network 1
Body 1

> NA, US, (NY), New York, ASN:567, Network 2
Body 2
`, string(outContent))
}

func Test_Output_Default_HTTP_Get_Share(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurement := &globalping.Measurement{
		Results: []globalping.ProbeMeasurement{
			{
				Probe: globalping.ProbeDetails{
					Continent: "EU",
					Country:   "DE",
					City:      "Berlin",
					ASN:       123,
					Network:   "Network 1",
				},
				Result: globalping.ProbeResult{
					RawOutput:  "Headers 1\nBody 1",
					RawHeaders: "Headers 1",
					RawBody:    "Body 1",
				},
			},

			{
				Probe: globalping.ProbeDetails{
					Continent: "NA",
					Country:   "US",
					City:      "New York",
					State:     "NY",
					ASN:       567,
					Network:   "Network 2",
				},
				Result: globalping.ProbeResult{
					RawOutput:  "Headers 2\nBody 2",
					RawHeaders: "Headers 2",
					RawBody:    "Body 2",
				},
			},
		},
	}

	gbMock := mocks.NewMockClient(ctrl)
	gbMock.EXPECT().GetMeasurement(measurementID1).Times(1).Return(measurement, nil)

	r, w, err := os.Pipe()
	assert.NoError(t, err)
	defer r.Close()
	defer w.Close()

	m := &globalping.MeasurementCreate{
		Options: &globalping.MeasurementOptions{
			Request: &globalping.RequestOptions{
				Method: "GET",
			},
		},
	}

	viewer := NewViewer(&Context{
		Cmd:   "http",
		CI:    true,
		Share: true,
	}, NewPrinter(w), nil, gbMock)

	viewer.Output(measurementID1, m)
	w.Close()

	outContent, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Equal(t, `> EU, DE, Berlin, ASN:123, Network 1
Body 1

> NA, US, (NY), New York, ASN:567, Network 2
Body 2
> View the results online: https://www.jsdelivr.com/globalping?measurement=nzGzfAGL7sZfUs3c
`, string(outContent))
}

func Test_Output_Default_HTTP_Get_Full(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurement := &globalping.Measurement{
		Results: []globalping.ProbeMeasurement{
			{
				Probe: globalping.ProbeDetails{
					Continent: "EU",
					Country:   "DE",
					City:      "Berlin",
					ASN:       123,
					Network:   "Network 1",
				},
				Result: globalping.ProbeResult{
					RawOutput:  "Headers 1\nBody 1",
					RawHeaders: "Headers 1",
					RawBody:    "Body 1",
				},
			},

			{
				Probe: globalping.ProbeDetails{
					Continent: "NA",
					Country:   "US",
					City:      "New York",
					State:     "NY",
					ASN:       567,
					Network:   "Network 2",
				},
				Result: globalping.ProbeResult{
					RawOutput:  "Headers 2\nBody 2",
					RawHeaders: "Headers 2",
					RawBody:    "Body 2",
				},
			},
		},
	}

	gbMock := mocks.NewMockClient(ctrl)
	gbMock.EXPECT().GetMeasurement(measurementID1).Times(1).Return(measurement, nil)

	r, w, err := os.Pipe()
	assert.NoError(t, err)
	defer r.Close()
	defer w.Close()

	m := &globalping.MeasurementCreate{
		Options: &globalping.MeasurementOptions{
			Request: &globalping.RequestOptions{
				Method: "GET",
			},
		},
	}

	viewer := NewViewer(&Context{
		Cmd:  "http",
		CI:   true,
		Full: true,
	}, NewPrinter(w), nil, gbMock)

	viewer.Output(measurementID1, m)
	w.Close()

	outContent, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Equal(t, `> EU, DE, Berlin, ASN:123, Network 1
Headers 1
Body 1

> NA, US, (NY), New York, ASN:567, Network 2
Headers 2
Body 2
`, string(outContent))
}

func Test_Output_Default_HTTP_Head(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurement := &globalping.Measurement{
		Results: []globalping.ProbeMeasurement{
			{
				Probe: globalping.ProbeDetails{
					Continent: "EU",
					Country:   "DE",
					City:      "Berlin",
					ASN:       123,
					Network:   "Network 1",
				},
				Result: globalping.ProbeResult{
					RawOutput:  "Headers 1",
					RawHeaders: "Headers 1",
				},
			},

			{
				Probe: globalping.ProbeDetails{
					Continent: "NA",
					Country:   "US",
					City:      "New York",
					State:     "NY",
					ASN:       567,
					Network:   "Network 2",
				},
				Result: globalping.ProbeResult{
					RawOutput:  "Headers 2",
					RawHeaders: "Headers 2",
				},
			},
		},
	}

	gbMock := mocks.NewMockClient(ctrl)
	gbMock.EXPECT().GetMeasurement(measurementID1).Times(1).Return(measurement, nil)

	r, w, err := os.Pipe()
	assert.NoError(t, err)
	defer r.Close()
	defer w.Close()

	m := &globalping.MeasurementCreate{
		Options: &globalping.MeasurementOptions{
			Request: &globalping.RequestOptions{
				Method: "HEAD",
			},
		},
	}

	viewer := NewViewer(&Context{
		Cmd: "http",
		CI:  true,
	}, NewPrinter(w), nil, gbMock)

	viewer.Output(measurementID1, m)
	w.Close()

	outContent, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Equal(t, `> EU, DE, Berlin, ASN:123, Network 1
Headers 1

> NA, US, (NY), New York, ASN:567, Network 2
Headers 2
`, string(outContent))
}

func Test_Output_Default_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurement := &globalping.Measurement{
		Results: []globalping.ProbeMeasurement{
			{
				Probe: globalping.ProbeDetails{
					Continent: "EU",
					Country:   "DE",
					City:      "Berlin",
					ASN:       123,
					Network:   "Network 1",
				},
				Result: globalping.ProbeResult{
					RawOutput: "Ping Results 1",
				},
			},

			{
				Probe: globalping.ProbeDetails{
					Continent: "NA",
					Country:   "US",
					City:      "New York",
					State:     "NY",
					ASN:       567,
					Network:   "Network 2",
				},
				Result: globalping.ProbeResult{
					RawOutput: "Ping Results 2",
				},
			},
		},
	}

	gbMock := mocks.NewMockClient(ctrl)
	gbMock.EXPECT().GetMeasurement(measurementID1).Times(1).Return(measurement, nil)

	r, w, err := os.Pipe()
	assert.NoError(t, err)
	defer r.Close()
	defer w.Close()

	m := &globalping.MeasurementCreate{}

	viewer := NewViewer(&Context{
		Cmd: "ping",
		CI:  true,
	}, NewPrinter(w), nil, gbMock)

	viewer.Output(measurementID1, m)
	w.Close()

	outContent, err := io.ReadAll(r)
	assert.NoError(t, err)
	assert.Equal(t, `> EU, DE, Berlin, ASN:123, Network 1
Ping Results 1

> NA, US, (NY), New York, ASN:567, Network 2
Ping Results 2
`, string(outContent))
}