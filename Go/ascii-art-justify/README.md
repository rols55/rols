# ascii-art-reverse

### Collaborators: 
    Roland Lehes (rols55)
    Priit Tuvike (\_parker\_)
    Jaan.Ginzul (TERSIT)

FOR AUDIT USE sh test.sh
### How to use:
Type in termianl ```go run . [OPTION] [STRING] [BANNER]```
#### [OPTION] Put your flags here, see flags section for more information
#### [STRING] type your text here between ""
#### [BANNER] choose banner format out of standard, shadow, thinkertoy
#### Flags:
 * `--output=` *outputs program's output into a file*
 * `--reverse=` *specify file where ascii art is located and program will output to you content of the file into terminal*
 * `--justify=` *aligns input string, this flag has 4 possible values: left, right, center, justify*

**syntax of flags follows golang's flag package permitted flag syntax**

### Assignment
https://github.com/01-edu/public/tree/master/subjects/ascii-art/reverse


### Type `bash test.sh` into terminal to run audit tests
### Example of program:

```
user$ go run . hello standard
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___  
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               