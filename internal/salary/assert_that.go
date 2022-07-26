package salary

import (
    "testing"
    "reflect"
)

func assertThat(got any, expected any, t *testing.T) {
    if !reflect.DeepEqual(got, expected) {
        t.Errorf("expected %+v; got %+v", expected, got)
    }
}
