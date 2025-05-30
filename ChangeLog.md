# Changelog

## 1.3.2 (2025-06-01)
- Improved
  - Support for reference types (e.g., *int, *string) in To*(), From*() and Equal()
  - ToTime() to support *string, []byte, and *time.Time as input types
  - To() to support *time.Time as output type

## v1.3.1 (2025-05-29)
- Improved
  - Equal() to support time.Time comparison and allow symmetric comparison by swapping variables.

## v1.3.0 (2025-05-25)
- Added conversion functions
  - Equal() function to compare two values for equality

## v1.2.7 (2025-05-07)
- Improved
  - ToInt8(), ToInt16(), ToInt32(), ToInt64(), and ToInt()
  - ToUint8(), ToUint16(), ToUint32(), ToUint64(), and ToUint()
    - Enhanced string-to-integer conversion by adding support for parsing float strings
  - FromInt8(), FromInt16(), FromInt32(), and FromInt64()
  - FromUint8(), FromUint16(), FromUint32(), and FromUint64()
    - Enhanced conversion to check unsigned integer ranges
  - FromFloat64(), FromFloat32()
    - Added proper range checks for integer conversions

## v1.2.6 (2025-05-05)
- Improved
  - ToString()
    - Enhanced conversion to string with specific source types

## v1.2.5 (2023-09-06)
- Added
  - To() and From() for generic conversion

## v1.2.4 (2023-09-02)
- Added
  - ToBytes() and FromBytes()

## v1.2.3 (2023-09-14)
- Updated 
  - ToTime()
    - Enabled to set multiple timestamp layouts
    - Added major timestamp layouts as default

## v1.2.2 (2023-09-04)
- Updated
  - ToUint*() and ToInt*()
    - Supported conversion from float32 and float64 values
 
## v1.2.1 (2023-07-27)
- Added
  - ToTime()
  - Examples for To*() and From*() functions for godoc
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
