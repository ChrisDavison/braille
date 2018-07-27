# Introduction

Go from formatted braille to english or vice-versa.

Formatted braille uses:

'O' - Raised portion
'.' - Lower portion

# Usage

```
braille -f braillefile.txt

braille -m message
```

## Braille to Text

Input of the text file *helloworld.txt*

```
O. O. O. O. O. .O O. O. O. OO
OO .O O. O. .O OO .O OO O. .O
.. .. O. O. O. .O O. O. O. ..
```

using command

```
braille -f helloworld.txt
```

results in
```
helloworld
```

## Text to Braille

Input of the message *chris davison*

```
braille -m "chris davison"
```

results in

```
OO O. O. .O .O .. OO O. O. .O .O O. OO 
.. OO OO O. O. .. .O .. O. O. O. .O .O 
.. .. O. .. O. .. .. .. OO .. O. O. O. 
```
