#!/bin.bash

echo test 1
go run . --output test00.txt banana standard
echo test 2
go run . --output=test00.txt "First\nTest" shadow
cat test00.txt
echo test 3
go run . --output=test01.txt "hello" standard
cat test01.txt
echo test 4
go run . --output=test02.txt "123 -> #$%" standard
cat test02.txt
echo test 5
go run . --output=test03.txt "432 -> #$%&@" shadow
cat test03.txt
echo test 6
go run . --output=test04.txt "There" shadow
cat test04.txt
echo test 7
go run . --output=test05.txt "123 -> \"#$%@" thinkertoy
cat test05.txt
echo test 8
go run . --output=test06.txt "2 you" thinkertoy
cat test06.txt
echo test 9 
go run . --output=test07.txt 'Testing long output!' standard
cat test07.txt
echo test 10
go run . --output=randomFileName.txt "raNdom StRing"
cat randomFileName.txt
echo test 11
go run . --output=sdad.txt "hdjska1829 dhsa21"
cat sdad.txt
echo test 12
go run . --output=special.txt ")(*&!@#|?><~"
cat special.txt
echo test 13
go run . --output=ultimate.txt "dsaj*(UASDjmk nd jk  jklj \n du8s9a \n ()*"
cat ultimate.txt