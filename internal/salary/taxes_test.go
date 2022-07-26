package salary

import (
    "testing"
    opt "github.com/eatmoreapple/optional"
)

func Test_shouldNotCreateTaxesForLowIncomes(t *testing.T){
    assertThat(CreateTaxes(10000), opt.None[Taxes](), t)
}

func Test_shouldCreateTaxesForBasicRateIncomes(t *testing.T){
   got := CreateTaxes(1200000)

   expected := opt.Some[Taxes](Taxes{
                  TaxFreeAllowanceInCents: 1100000,
                  TaxableIncomeInCents: 100000,
                  TaxPayableInCents: 20000,
              })
   assertThat(got, expected, t)
}

func Test_shouldCreateTaxesForHigherRateIncomes(t *testing.T){
   got := CreateTaxes(4500000)

   expected := opt.Some[Taxes](Taxes{
                  TaxFreeAllowanceInCents: 1100000,
                  TaxableIncomeInCents: 3400000,
                  TaxPayableInCents: 720000,
              })
   assertThat(got, expected, t)
}
