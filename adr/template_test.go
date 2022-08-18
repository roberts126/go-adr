package adr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestTemplate(t *testing.T) {
	t.Run("TestDefaultTemplate", func(t *testing.T) {
		dv := "default value"
		truePtr := true
		min := 1
		max := 5

		expected := Template{
			data: nil,
			Variables: []*Variables{
				{
					Name:    "VariableName",
					Prompt:  "Prompt shown to user",
					Default: &dv,
					Validation: []*Validation{
						{
							Operation: "match",
							Args:      []string{"^[A-Z][A-Za-z0-9]+$"},
							Message:   "Invalid pattern for VariableName",
						},
					},
				},
				{
					Name:   "AnotherVariable",
					Prompt: "Enter another value:",
				},
				{
					Name:         "DecisionDrivers",
					Prompt:       "Please enter the initial decision driver:",
					RepeatPrompt: "Please enter an additional decision driver:",
					MinItems:     &min,
					Optional:     &truePtr,
					Repeats:      &truePtr,
					ExitValue:    "",
				},
				{
					Name:      "ConsideredOptions",
					Prompt:    "Please enter a considered option:",
					Repeats:   &truePtr,
					ExitValue: "",
					MaxItems:  &max,
				},
			},
			Contents: `# {{ .Title }}

## {{ .Context }}

{{ with .DecisionDrivers }}
## Decision Drivers
{{ range $driver := .DecisionDrivers }}
* {{ $driver }}
{{ end }}
{{ end }}

## Considered Options
{{ range $option := .ConsideredOptions }}
{{ $option }}
{{ end }}

## Decision Outcome

Chosen option: "{title of option 1}", because

{justification. e.g., only option, which meets k.o. criterion decision driver | which resolves force {force} | … | comes out best (see below)}.

<!-- This is an optional element. Feel free to remove. -->
## Positive Consequences

* Consequence #1
* Consequence #2
* Consequence #N

<!-- This is an optional element. Feel free to remove. -->
## Negative Consequences

* Consequence #1
* Consequence #2
* Consequence #N

<!-- This is an optional element. Feel free to remove. -->
## Validation

Validation information.

<!-- This is an optional element. Feel free to remove. -->
## Pros and Cons of the Options

### {title of option 1}

<!-- This is an optional element. Feel free to remove. -->
{example | description | pointer to more information | …}

* Good, because {argument a}
* Good, because {argument b}
<!-- use "neutral" if the given argument weights neither for good nor bad -->
* Neutral, because {argument c}
* Bad, because {argument d}
* … <!-- numbers of pros and cons can vary -->

### {title of other option}

{example | description | pointer to more information | …}

* Good, because {argument a}
* Good, because {argument b}
* Neutral, because {argument c}
* Bad, because {argument d}
* …

<!-- This is an optional element. Feel free to remove. -->
## More Information

{You might want to provide additional evidence/confidence for the decision outcome here and/or
document the team agreement on the decision and/or
define when this decision when and how the decision should be realized and if/when it should be re-visited and/or
how the decision is validated.
Links to other decisions and resources might here appear as well.}
`,
		}

		var actual Template
		if assert.NoError(t, yaml.Unmarshal([]byte(DefaultTemplate), &actual)) {
			assert.Equal(t, expected.Contents, actual.Contents, "Template contents must match")

			// Loop through the variables to identify failures more easily
			for i := 0; i < len(expected.Variables); i++ {
				val := expected.Variables[i]
				assert.Equalf(t, expected.Variables[i], actual.Variables[i], "Template variable %s must match", val.Name)
			}
		}
	})
}
