package model

import "fmt"

func ErrCommit(err error) error {
	return fmt.Errorf("failed to commit Tx due to error: %v", err)
}

func ErrRollback(err error) error {
	return fmt.Errorf("failed to rollback Tx due to error: %v", err)
}

func ErrCreateTx(err error) error {
	return fmt.Errorf("failed to create Tx due to error: %v", err)
}

func ErrCreateQuery(err error) error {
	return fmt.Errorf("failed to create SQL Query due to error: %v", err)
}

func ErrScan(err error) error {
	return fmt.Errorf("failed to scan due to error: %v", err)
}

func ErrDoQuery(err error) error {
	return fmt.Errorf("failed to query due to error: %v", err)
}
