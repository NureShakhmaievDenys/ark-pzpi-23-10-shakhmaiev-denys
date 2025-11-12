package main

import "fmt"

func proc(d map[string]string) error {
	if len(d["password"]) < 8 {
		return fmt.Errorf("password too short")
	}

	if d["email"] == "" {
		return fmt.Errorf("invalid email")
	}
	return nil
}
func ValidateUserCredentials(userData map[string]string) error {
	if len(userData["password"]) < 8 {
		return fmt.Errorf("password too short")
	}

	if userData["email"] == "" {
		return fmt.Errorf("invalid email")
	}
	return nil
}

type ReportGenerator struct{}

func (r *ReportGenerator) CreateReport() {

	connectionString := r.getDBConnection()
	fmt.Println("Connecting with:", connectionString)

}
func (r *ReportGenerator) getDBConnection() string {
	return "user:pass@tcp(127.0.0.1)/db"
}

func (r *ReportGenerator) GetDBConnection() string {
	return "user:pass@tcp(127.0.0.1)/db"
}

/*
type TextReport struct {}
func (t *TextReport) Generate() {
	fmt.Println("1. Getting text data...")
	fmt.Println("2. Formatting data as Text")
	fmt.Println("3. Saving to .txt file")
}

type CsvReport struct {}
func (c *CsvReport) Generate() {
	fmt.Println("1. Getting CSV data...")
	fmt.Println("2. Formatting data as CSV")
	fmt.Println("3. Saving to .csv file")
}
*/
type reportImplementation interface {
	getData()
	formatData()
	saveFile()
}

func GenerateReport(impl reportImplementation) {
	impl.getData()
	impl.formatData()
	impl.saveFile()
}

type textReport struct{}

func (t *textReport) getData()    { fmt.Println("1. Getting text data...") }
func (t *textReport) formatData() { fmt.Println("2. Formatting data as Text") }
func (t *textReport) saveFile()   { fmt.Println("3. Saving to .txt file") }

type csvReport struct{}

func (c *csvReport) getData()    { fmt.Println("1. Getting CSV data...") }
func (c *csvReport) formatData() { fmt.Println("2. Formatting data as CSV") }
func (c *csvReport) saveFile()   { fmt.Println("3. Saving to .csv file") }

func main() {
	fmt.Println("--- Запуск Прикладу 1 (Rename Method) ---")

	userData := map[string]string{"email": "test@ex.com", "password": "123"}
	if err := ValidateUserCredentials(userData); err != nil {
		fmt.Println("Помилка валідації:", err)
	}

	userDataGood := map[string]string{"email": "test@ex.com", "password": "longpassword"}
	if err := ValidateUserCredentials(userDataGood); err == nil {
		fmt.Println("Валідація успішна")
	}

	fmt.Println("\n--- Запуск Прикладу 2 (Hide Method) ---")
	reportGen := &ReportGenerator{}
	reportGen.CreateReport()
	fmt.Println("\n--- Запуск Прикладу 3 (Template Method) ---")
	txtReport := &textReport{}
	GenerateReport(txtReport)

	fmt.Println("...")

	csvRep := &csvReport{}
	GenerateReport(csvRep)
}
