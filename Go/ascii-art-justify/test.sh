
echo "Test 00:"
go run ascii-art.go --align right something standard
echo "\nTest 01:"
go run . --align=right left standard
echo "\nTest 02:"
go run . --align=left right standard
echo "\nTest 03:"
go run . --align=center hello shadow
echo "\nTest 04:"
go run . --align=justify "1 Two 4" shadow
echo "\nTest 05:"
go run . --align=right 23/32 standard
echo "\nTest 06:"
go run . --align=right ABCabc123 thinkertoy
echo "\nTest 07:"
go run . --align=center "#$%&\"" thinkertoy
echo "\nTest 08:"
go run . --align=left "23Hello World\!" standard
echo "\nTest 09:"
go run . --align=justify "HELLO there HOW are YOU?\!" thinkertoy
echo "\nTest 10:"
go run . --align=right "a -> A b -> B c -> C" shadow
echo "\nTest 11:"
go run . --align=right abcd shadow
echo "\nTest 12:"
go run . --align=center ola standard