package salary
import (
    "math"
    opt "github.com/eatmoreapple/optional"
)

var salaryThresholdForHigherContributions int64 = 4300000

type Employee struct {
     Id string
     Name string
     AnnualGrossSalaryInCents  int64
}

type SalarySlip struct {
     EmployeeId string
     EmployeeName string
     GrossSalaryInCents int64
     NationalInsuranceContributionsInCents opt.Option[int64]
     Taxes opt.Option[Taxes]
}

func GenerateSalarySlipFor(employee Employee) SalarySlip {
    nationalInsuranceContributions := CalculateNationalInsurance(employee.AnnualGrossSalaryInCents)
    grossSalary := int64(math.Ceil(float64(employee.AnnualGrossSalaryInCents) / 12))
    taxes := CreateTaxes(employee.AnnualGrossSalaryInCents)
    salarySlip := SalarySlip{
                    EmployeeId: employee.Id,
                    EmployeeName: employee.Name,
                    GrossSalaryInCents: grossSalary,
                    NationalInsuranceContributionsInCents: nationalInsuranceContributions,
                    Taxes: taxes,
                  }
    return salarySlip
}
