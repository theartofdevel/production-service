package validator

type chainValidator []Validator

func (v chainValidator) Validate() error {
	for _, validator := range v {
		if err := validator.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func ChainValidator(validators ...Validator) Validator {
	return chainValidator(validators)
}
