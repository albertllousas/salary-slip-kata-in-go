package salary

import (
    "testing"
    opt "github.com/eatmoreapple/optional"
)

func Test_shouldNotCalculateNationalInsuranceForLowIncomes(t *testing.T){
    assertThat(CalculateNationalInsurance(800000), opt.None[int64](), t)
}

func Test_shouldCalculateNationalInsuranceForBasicRateIncomes(t *testing.T){
    assertThat(CalculateNationalInsurance(906000), opt.Some[int64](12000), t)
}

func Test_shouldCalculateNationalInsuranceForHigherRateIncomes(t *testing.T){
    assertThat(CalculateNationalInsurance(4500000), opt.Some[int64](423280), t)
}
