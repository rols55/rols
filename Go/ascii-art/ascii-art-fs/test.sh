#these are all the tests from the audit page

echo  test 1 

go run . "hello" | cat -e

echo  test 2 

go run . "hello world" shadow | cat -e

echo  test 3 

go run . "nice 2 meet you" thinkertoy | cat -e

echo  test 4 

go run . "you & me" standard | cat -e

echo  test 5 

go run . "123" shadow | cat -e

echo  test 6 

go run . "/(\")" thinkertoy | cat -e

echo  test 7 

go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" shadow | cat -e

echo  test 8 

go run . "\"#$%&/()*+,-./" thinkertoy | cat -e

echo  test 9

go run . "It's Working" thinkertoy | cat -e

echo  test 10

go run . "RaNdOm StRiNg" standard | cat -e

echo  test 11 

go run . "random 123" standard | cat -e

echo  test 12 

go run . "[\]^_ 'a123 AbSd" standard | cat -e