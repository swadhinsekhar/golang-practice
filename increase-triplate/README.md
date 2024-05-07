


`int(^uint(0) >> 1)`
-  `uint(0)` creates an unsigned integer with value 0.
- `^uint(0)` flips all the bits of this 0 value, resulting in a value where all bits are 1. This is the maximum value that can be held by an unsigned integer.
- `>> 1` shifts the bits of this value one place to the right. This effectively divides the value by 2. The reason for this is that the maximum value of a signed integer is half that of an unsigned integer, because one bit is used to hold the sign.
`int(^uint(0) >> 1)` converts this value to an integer.