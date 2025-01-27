package cyclonedxjson

import (
	"io"

	"github.com/CycloneDX/cyclonedx-go"
	"github.com/anchore/syft/internal/formats/common/cyclonedxhelpers"
	"github.com/anchore/syft/syft/sbom"
)

func encoder(output io.Writer, s sbom.SBOM) error {
	bom := cyclonedxhelpers.ToFormatModel(s)
	enc := cyclonedx.NewBOMEncoder(output, cyclonedx.BOMFileFormatJSON)

	err := enc.Encode(bom)
	return err
}
