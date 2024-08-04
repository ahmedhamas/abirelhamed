package handlers

import (
	"abir-el-hamd/internal/config"
	"abir-el-hamd/internal/middleware"
	"abir-el-hamd/internal/models"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/base.html", "template/index.html"))
	tmpl.Execute(w, nil)
}

func HomeApi(w http.ResponseWriter, r *http.Request) {
	database := config.Database()
	defer database.Close()

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil {
		er := errors.New("invalid number")
		middleware.ErrorResopnse(w, er)
		return
	}

	offset := limit - 30

	CasesTable := models.Cases{}

	Cases, err := CasesTable.GetAll(database, 30, offset)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	pages, err := CasesTable.NumberOfPages(database)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	Res := map[string]interface{}{
		"Cases": Cases,
		"Pages": pages,
	}

	if err := json.NewEncoder(w).Encode(Res); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}

func AddCase(w http.ResponseWriter, r *http.Request) {
	db := config.Database()
	defer db.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	bodyData := make(map[string]interface{})

	if err := json.Unmarshal(body, &bodyData); err != nil {
		fmt.Println(err)
		middleware.ErrorResopnse(w, err)
		return
	}

	totalIncome, err := strconv.Atoi(bodyData["total_income"].(string))
	if err != nil {
		log.Println("Error converting total_income:", err)
		return
	}

	fixedExpenses, err := strconv.Atoi(bodyData["fixed_expenses"].(string))
	if err != nil {
		log.Println("Error converting fixed_expenses:", err)
		return
	}

	Age, err := strconv.Atoi(bodyData["age"].(string))
	if err != nil {
		log.Println("Error converting age:", err)
		return
	}

	now := time.Now()
	createdAt := sql.NullString{String: now.Format(time.RFC3339), Valid: true}

	Case := models.Cases{
		Case_name:                     getString(bodyData["case_name"]),
		National_id:                   getString(bodyData["national_id"]),
		Devices_needed_for_the_case:   getString(bodyData["devices_needed_for_the_case"]),
		Total_income:                  totalIncome,
		Fixed_expenses:                fixedExpenses,
		Pension_from_husband:          getString(bodyData["pension_from_husband"]),
		Pension_from_father:           getString(bodyData["pension_from_father"]),
		Debts:                         getString(bodyData["debts"]),
		Case_type:                     getString(bodyData["case_type"]),
		Date_of_birth:                 sql.NullString{String: getString(bodyData["date_of_birth"]), Valid: true},
		Age:                           Age,
		Gender:                        getString(bodyData["gender"]),
		Job:                           getString(bodyData["job"]),
		Social_situation:              getString(bodyData["social_situation"]),
		Address_from_national_id_card: getString(bodyData["address_from_national_id_card"]),
		Actual_address:                getString(bodyData["actual_address"]),
		District:                      getString(bodyData["district"]),
		Created_at:                    createdAt,
		Updated_at:                    createdAt,
	}

	if err := Case.Create(db); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"status": "success"}); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}

func getString(value interface{}) string {
	if value == nil {
		return ""
	}
	return value.(string)
}

func DeleteCase(w http.ResponseWriter, r *http.Request) {
	db := config.Database()
	defer db.Close()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Case := models.Cases{
		Id: id,
	}

	fmt.Println(Case)

	if err := Case.Delete(db); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	if err := json.NewEncoder(w).Encode(map[string]interface{}{"status": "success"}); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

}

func GetCase(w http.ResponseWriter, r *http.Request) {
	var script string = "/assets/case.js"
	var style string = "/assets/case.css"

	data := map[string]interface{}{
		"script": script,
		"style":  style,
	}

	tmpl := template.Must(template.ParseFiles("template/base.html", "template/case.html"))

	tmpl.Execute(w, data)
}

func CaseApi(w http.ResponseWriter, r *http.Request) {
	db := config.Database()
	defer db.Close()

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	Case := models.CaseDitails{
		Id: id,
	}

	cas, err := Case.Get(db)

	if err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cas); err != nil {
		middleware.ErrorResopnse(w, err)
		return
	}
}
