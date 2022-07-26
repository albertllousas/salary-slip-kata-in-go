package salary

import (
    "testing"
    opt "github.com/eatmoreapple/optional"
)

func Test_shouldCalculateTheSalarySlipForAnEmployee(t *testing.T){
    employee := Employee{Id: "12345", Name: "John J Doe", AnnualGrossSalaryInCents: 500000}

    got := GenerateSalarySlipFor(employee)

    expected := SalarySlip{
                    GrossSalaryInCents: 41667,
                    EmployeeName: "John J Doe",
                    EmployeeId: "12345",
                    NationalInsuranceContributionsInCents: opt.None[int64](),
                    Taxes: opt.None[Taxes](),
                }
    assertThat(got, expected, t)
}

func Test_shouldCalculateTheSalarySlipForAnEmployeeWithNationalInsuranceContribution(t *testing.T){
    employee := Employee{Id: "12345", Name: "John J Doe", AnnualGrossSalaryInCents: 906000}

    got := GenerateSalarySlipFor(employee)

    expected := SalarySlip{
                    GrossSalaryInCents: 75500,
                    EmployeeName: "John J Doe",
                    EmployeeId: "12345",
                    NationalInsuranceContributionsInCents: opt.Some[int64](12000),
                    Taxes: opt.None[Taxes](),
                }
    assertThat(got, expected, t)
}

func Test_shouldCalculateTheSalarySlipForAnEmployeeWithTaxes(t *testing.T){
    employee := Employee{Id: "12345", Name: "John J Doe", AnnualGrossSalaryInCents: 1200000}

    got := GenerateSalarySlipFor(employee)

    expected := SalarySlip{
                    GrossSalaryInCents: 100000,
                    EmployeeName: "John J Doe",
                    EmployeeId: "12345",
                    NationalInsuranceContributionsInCents: opt.Some[int64](47280),
                    Taxes: opt.Some[Taxes](Taxes{
                                               TaxFreeAllowanceInCents: 1100000,
                                               TaxableIncomeInCents: 100000,
                                               TaxPayableInCents: 20000,
                                           }),
                }
    assertThat(got, expected, t)
}
