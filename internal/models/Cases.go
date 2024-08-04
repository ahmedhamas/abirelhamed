package models

import (
	"database/sql"
	"fmt"
)

type Cases struct {
	Id                            int            `json:"id"`
	Case_name                     string         `json:"case_name"`
	National_id                   string         `json:"national_id"`
	Devices_needed_for_the_case   string         `json:"devices_needed_for_the_case"`
	Total_income                  int            `json:"total_income"`
	Fixed_expenses                int            `json:"fixed_expenses"`
	Pension_from_husband          string         `json:"pension_from_husband"`
	Pension_from_father           string         `json:"pension_from_father"`
	Debts                         string         `json:"debts"`
	Case_type                     string         `json:"case_type"`
	Date_of_birth                 sql.NullString `json:"date_of_birth"`
	Age                           int            `json:"age"`
	Gender                        string         `json:"gender"`
	Job                           string         `json:"job"`
	Social_situation              string         `json:"social_situation"`
	Address_from_national_id_card string         `json:"address_from_national_id_card"`
	Actual_address                string         `json:"actual_address"`
	District                      string         `json:"district"`
	Subsidies_id                  sql.NullInt32  `json:"subsidies_id"`
	Social_status                 sql.NullInt32  `json:"social_status"`
	Husband_id                    sql.NullInt32  `json:"husband_id"`
	Created_at                    sql.NullString `json:"created_at"`
	Updated_at                    sql.NullString `json:"updated_at"`
}

type Relative struct {
	Relative_type             sql.NullString
	Relative_name             string
	Relative_national_id      sql.NullString
	Relative_date_of_birth    sql.NullString
	Relative_age              int
	Relative_gender           string
	Relative_job              string
	Relative_social_situation string
	Relative_health_status    string
	Relative_education        string
}

type CaseDitails struct {
	Id                                            int
	Case_name                                     string
	National_id                                   sql.NullString
	Devices_needed_for_the_case                   string
	Total_income                                  int
	Fixed_expenses                                int
	Pension_from_husband                          int
	Pension_from_father                           int
	Debts                                         string
	Case_type                                     string
	Date_of_birth                                 sql.NullString
	Age                                           int
	Gender                                        string
	Job                                           string
	Social_situation                              string
	Address_from_national_id_card                 string
	Actual_address                                string
	District                                      string
	Created_at                                    sql.NullString
	Updated_at                                    sql.NullString
	Husband_name                                  string
	Husband_national_id                           sql.NullString
	Husband_date_of_birth                         sql.NullString
	Husband_age                                   int
	Husband_gender                                string
	Properties                                    string
	Health_status                                 string
	Education                                     string
	Number_of_family_members                      int
	Number_of_registered_children                 int
	Total_number_of_children                      int
	Grants_from_outside_the_association           string
	Grants_from_outside_the_association_financial string
	Grants_from_the_association_financial         string
	Grants_from_the_association_inKind            string
	Total_Subsidies                               int
	Relatives                                     []Relative
}

