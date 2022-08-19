package adr

// ExampleConfiguration is an example config used in the init commands and testing.
const ExampleConfiguration = `# ---
# projects:
#   '/some/path/to/repo/docs/adr':
#     name: Example
#     directory: '/some/path/to/repo/docs/adr'
#     template: '~/.config/adr/templates/default.tpl'`

// DefaultTemplate is a Go Template implementation of
// https://raw.githubusercontent.com/adr/madr/main/template/adr-template.md
const DefaultTemplate = `---
variables:
#    Example variable declaration
#    - name: Example # The name of the variable. Used in templates as {{ .VAR_NAME }}
#      prompt: Prompt shown to user # The initial prompt shown to the user
#      default: default value # Optional
#      validation:  # Optional
#        - operation: match # The name of the validation function
#          message: "Invalid pattern for VariableName" # The message displayed on failure
#          args: # An array of additional arguments supplied to the validator
#          - '^[A-Z][A-Za-z0-9]+$'
#      optional: false # Whether the variable is optional or not. Defaults to false
#      repeat: # Whether the repeats or not. Used to build lists. Defaults to nil
#        exitValue: '' # The value a user can enter to "exit" a repeated prompt.
#        prompt: 'Repeated Prompt:' # A prompt shown on repeated prompts. If empty the variable prompts is shown.
#        maxItems: 1 # Max number of items a list can contain
#        minItems: 1 # Min number of items a list must contain
    - name: DecisionDrivers
      prompt: 'Please enter the initial decision driver:'
      optional: true
      repeat:
        minItems: 1
        prompt: 'Please enter an additional decision driver:'
        exitValue: ''
    - name: ConsideredOptions
      prompt: 'Please enter a considered option:'
      repeat:
        maxItems: 5
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
