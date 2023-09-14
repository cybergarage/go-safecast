# Changelog

## v1.2.3 (2023-09-14)
- Updated ToTime() to enable setting multiple layouts

## v1.2.2 (2023-09-04)
- Updated ToUint*() and ToInt*() functions to support conversion from float32 and float64 values
 
## v1.2.1 (2023-07-27)
- Added
  - ToTime()
  - Examples for To*() and From*() functions
- Fixed
  -  string conversion bugs in ToInt*() and ToUint*()

## v1.2.0 (2023-05-20)
- Added 
  - ToInt(), ToInt64(), ToInt32(), ToInt16() and ToInt8()
  - ToUint(), ToUint64(), ToUint32(), ToUint16() and ToUint8()
  - ToFloat64() and ToFloat32()

## v1.1.0 (2022-11-20)
- Added
  - FromFloat64(), FromFloat32(), FromString(), and FromBool()
- Updated
  - FromInt64() and FromUint64() to support float32, float64, string and bool

## v1.0.0 (2022-11-19)
- Initial release  
- Supported
  - Signed integers
    - FromInt(), FromInt8(), FromInt16(), FromInt32(), FromInt64() 
  - Unsigned integers
    - FromUint(), FromUint8(), FromUint16(), FromUint32(), FromUint64() 
