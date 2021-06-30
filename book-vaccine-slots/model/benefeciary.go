package model


type Beneficiary struct {
	Beneficiaries []struct {
		BeneficiaryReferenceID string `json:"beneficiary_reference_id"`
		Name                   string `json:"name"`
		BirthYear              string `json:"birth_year"`
		Gender                 string `json:"gender"`
		MobileNumber           string `json:"mobile_number"`
		PhotoIDType            string `json:"photo_id_type"`
		PhotoIDNumber          string `json:"photo_id_number"`
		ComorbidityInd         string `json:"comorbidity_ind"`
		VaccinationStatus      string `json:"vaccination_status"`
		Vaccine                string `json:"vaccine"`
		Dose1Date              string `json:"dose1_date"`
		Dose2Date              string `json:"dose2_date"`
		Appointments           []struct {
			AppointmentID string `json:"appointment_id"`
			CenterID      int    `json:"center_id"`
			Name          string `json:"name"`
			StateName     string `json:"state_name"`
			DistrictName  string `json:"district_name"`
			BlockName     string `json:"block_name"`
			From          string `json:"from"`
			To            string `json:"to"`
			Dose          int    `json:"dose"`
			SessionID     string `json:"session_id"`
			Date          string `json:"date"`
			Slot          string `json:"slot"`
		} `json:"appointments"`
	} `json:"beneficiaries"`
}
