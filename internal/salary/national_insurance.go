package salary
import (
    opt "github.com/eatmoreapple/optional"
)

var basicThreshold int64 = 806000
var higherThreshold int64 = 4300000
var basicRatePercentage = 0.12
var higherRatePercentage = 0.02

func CalculateNationalInsurance(annualGrossSalary int64) opt.Option[int64] {
    if (annualGrossSalary - basicThreshold) > 0 {
        nationalInsuranceForBasicRate := int64(float64(annualGrossSalary - basicThreshold) * basicRatePercentage)
        nationalInsuranceForHigherRate := int64(0)
        if(annualGrossSalary - higherThreshold) > 0 {
            nationalInsuranceForBasicRate = int64(float64(higherThreshold - basicThreshold) * basicRatePercentage)
            nationalInsuranceForHigherRate = int64(float64(annualGrossSalary - higherThreshold) * higherRatePercentage)
        }
        return opt.Some[int64](nationalInsuranceForBasicRate + nationalInsuranceForHigherRate)
    }
    return opt.None[int64]()
}