func (ca Cases) Create(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO `cases` (`case_name`, `national_id`, `devices_needed_for_the_case`, `total_income`, `fixed_expenses`, `pension_from_husband`, `pension_from_father`, `debts`, `case_type`, `date_of_birth`, `age`, `gender`, `job`, `social_situation`, `address_from_national_id_card`, `actual_address`, `district`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		ca.Case_name, ca.National_id, ca.Devices_needed_for_the_case, ca.Total_income, ca.Fixed_expenses,
		ca.Pension_from_husband, ca.Pension_from_father, ca.Debts, ca.Case_type,
		ca.Date_of_birth, ca.Age, ca.Gender, ca.Job, ca.Social_situation,
		ca.Address_from_national_id_card, ca.Actual_address, ca.District, ca.Created_at, ca.Updated_at)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (ca Cases) Update(db *sql.DB) error {

	_, err := db.Exec("UPDATE `cases` SET `case_name` = ?,`national_id` = ?,`devices_needed_for_the_case` = ?,`total_income` = ?,`fixed_expenses` = ?,`pension_from_husband` = ?,`pension_from_father` = ?,`debts` = ?,`case_type` = ?,`date_of_birth` = ?,`age` = ?,`gender` = ?,`job` = ?,`social_situation` = ?,`address_from_national_id_card` = ?,`actual_address` = ?,`district` = ?,`subsidies_id` = ?,`social_status` = ?,`husband_id` = ?,`created_at` = ?,`updated_at` = ? WHERE `id` = ?",
		ca.Case_name, ca.National_id, ca.Devices_needed_for_the_case, ca.Total_income, ca.Fixed_expenses,
		ca.Pension_from_husband, ca.Pension_from_father, ca.Debts, ca.Case_type,
		ca.Date_of_birth, ca.Age, ca.Gender, ca.Job, ca.Social_situation,
		ca.Address_from_national_id_card, ca.Actual_address, ca.District, ca.Subsidies_id,
		ca.Social_status, ca.Husband_id, ca.Created_at, ca.Updated_at, ca.Id)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (ca Cases) Delete(db *sql.DB) error {

	_, err := db.Exec("DELETE FROM cases WHERE id = ?", ca.Id)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (ca CaseDitails) Get(db *sql.DB) (CaseDitails, error) {
	query := `
		SELECT DISTINCT
			cases.id, cases.case_name, cases.national_id, cases.devices_needed_for_the_case, cases.total_income,
			cases.fixed_expenses, cases.pension_from_husband, cases.pension_from_father, cases.debts, cases.case_type,
			cases.date_of_birth, cases.age, cases.gender, cases.job, cases.social_situation, cases.address_from_national_id_card,
			cases.actual_address, cases.district, cases.created_at, cases.updated_at, husband.name AS husband_name,
			husband.national_id AS husband_national_id, husband.date_of_birth AS husband_date_of_birth, husband.age AS husband_age,
			husband.gender AS husband_gender, socialstatusofthecase.properties, socialstatusofthecase.health_status,
			socialstatusofthecase.education, socialstatusofthecase.number_of_family_members, socialstatusofthecase.number_of_registered_children,
			socialstatusofthecase.total_number_of_children, subsidies.grants_from_outside_the_association,
			subsidies.grants_from_outside_the_association_financial, subsidies.grants_from_the_association_financial,
			subsidies.grants_from_the_association_inKind, subsidies.total_Subsidies
		FROM eabir_alhamd.cases
		INNER JOIN subsidies ON subsidies.id = cases.subsidies_id
		INNER JOIN socialstatusofthecase ON socialstatusofthecase.id = cases.social_status
		INNER JOIN husband ON husband.id = cases.husband_id
		WHERE cases.id = ?
	`

	Case := db.QueryRow(query, ca.Id)

	var cas CaseDitails
	if err := Case.Scan(
		&cas.Id, &cas.Case_name, &cas.National_id, &cas.Devices_needed_for_the_case, &cas.Total_income,
		&cas.Fixed_expenses, &cas.Pension_from_husband, &cas.Pension_from_father, &cas.Debts, &cas.Case_type,
		&cas.Date_of_birth, &cas.Age, &cas.Gender, &cas.Job, &cas.Social_situation, &cas.Address_from_national_id_card,
		&cas.Actual_address, &cas.District, &cas.Created_at, &cas.Updated_at, &cas.Husband_name, &cas.Husband_national_id,
		&cas.Husband_date_of_birth, &cas.Husband_age, &cas.Husband_gender, &cas.Properties, &cas.Health_status,
		&cas.Education, &cas.Number_of_family_members, &cas.Number_of_registered_children, &cas.Total_number_of_children,
		&cas.Grants_from_outside_the_association, &cas.Grants_from_outside_the_association_financial, &cas.Grants_from_the_association_financial,
		&cas.Grants_from_the_association_inKind, &cas.Total_Subsidies,
	); err != nil {
		fmt.Println(err)
		return CaseDitails{}, fmt.Errorf("error: %v", err)
	}

	relativesQuery := `
		SELECT DISTINCT relative.relative_type, relative.name, relative.national_id, relative.date_of_birth, relative.age,
			relative.gender, relative.job, relative.social_situation, relative.health_status, relative.education
		FROM relative
		WHERE relative.case_id = ?
	`

	rows, err := db.Query(relativesQuery, cas.Id)
	if err != nil {
		return CaseDitails{}, fmt.Errorf("error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var rel Relative

		if err := rows.Scan(
			&rel.Relative_type, &rel.Relative_name, &rel.Relative_national_id, &rel.Relative_date_of_birth, &rel.Relative_age,
			&rel.Relative_gender, &rel.Relative_job, &rel.Relative_social_situation, &rel.Relative_health_status, &rel.Relative_education,
		); err != nil {
			fmt.Println(err)
			return CaseDitails{}, fmt.Errorf("error: %v", err)
		}
		cas.Relatives = append(cas.Relatives, rel)
	}

	if err := rows.Err(); err != nil {
		return CaseDitails{}, fmt.Errorf("error: %v", err)
	}

	return cas, nil
}

func (ca Cases) GetAll(db *sql.DB, limit, offset int) ([]Cases, error) {
	cases := []Cases{}

	rows, err := db.Query("SELECT * FROM `cases` LIMIT ? OFFSET ?", limit, offset)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var cas Cases

		if err := rows.Scan(&cas.Id, &cas.Case_name, &cas.National_id, &cas.Devices_needed_for_the_case, &cas.Total_income,
			&cas.Fixed_expenses, &cas.Pension_from_husband, &cas.Pension_from_father, &cas.Debts,
			&cas.Case_type, &cas.Date_of_birth, &cas.Age, &cas.Gender, &cas.Job, &cas.Social_situation,
			&cas.Address_from_national_id_card, &cas.Actual_address, &cas.District, &cas.Subsidies_id,
			&cas.Social_status, &cas.Husband_id, &cas.Created_at, &cas.Updated_at); err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("error: %v", err)
		}

		cases = append(cases, cas)
	}

	return cases, nil
}

func (ca Cases) NumberOfPages(db *sql.DB) (int, error) {
	Case := db.QueryRow("SELECT COUNT(*) AS length FROM cases")

	var length int
	if err := Case.Scan(&length); err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}

	return length, nil
}
