#these are all the tests from the audit page

echo  test 1 

go run . "hello" | cat -e

echo  test 2 

go run . "HELLO" | cat -e

echo  test 3 

go run . "HeLlo HuMaN" | cat -e

echo  test 4 

go run . "1Hello 2There" | cat -e

echo  test 5 

go run . "Hello\nThere" | cat -e

echo  test 6 

go run . "Hello\n\nThere" | cat -e

echo  test 7 

go run . "{Hello & There #}" | cat -e

echo  test 8 

go run . "hello There 1 to 2\!" | cat -e

echo  test 9

go run . "MaD3IrA&LiSboN" | cat -e

echo  test 10

go run . "1a\"#FdwHywR&/()=" | cat -e

echo  test 11 

go run . "{|}~" | cat -e

echo  test 12 

go run . "[\]^_ 'a" | cat -e

echo  test 13 

go run . "RGB" | cat -e

echo  test 14 

go run . ":;<=>?@" | cat -e

echo  test 15 

go run . "\\!\" #$%&'()*+,-./" | cat -e

echo  test 16 

go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e

echo  test 16 

go run . "abcdefghijklmnopqrstuvwxyz" | cat -e

echo  test 17 

go run . "THIS is a random STRING" | cat -e

echo  test 18 

go run . "random STRING 55" | cat -e

