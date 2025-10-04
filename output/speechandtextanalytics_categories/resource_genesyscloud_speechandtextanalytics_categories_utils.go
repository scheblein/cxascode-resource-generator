package speechandtextanalytics_categories

import (
	"terraform-provider-genesyscloud/genesyscloud/util/resourcedata"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mypurecloud/platform-client-sdk-go/v133/platformclientv2"
)

/*
The resource_genesyscloud_speechandtextanalytics_categories_utils.go file contains various helper methods to marshal
and unmarshal data into formats consumable by Terraform and/or Genesys Cloud.
*/

// getSpeechandtextanalyticsCategoriesFromResourceData maps data from schema ResourceData object to a platformclientv2.Stacategory
func getSpeechandtextanalyticsCategoriesFromResourceData(d *schema.ResourceData) platformclientv2.Stacategory {
	return platformclientv2.Stacategory{
		Name:            platformclientv2.String(d.Get("name").(string)),
		Description:     platformclientv2.String(d.Get("description").(string)),
		InteractionType: platformclientv2.String(d.Get("interaction_type").(string)),
		Criteria:        buildOperand(d.Get("criteria").([]interface{})),
	}
}

// buildTerms maps an []interface{} into a Genesys Cloud *[]platformclientv2.Term
func buildTerms(terms []interface{}) *[]platformclientv2.Term {
	termsSlice := make([]platformclientv2.Term, 0)
	for _, term := range terms {
		var sdkTerm platformclientv2.Term
		termsMap, ok := term.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkTerm.Word, termsMap, "word")
		resourcedata.BuildSDKStringValueIfNotNil(&sdkTerm.ParticipantType, termsMap, "participant_type")

		termsSlice = append(termsSlice, sdkTerm)
	}

	return &termsSlice
}

// buildOperandPositions maps an []interface{} into a Genesys Cloud *[]platformclientv2.Operandposition
func buildOperandPositions(operandPositions []interface{}) *[]platformclientv2.Operandposition {
	operandPositionsSlice := make([]platformclientv2.Operandposition, 0)
	for _, operandPosition := range operandPositions {
		var sdkOperandPosition platformclientv2.Operandposition
		operandPositionsMap, ok := operandPosition.(map[string]interface{})
		if !ok {
			continue
		}

		sdkOperandPosition.StartingPositionValue = platformclientv2.Int(operandPositionsMap["starting_position_value"].(int))
		resourcedata.BuildSDKStringValueIfNotNil(&sdkOperandPosition.StartingPositionDirection, operandPositionsMap, "starting_position_direction")
		sdkOperandPosition.EndingPositionValue = platformclientv2.Int(operandPositionsMap["ending_position_value"].(int))
		resourcedata.BuildSDKStringValueIfNotNil(&sdkOperandPosition.EndingPositionDirection, operandPositionsMap, "ending_position_direction")

		operandPositionsSlice = append(operandPositionsSlice, sdkOperandPosition)
	}

	return &operandPositionsSlice
}

// buildOperatorPositions maps an []interface{} into a Genesys Cloud *[]platformclientv2.Operatorposition
func buildOperatorPositions(operatorPositions []interface{}) *[]platformclientv2.Operatorposition {
	operatorPositionsSlice := make([]platformclientv2.Operatorposition, 0)
	for _, operatorPosition := range operatorPositions {
		var sdkOperatorPosition platformclientv2.Operatorposition
		operatorPositionsMap, ok := operatorPosition.(map[string]interface{})
		if !ok {
			continue
		}

		sdkOperatorPosition.VoiceSecondsPosition = platformclientv2.Int(operatorPositionsMap["voice_seconds_position"].(int))
		sdkOperatorPosition.DigitalWordsPosition = platformclientv2.Int(operatorPositionsMap["digital_words_position"].(int))

		operatorPositionsSlice = append(operatorPositionsSlice, sdkOperatorPosition)
	}

	return &operatorPositionsSlice
}

// buildInfixOperators maps an []interface{} into a Genesys Cloud *[]platformclientv2.Infixoperator
func buildInfixOperators(infixOperators []interface{}) *[]platformclientv2.Infixoperator {
	infixOperatorsSlice := make([]platformclientv2.Infixoperator, 0)
	for _, infixOperator := range infixOperators {
		var sdkInfixOperator platformclientv2.Infixoperator
		infixOperatorsMap, ok := infixOperator.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkInfixOperator.OperatorType, infixOperatorsMap, "operator_type")
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkInfixOperator.OperatorPosition, infixOperatorsMap, "operator_position", buildOperatorPosition)

		infixOperatorsSlice = append(infixOperatorsSlice, sdkInfixOperator)
	}

	return &infixOperatorsSlice
}

// buildOperands maps an []interface{} into a Genesys Cloud *[]platformclientv2.Operand
func buildOperands(operands []interface{}) *[]platformclientv2.Operand {
	operandsSlice := make([]platformclientv2.Operand, 0)
	for _, operand := range operands {
		var sdkOperand platformclientv2.Operand
		operandsMap, ok := operand.(map[string]interface{})
		if !ok {
			continue
		}

		operandsSlice = append(operandsSlice, sdkOperand)
	}

	return &operandsSlice
}

