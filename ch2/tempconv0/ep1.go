package tempconv0

import (

)

// 摄氏温度
type Cel float64
// 华氏温度
type Fah float64

const (
    AbsluteZeros Cel = -273.15
    FreezingC    Cel = 0
    BoilingC     Cel = 100
)

func CToF(c Cel) Fah {

}
