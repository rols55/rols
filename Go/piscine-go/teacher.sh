var1=$(sed -n "179p" streets/Buckingham_Place | cut -d "#" -f2)
echo $var1
cat interviews/interview-"$var1"
echo $MAIN_SUSPECT