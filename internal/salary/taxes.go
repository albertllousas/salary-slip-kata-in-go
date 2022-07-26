package salary
import (
    opt "github.com/eatmoreapple/optional"
)

var basicRateThreshold int64 = 1100000
var higherRateThreshold int64 = 4300000
var basicRateTax = 0.20
var higherRateTax = 0.40

type Taxes struct {
     TaxFreeAllowanceInCents int64
     TaxableIncomeInCents int64
     TaxPayableInCents int64
}

func CreateTaxes(annualGrossSalary int64) opt.Option[Taxes] {
    if annualGrossSalary > basicRateThreshold  {
        return opt.Some[Taxes](Taxes{
            TaxFreeAllowanceInCents: basicRateThreshold,
            TaxableIncomeInCents: (annualGrossSalary - basicRateThreshold),
            TaxPayableInCents: amountToTax(annualGrossSalary),
        })
    }
    return opt.None[Taxes]()
}

func amountToTax(annualGrossSalary int64) int64 {
    amountToTaxBasicRate := int64(float64(annualGrossSalary - basicRateThreshold) * basicRateTax)
    if annualGrossSalary > higherRateThreshold {
        amountToTaxBasicRate = int64(float64(higherRateThreshold - basicRateThreshold) * basicRateTax)
    }
    amountToTaxHigherRate := int64(0)
    if annualGrossSalary > higherRateThreshold {
        amountToTaxHigherRate = int64(float64(annualGrossSalary - higherRateThreshold) * higherRateTax)
    }
    return amountToTaxBasicRate + amountToTaxHigherRate
}
