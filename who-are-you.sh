returnValue=`curl -s https://support.labs.01-edu.org/assets/superhero/all.json`
echo $returnValue  | cat > list.json
jq '.[]  | select(.id==70) | .name' list.json