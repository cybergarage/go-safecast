# Changelog

## v1.3.6 (Unreleased)
- Fixed
  - `ToUint8()` now rejects overflowing `int16` and `int32` inputs, and `ToUint16()` rejects overflowing `int32` inputs, including pointer inputs and generic `To()` calls.
  - `FromFloat64()` rejects the exact exclusive upper bounds of 64-bit integer destinations (`2^63` for signed integers and `2^64` for unsigned integers), which previously passed rounded maximum checks. The fix also applies to `FromFloat32()` and generic `From()` calls.
  - `Compare()` preserves the original operand order when retrying a failed conversion with reversed operands.
- Improved
  - Added regression tests for unsigned narrowing, floating-point integer boundaries, and mixed-width comparison ordering.
  - Expanded tests for generic conversion, comparison, equality, and byte conversion functions.
  - Updated GitHub Actions and golangci-lint configuration.
- Compatibility
  - Public function signatures are unchanged. Previously accepted overflowing inputs now return errors, and affected fallback comparisons return the corrected ordering.

## v1.3.5 (2025-11-23)
- Improved
  - Equal() supports slice ([]any) and map (map[string]any, map[any]any) comparisons
- Fixed 
  - ToUint() and ToUint64() string conversion to support uint64 maximum value (18446744073709551615)
    - Replaced strconv.Atoi with strconv.ParseUint in fromString helper functions
  - golangci-lint warnings
  
## v1.3.4 (2025-07-08)
- Improved
  - Examples
    - Added examples for To*() and From*() functions in godoc
  - Testing
    - Added comprehensive tests for From() and To() generic functions
    - Extended test coverage for Compare(), FromFloat64(), FromString(), ToString() functions
    - Added edge case testing for special values (NaN, Inf, nil, pointers)
    - Added symmetric comparison tests and error condition coverage

## v1.3.3 (2025-06-18)
- Added conversion functions
  - Compare() function to compare two values for equality and return a integer result

## 1.3.2 (2025-05-30)
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