// buildOperands maps an []interface{} into a Genesys Cloud *[]platformclientv2.Operand
func buildOperands(operands []interface{}) *[]platformclientv2.Operand {
	operandsSlice := make([]platformclientv2.Operand, 0)
	for _, operand := range operands {
		var sdkOperand platformclientv2.Operand
		operandsMap, ok := operand.(map[string]interface{})
		if !ok {
			continue
		}

		resourcedata.BuildSDKStringValueIfNotNil(&sdkOperand.Type, operandsMap, "type")
		sdkOperand.Occurrence = platformclientv2.Int(operandsMap["occurrence"].(int))
		sdkOperand.Inverted = platformclientv2.Bool(operandsMap["inverted"].(bool))
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkOperand.Term, operandsMap, "term", buildTerm)
		resourcedata.BuildSDKStringValueIfNotNil(&sdkOperand.TopicId, operandsMap, "topic_id")
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkOperand.VoiceSecondsPosition, operandsMap, "voice_seconds_position", buildOperandPosition)
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkOperand.DigitalWordsPosition, operandsMap, "digital_words_position", buildOperandPosition)
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkOperand.InfixOperator, operandsMap, "infix_operator", buildInfixOperator)
		resourcedata.BuildSDKInterfaceArrayValueIfNotNil(&sdkOperand.Operands, operandsMap, "operands", buildOperands)

		operandsSlice = append(operandsSlice, sdkOperand)
	}

	return &operandsSlice
}

// flattenTerms maps a Genesys Cloud *[]platformclientv2.Term into a []interface{}
func flattenTerms(terms *[]platformclientv2.Term) []interface{} {
	if len(*terms) == 0 {
		return nil
	}

	var termList []interface{}
	for _, term := range *terms {
		termMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(termMap, "word", term.Word)
		resourcedata.SetMapValueIfNotNil(termMap, "participant_type", term.ParticipantType)

		termList = append(termList, termMap)
	}

	return termList
}

// flattenOperandPositions maps a Genesys Cloud *[]platformclientv2.Operandposition into a []interface{}
func flattenOperandPositions(operandPositions *[]platformclientv2.Operandposition) []interface{} {
	if len(*operandPositions) == 0 {
		return nil
	}

	var operandPositionList []interface{}
	for _, operandPosition := range *operandPositions {
		operandPositionMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(operandPositionMap, "starting_position_value", operandPosition.StartingPositionValue)
		resourcedata.SetMapValueIfNotNil(operandPositionMap, "starting_position_direction", operandPosition.StartingPositionDirection)
		resourcedata.SetMapValueIfNotNil(operandPositionMap, "ending_position_value", operandPosition.EndingPositionValue)
		resourcedata.SetMapValueIfNotNil(operandPositionMap, "ending_position_direction", operandPosition.EndingPositionDirection)

		operandPositionList = append(operandPositionList, operandPositionMap)
	}

	return operandPositionList
}

// flattenOperatorPositions maps a Genesys Cloud *[]platformclientv2.Operatorposition into a []interface{}
func flattenOperatorPositions(operatorPositions *[]platformclientv2.Operatorposition) []interface{} {
	if len(*operatorPositions) == 0 {
		return nil
	}

	var operatorPositionList []interface{}
	for _, operatorPosition := range *operatorPositions {
		operatorPositionMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(operatorPositionMap, "voice_seconds_position", operatorPosition.VoiceSecondsPosition)
		resourcedata.SetMapValueIfNotNil(operatorPositionMap, "digital_words_position", operatorPosition.DigitalWordsPosition)

		operatorPositionList = append(operatorPositionList, operatorPositionMap)
	}

	return operatorPositionList
}

// flattenInfixOperators maps a Genesys Cloud *[]platformclientv2.Infixoperator into a []interface{}
func flattenInfixOperators(infixOperators *[]platformclientv2.Infixoperator) []interface{} {
	if len(*infixOperators) == 0 {
		return nil
	}

	var infixOperatorList []interface{}
	for _, infixOperator := range *infixOperators {
		infixOperatorMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(infixOperatorMap, "operator_type", infixOperator.OperatorType)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(infixOperatorMap, "operator_position", infixOperator.OperatorPosition, flattenOperatorPosition)

		infixOperatorList = append(infixOperatorList, infixOperatorMap)
	}

	return infixOperatorList
}

// flattenOperands maps a Genesys Cloud *[]platformclientv2.Operand into a []interface{}
func flattenOperands(operands *[]platformclientv2.Operand) []interface{} {
	if len(*operands) == 0 {
		return nil
	}

	var operandList []interface{}
	for _, operand := range *operands {
		operandMap := make(map[string]interface{})

		operandList = append(operandList, operandMap)
	}

	return operandList
}

// flattenOperands maps a Genesys Cloud *[]platformclientv2.Operand into a []interface{}
func flattenOperands(operands *[]platformclientv2.Operand) []interface{} {
	if len(*operands) == 0 {
		return nil
	}

	var operandList []interface{}
	for _, operand := range *operands {
		operandMap := make(map[string]interface{})

		resourcedata.SetMapValueIfNotNil(operandMap, "type", operand.Type)
		resourcedata.SetMapValueIfNotNil(operandMap, "occurrence", operand.Occurrence)
		resourcedata.SetMapValueIfNotNil(operandMap, "inverted", operand.Inverted)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(operandMap, "term", operand.Term, flattenTerm)
		resourcedata.SetMapValueIfNotNil(operandMap, "topic_id", operand.TopicId)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(operandMap, "voice_seconds_position", operand.VoiceSecondsPosition, flattenOperandPosition)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(operandMap, "digital_words_position", operand.DigitalWordsPosition, flattenOperandPosition)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(operandMap, "infix_operator", operand.InfixOperator, flattenInfixOperator)
		resourcedata.SetMapInterfaceArrayWithFuncIfNotNil(operandMap, "operands", operand.Operands, flattenOperands)

		operandList = append(operandList, operandMap)
	}

	return operandList
}
