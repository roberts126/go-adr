package adr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestTemplate(t *testing.T) {
	t.Run("TestDefaultTemplate", func(t *testing.T) {
		truePtr := true
		min := 1
		max := 5

		expected := Template{
			data: nil,
			Variables: []*Variable{
				{
					Name:     "DecisionDrivers",
					Prompt:   "Please enter the initial decision driver:",
					Optional: &truePtr,
					Repeat: &Repeat{
						MinItems: &min,
						Prompt:   "Please enter an additional decision driver:",
					},
				},
				{
					Name:   "ConsideredOptions",
					Prompt: "Please enter a considered option:",
					Repeat: &Repeat{
						MaxItems: &max,
					},
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

			expectedVars := make(map[string]*Variable, len(expected.Variables))
			for _, v := range expected.Variables {
				expectedVars[v.Name] = v
			}

			actualVars := make(map[string]*Variable, len(actual.Variables))
			for _, v := range actual.Variables {
				actualVars[v.Name] = v
			}

			// Loop through the variables to identify failures more easily and ensure completeness
			for k, expectedVar := range expectedVars {
				actualVar, ok := actualVars[k]

				if assert.Truef(t, ok, "Actual variables must contain %s", k) {
					assert.Equalf(t, expectedVar, actualVar, "Template variable %s must match", k)
				}
			}
		}
	})
}
