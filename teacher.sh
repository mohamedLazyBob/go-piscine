key=$(grep "Church" -r interviews -l | cut -d "-" -f2)

#   Same that the previous one
key1=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $key
cat interviews/interview-"$key"
echo $MAIN_SUSPECT