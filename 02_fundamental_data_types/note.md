# Data types

## Numeric types

### Integers
There are two types of integers, signed and unsigned.
1. Signed integers: These can store both positive and negative numbers. int, int8, int16, int32, int64. **int** refers to 64 bit integer in most modern computers.
2. Unsigned integers: These can hold only positive numbers (and zero) because they don't need a bit to track the plus/minus sign. uint, uint8, uint16, uint32, uint64. **uint** refers to 64 bit unsigned integer in most modern computers.

### Floating point
- These store decimal values.
- float32 and float64
- default type for decimal values is float64

## Textual types

### String
- These are sequence of characters.
- In Go, strings are immutable.

### rune
- This is an alias for int32.
- It can be used to store characters.
- As Go uses UTF-8, emojis can also be stored in a rune value.

## Boolean
- It is the simplest type with only true and false as values.