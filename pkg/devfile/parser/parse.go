package parser

import (
	"encoding/json"

	devfileCtx "github.com/openshift/odo/pkg/devfile/parser/context"
	"github.com/openshift/odo/pkg/devfile/parser/data"
	"github.com/pkg/errors"
)

// ParseDevfile func validates the devfile integrity.
// Creates devfile context and runtime objects
func ParseDevfile(d DevfileObj) (DevfileObj, error) {

	// Validate devfile
	err := d.Ctx.Validate()
	if err != nil {
		return d, err
	}

	// Create a new devfile data object
	d.Data, err = data.NewDevfileData(d.Ctx.GetApiVersion())
	if err != nil {
		return d, err
	}

	// Unmarshal devfile content into devfile struct
	err = json.Unmarshal(d.Ctx.GetDevfileContent(), &d.Data)
	if err != nil {
		return d, errors.Wrapf(err, "failed to decode devfile content")
	}

	// Successful
	return d, nil
}

// Parse func parses and validates the devfile integrity.
// Creates devfile context and runtime objects
func Parse(path string) (d DevfileObj, err error) {

	// NewDevfileCtx
	d.Ctx = devfileCtx.NewDevfileCtx(path)

	// Fill the fields of DevfileCtx struct
	err = d.Ctx.Populate()
	if err != nil {
		return d, err
	}
	return ParseDevfile(d)
}

// ParseInMemory func parses and validates the devfile integrity.
// Creates devfile context and runtime objects
func ParseInMemory(bytes []byte) (d DevfileObj, err error) {

	// Fill the fields of DevfileCtx struct
	err = d.Ctx.PopulateFromBytes(bytes)
	if err != nil {
		return d, err
	}
	return ParseDevfile(d)
}
