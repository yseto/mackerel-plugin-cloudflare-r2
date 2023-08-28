package main

// go:generate go run github.com/Khan/genqlient

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/Khan/genqlient/generate"
	"github.com/Khan/genqlient/graphql"
	mp "github.com/mackerelio/go-mackerel-plugin"
)

type AddHeaderTransport struct {
	T     http.RoundTripper
	email string
	key   string
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-AUTH-EMAIL", adt.email)
	req.Header.Add("X-AUTH-KEY", adt.key)

	return adt.T.RoundTrip(req)
}

const defaultPrefix = "r2"

type R2Plugin struct {
	prefix     string
	accountTag string
	email      string
	key        string
	t0         time.Time
}

func (m *R2Plugin) MetricKeyPrefix() string {
	if m.prefix == "" {
		m.prefix = defaultPrefix
	}
	return m.prefix
}

func (m R2Plugin) FetchMetrics() (map[string]float64, error) {
	ctx := context.Background()

	httpClient := http.DefaultClient
	httpClient.Transport = &AddHeaderTransport{
		T:     http.DefaultTransport,
		email: m.email,
		key:   m.key,
	}

	client := graphql.NewClient("https://api.cloudflare.com/client/v4/graphql", httpClient)

	t1 := m.t0.Add(59 * time.Second)

	resp, err := GetR2Metrics(ctx, client, m.accountTag, m.t0.Format(time.RFC3339), t1.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	v := make(map[string]float64, 0)

	for _, a := range resp.Viewer.Accounts {
		for _, c := range a.ClassA {
			v["Requests.classA."+c.Dimensions.ActionType] = float64(c.Sum.Requests)
			v["ResponseObjectSize.classA."+c.Dimensions.ActionType] = float64(c.Sum.ResponseObjectSize)
		}

		for _, c := range a.ClassB {
			v["Requests.classB."+c.Dimensions.ActionType] = float64(c.Sum.Requests)
			v["ResponseObjectSize.classB."+c.Dimensions.ActionType] = float64(c.Sum.ResponseObjectSize)
		}

		for _, b := range a.ClassAperBucket {
			var s = b.Dimensions.BucketName
			if s == "" {
				s = "_wildcard_"
			}
			v["bucket.Aops."+s] = float64(b.Sum.Requests)
		}

		for _, b := range a.ClassBperBucket {
			var s = b.Dimensions.BucketName
			if s == "" {
				s = "_wildcard_"
			}
			v["bucket.Bops."+s] = float64(b.Sum.Requests)
		}
	}

	t2 := m.t0.Add(-24 * time.Hour)
	// fmt.Println(t2.Format(time.RFC3339))
	// fmt.Println(t1.Format(time.RFC3339))
	// fmt.Println(m.t0.Format(time.RFC3339))

	// Widen the acquisition window for stabilize
	resp2, err := GetR2MetricsStorage(ctx, client, m.accountTag, t2.Format(time.RFC3339), t1.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}
	for _, a := range resp2.Viewer.Accounts {
		for _, s := range a.Storage {
			v["Storage.MetadataSize."+s.Dimensions.BucketName] = float64(s.Max.MetadataSize)
			v["Storage.ObjectCount."+s.Dimensions.BucketName] = float64(s.Max.ObjectCount)
			v["Storage.PayloadSize."+s.Dimensions.BucketName] = float64(s.Max.PayloadSize)
			v["Storage.UploadCount."+s.Dimensions.BucketName] = float64(s.Max.UploadCount)
		}
	}

	return v, nil
}

func (m R2Plugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Replace(m.MetricKeyPrefix(), defaultPrefix, "R2", -1)

	return map[string]mp.Graphs{
		"ResponseObjectSize.classA": {
			Label:   labelPrefix + " ResponseObjectSize (classA)",
			Unit:    mp.UnitInteger,
			Metrics: []mp.Metrics{{Name: "*", Label: "%1"}},
			// %1 is to be like ListObjects, PutObject
			// `*` is value of set a postfix.
		},
		"ResponseObjectSize.classB": {
			Label:   labelPrefix + " ResponseObjectSize (classB)",
			Unit:    mp.UnitInteger,
			Metrics: []mp.Metrics{{Name: "*", Label: "%1"}},
		},
		"Requests.classA": {
			Label:   labelPrefix + " Requests (classA)",
			Unit:    mp.UnitInteger,
			Metrics: []mp.Metrics{{Name: "*", Label: "%1"}},
		},
		"Requests.classB": {
			Label:   labelPrefix + " Requests (classB)",
			Unit:    mp.UnitInteger,
			Metrics: []mp.Metrics{{Name: "*", Label: "%1"}},
		},
		"Storage.MetadataSize": {
			Label:   labelPrefix + " MetadataSize",
			Unit:    mp.UnitInteger,
			Metrics: []mp.Metrics{{Name: "*", Label: "%1"}},
		},
		"Storage.ObjectCount": {
			Label:   labelPrefix + " ObjectCount",
			Unit:    mp.UnitInteger,
			Metrics: []mp.Metrics{{Name: "*", Label: "%1"}},
		},
		"Storage.PayloadSize": {
			Label:   labelPrefix + " PayloadSize",
			Unit:    mp.UnitInteger,
			Metrics: []mp.Metrics{{Name: "*", Label: "%1"}},
		},
		"Storage.UploadCount": {
			Label:   labelPrefix + " UploadCount",
			Unit:    mp.UnitInteger,
			Metrics: []mp.Metrics{{Name: "*", Label: "%1"}},
		},
		"bucket.Aops": {
			Label:   labelPrefix + " classA operation per Buckets",
			Unit:    mp.UnitInteger,
			Metrics: []mp.Metrics{{Name: "*", Label: "%1"}},
		},
		"bucket.Bops": {
			Label:   labelPrefix + " classB operation per Buckets",
			Unit:    mp.UnitInteger,
			Metrics: []mp.Metrics{{Name: "*", Label: "%1"}},
		},
	}
}

func main() {
	optPrefix := flag.String("prefix", "", "prefix default:"+defaultPrefix)
	optAccountTag := flag.String("account-id", "", "account ID")
	optEmail := flag.String("email", "", "email address")
	optKey := flag.String("global-api-key", "", "Global API Key")

	optTempfile := flag.String("tempfile", "", "Temp file name")
	flag.Parse()

	if *optAccountTag == "" || *optEmail == "" || *optKey == "" {
		flag.Usage()
		os.Exit(1)
	}

	// this value was determined by performing actual operation.
	offset := -3
	t0 := time.Now().Truncate(time.Minute).Add(time.Duration(offset) * time.Minute)

	r2 := R2Plugin{
		prefix:     *optPrefix,
		accountTag: *optAccountTag,
		email:      *optEmail,
		key:        *optKey,
		t0:         t0,
	}

	helper := mp.NewMackerelPlugin(r2)
	helper.Tempfile = *optTempfile
	// helper.Run()

	if os.Getenv("MACKEREL_AGENT_PLUGIN_META") != "" {
		helper.OutputDefinitions()
		return
	}
	stat, err := r2.FetchMetrics()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// i want write actual timestmap
	for i, v := range stat {
		fmt.Printf("%s\t%d\t%d\n", i, int64(v), t0.Unix())
	}

	// fmt.Printf("%+v\n", stat)
}
