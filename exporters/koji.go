package exporters

import (
	"context"
	"net/url"

	"github.com/paulmach/orb/geojson"
	"github.com/sirupsen/logrus"

	"github.com/UnownHash/Fletchling/koji_client"
)

type KojiExporter struct {
	logger        *logrus.Logger
	kojiCli       *koji_client.APIClient
	projectName   string
	projectParams url.Values
}

func (*KojiExporter) ExporterName() string {
	return "koji"
}

func (exporter *KojiExporter) ExportFeatures(ctx context.Context) ([]*geojson.Feature, error) {
	fc, err := exporter.kojiCli.GetFeatureCollection(ctx, exporter.projectName, exporter.projectParams)
	if err != nil {
		return nil, err
	}

	if len(fc.Features) == 0 {
		return nil, nil
	}

	features := make([]*geojson.Feature, len(fc.Features))
	idx := 0

	for _, feature := range fc.Features {
		name, _ := feature.Properties["name"].(string)
		if name == "" {
			name = "<unknown>"
		}
		features[idx] = feature
		idx++
	}

	features = features[:idx]

	return features, nil

}

func NewKojiExporter(logger *logrus.Logger, kojiCli *koji_client.APIClient, projectName string, projectParams url.Values) (*KojiExporter, error) {
	exporter := &KojiExporter{
		logger:        logger,
		kojiCli:       kojiCli,
		projectName:   projectName,
		projectParams: projectParams,
	}
	return exporter, nil
}
