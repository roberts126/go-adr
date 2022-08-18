package adr

const ExampleConfiguration = `# ---
# projects:
#   - name: Example
#     directory: /some/path/to/repo/docs/adr
#     template: '~/.config/adr/templates/default.tpl'`

// The default template is a Go Template implementation of
// https://raw.githubusercontent.com/adr/madr/main/template/adr-template.md

const DefaultTemplate = `---
variables:
    - name: VariableName
      prompt: Prompt shown to user
      default: default value  # Optional
      validation:  # Optional
        - operation: match
          message: "Invalid pattern for VariableName"
          args:
          - '^[A-Z][A-Za-z0-9]+$'
    - name: AnotherVariable
      prompt: 'Enter another value:'
    - name: DecisionDrivers
      prompt: 'Please enter the initial decision driver:'
      optional: true
      repeats: true
      minItems: 1
      repeatPrompt: 'Please enter an additional decision driver:'
      exitValue: ''
    - name: ConsideredOptions
      prompt: 'Please enter a considered option:'
      maxItems: 5
      repeats: true
      exitValue: ''
contents: |
    # {{ .Title }}

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
`
