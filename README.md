Simple library for converting ints (as *big.Int) to base62 encoded string.

Why base62? Sometimes you just want an alphanumeric representation of a number, and base64's extra non alphanumeric chars ("-" and "+") just won't do.

base62.EncodeBig() - Takes a pointer to big.Int and returns base62 encoded string.
base62.EncodeStr() - Takes a string representation of an integer (like "150") and returns a base62 encoded string.
