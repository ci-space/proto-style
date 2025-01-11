package rules

func NewFieldWithBehaviorRule() *FieldWithOptionRule {
	return NewFieldWithOptionRule("PROTOSTYLE_FIELD_WITH_BEHAVIOR", "google.api.field_behavior", []string{
		"OPTIONAL",
		"REQUIRED",
		"OUTPUT_ONLY",
		"INPUT_ONLY",
		"IMMUTABLE",
		"UNORDERED_LIST",
		"NON_EMPTY_DEFAULT",
		"IDENTIFIER",
	})
}
