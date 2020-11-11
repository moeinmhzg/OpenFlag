package request

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/OpenFlag/OpenFlag/internal/app/openflag/constraint"
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	minVariantLen = 1
	minSegmentLen = 1

	nameFormat = `^[a-z0-9]+(?:\.[a-z0-9]+)*$`
)

//nolint:gochecknoglobals
var (
	nameRegex = regexp.MustCompile(nameFormat)
)

type (
	// Variant represents the possible variation of a flag. For example, control/treatment, green/yellow/red, etc.
	// VariantAttachment represents the dynamic configuration of a variant. For example,
	// if you have a variant for the green button,
	// you can dynamically control what's the hex color of green you want to use (e.g. {"hex_color": "#42b983"}).
	Variant struct {
		Key        string          `json:"key"`
		Attachment json.RawMessage `json:"attachment,omitempty"`
	}

	// Constraint represents rules that we can use to define the audience of the segment.
	// In other words, the audience in the segment is defined by a set of constraints.
	Constraint struct {
		Name       string          `json:"name"`
		Parameters json.RawMessage `json:"parameters"`
	}

	// Segment represents the segmentation, i.e. the set of audience we want to target.
	Segment struct {
		Description     string     `json:"description,omitempty"`
		Constraint      Constraint `json:"constraint"`
		Variants        []Variant  `json:"variants"`
		DefaultVariants []Variant  `json:"default_variants,omitempty"`
	}

	// Flag represents a feature flag, an experiment, or a configuration.
	Flag struct {
		Tags        []string  `json:"tags,omitempty"`
		Description string    `json:"description"`
		Flag        string    `json:"flag"`
		Segments    []Segment `json:"segments"`
	}

	// CreateFlagRequest represents a request body for creating a flag.
	CreateFlagRequest struct {
		Flag
	}
)

// Validate validates Variant struct.
func (v Variant) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(
			&v.Key,
			validation.Required,
			validation.Match(nameRegex),
		),
	)
}

// Validate validates Constraint struct.
func (c Constraint) Validate() error {
	return constraint.Validate(c.Name, c.Parameters)
}

// Validate validates Segment struct.
func (s Segment) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(
			&s.Constraint,
			validation.Required,
		),
		validation.Field(
			&s.Variants,
			validation.Required,
			validation.Length(minVariantLen, 0),
		),
		validation.Field(
			&s.DefaultVariants,
			validation.Length(minVariantLen, 0),
		),
	)
}

// Validate validates Flag struct.
func (f Flag) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(
			&f.Description,
			validation.Required,
		),
		validation.Field(
			&f.Flag,
			validation.Required,
			validation.Match(nameRegex),
		),
		validation.Field(
			&f.Segments,
			validation.Required,
			validation.Length(minSegmentLen, 0),
		),
		validation.Field(
			&f.Tags,
			validation.By(func(value interface{}) error {
				for _, tag := range f.Tags {
					if !nameRegex.MatchString(tag) {
						return errors.New("invalid tag format")
					}
				}

				return nil
			}),
		),
	)
}