package rules

func NewFieldWithBehaviorRule() *FieldWithOptionRule {
	return NewFieldWithOptionRule("PROTOSTYLE_FIELD_WITH_BEHAVIOR", "google.api.field_behavior")
}
