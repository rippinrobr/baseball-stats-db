#!/bin/bash

# check for required params

#echo "Baseball Databank Directory: $1"
#echo "Language to create structures for: $2"
#echo "Ouput directory: $3"
usage() 
{
    echo "usage: $0 -i|--input <baseball databank files dir> -l|--language go|haskell"
    echo "usage:   [-o|--ouptut-dir <dir to create the files in>]"
}

csv=
db=
input_dir=
json=
language=
name=
output_dir=

#echo "about to do the while check"
while [ "$1" != "" ]; do 
    case $1 in 
        -i | --input )          shift
                                input_dir=$1
                                ;;
        -l | --language )       shift
                                language=$1
                                ;;
        -o | --output-dir )     shift
                                output_dir=$1
                                ;;
        --csv )                 csv=1
                                ;;
        --db )                  db=1
                                ;;
        --json )                json=1
                                ;;
        -h | --help )           usage
                                exit
                                ;;
        -n | --name )           shift 
                                name=$1
                                ;;
        -p | --package )        print_package=1
                                ;;
        * )                     usage
                                exit 1
                                ;;
    esac
    shift
done

echo "input: $input_dir"
echo "output: $output_dir"
echo "language: $language"

if [ "$language" == "" ] || [ "$input_dir" == "" ]
then
   echo ""
   echo "ERROR: language, input are both required options"
   echo ""
   usage
   exit 1
fi

if [ "$language" != "go" ] && [ "$language" != "haskell" ]
then 
   echo ""
   echo "ERROR: the language '$language' is not supported, only go and haskell are supported."
   exit 1
fi

optional_params=""
if [ "$csv" == "1" ]; then 
   optional_params=" --csv"
fi

if [ "$db" == "1" ]; then
   optional_params="$optional_params --db"
fi

if [ "$json" == "1" ]; then
   optional_params="$optional_params --json"
fi

if [ "$name" != "" ]; then
   optional_params="$optional_params --name $name"
fi

if [ "$print_package" == "1" ]; then 
   optional_params="$optional_params --package"
fi

if [ "$output_dir" != "" ]; then 
   optional_params="$optional_params --output-dir $output_dir"
fi

cd ./generators 
for file in $( ls $input_dir/*.csv )
do
    python gen-data-structures.py --language $language --input $file $optional_params
done
cd ..
