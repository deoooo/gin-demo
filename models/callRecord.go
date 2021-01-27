package models

type CallRecord struct {
	Model

	MobileNumber string `json:"mobileNumber"`
	Status       string `json:"status"`
}

func CreateCallRecord(mobileNumber string) (*CallRecord, error) {
	record := &CallRecord{
		MobileNumber: mobileNumber,
		Status:       "init",
	}
	result := db.Create(&record)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}
