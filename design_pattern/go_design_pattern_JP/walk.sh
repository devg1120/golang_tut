

dir_path="./*"
#dirs=`find $dir_path -maxdepth 0 -type f -name *.jpg`
dirs=`find $dir_path -maxdepth 0 -type d`

for dir in $dirs;
do
    name=$(echo "$dir" | cut -c 3-)
    echo  ""
    #echo [ $name ]
    echo -e "\033[32m[ $name ]\033[0m"
    cd $dir
    if [ -f ./go.mod ]; then
          #echo  exist go.mod
	  echo -n ""
    else
          echo $name
          echo  "     " not exist go.mod
          go mod init example.com/${name}
    fi
    go test
    retval=$?
    if [ $retval -eq 0 ]
    then
        echo "...Success"
    else
        echo -e "\033[31m...Error\033[0m"
    fi
    cd ..
done

#echo -e "\033[31mThis text is red\033[0m"
#echo -e "\033[32mGreen text\033[0m"
#echo -e "\033[34mBlue text\033[0m"
#echo -e "\033[33mYellow text\033[0m"
