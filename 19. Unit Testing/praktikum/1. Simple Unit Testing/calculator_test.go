
package calculator

import (
  "testing"
  // "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
  "fmt"

)

func TestAddition(t *testing.T) { 
    t.Run("Non Decimal", func(t *testing.T){
      result := Addition(1, 2)
      require.Equal(t, float64(3), result, "harus 3")
    })

    t.Run("Decimal", func (t *testing.T){
      result := Addition(1.2, 2.3)
      require.Equal(t, float64(3.5), result, "harus 3")
    })
}

func TestSubtraction(t *testing.T){
  t.Run("Non Decimal", func(t *testing.T){
    result := Subtraction(3, -2)
    require.Equal(t, float64(5), result, "harus 1")
  })

  t.Run("Decimal", func(t *testing.T){
    result := Subtraction(3.5, -2)
    require.Equal(t, float64(5.5), result, "harus 1")
  })
}

func TestDivision(t *testing.T){
  t.Run("Non Decimal", func(t *testing.T){
    result := Division(10, 2)
    require.Equal(t, float64(5), result, "harus 5")
  })

  t.Run("Decimal", func(t *testing.T){
    result := Division(7, 2)
    require.Equal(t, float64(3.5), result, "harus 1")
  })
}

func TestMultiplication(t *testing.T){
  t.Run("Non Decimal", func(t *testing.T){
    result := Multiplication(5, 6)
    require.Equal(t, float64(30), result, "harus 5")
  })

  t.Run("Decimal", func(t *testing.T) {
    result := Multiplication(2.6, 9)
    formattedResult := fmt.Sprintf("%.1f", result) // Format hasil dengan satu angka desimal
    require.Equal(t, "23.4", formattedResult, "Harus memiliki 1 angka desimal")
  })
}



