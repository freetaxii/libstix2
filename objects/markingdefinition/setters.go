package markingdefinition

/*
SetDefinitionType - This method sets the definition type property to s.
*/
func (o *MarkingDefinition) SetDefinitionType(s string) (err error) {
	o.DefinitionType = s
	return nil
}

/*
GetDefinitionType - This method get the definition type.
*/
func (o *MarkingDefinition) GetDefinitionType() (s string) {
	return o.DefinitionType
}
