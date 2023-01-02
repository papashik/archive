#!/bin/bash

tree1(){
local DIR=$1
local NUM=0 # Номер текущей папки/файла - чтобы проверять, последняя ли

if [[ $KEY_LIST_DIR == true ]]; then
  local AMOUNT=$(ls -ld -- "$DIR"/*/ | wc -l) # Количество только папок
else
  local AMOUNT=$(ls "$DIR" | wc -l) # Количество файлов и папок вместе
fi

if [[ -d "$DIR" ]]; then
  let DIRNUM++
  let "V = $V + 1"
  while read NAME; do
    if [[ $KEY_LIST_DIR == true && -f "$DIR/$NAME" ]]; then
      continue
    fi
    let NUM++
    # echo -n "$NUM/$AMOUNT "
    if [[ $NUM -eq $AMOUNT ]]; then
      echo -e ${SPACES:0:$(($V*$S))}'`--' $NAME
    else
      echo -e ${SPACES:0:$(($V*$S))}'|--' $NAME
    fi
    tree1 "$DIR/$NAME" # Рекурсивный вызов
  done < <(ls "$DIR")
  let "V = $V - 1"
else
  let FILENUM++
fi
}

SPACES='| . | . | . | . | . | . | . | . | . | . | . | . '
# SPACES='|\t\b\b\b\b|\t\b\b\b\b|\t\b\b\b\b|\t\b\b\b\b'
S=4 # Размер одного пробела
let "V = -1" # Вложенность
DIRNUM=-1 # Общее количество директорий
FILENUM=0 # Общее количество файлов
KEY_LIST_DIR=false # Наличие ключа -d (true или false)
KEY_FILE_OUTPUT=false # Наличие ключа -o (true или false)
FILE_OUTPUT=1 # Путь к файлу для вывода
DIR="."

while [[ ! -z $1 ]]; do
  if [[ -d $1 ]]; then
    DIR=$1
    shift
  elif [[ $1 == '-d' ]]; then
    KEY_LIST_DIR=true
    shift
  elif [[ $1 == '-o' ]]; then
    KEY_FILE_OUTPUT=true
    
    if [[ -z $2 ]]; then
      echo "no file path to output"
      exit 0
    elif [[ -d $2 ]]; then
      echo "is a directory \"$2\""
      exit 0
    else
      FILE_OUTPUT=$2
      exec 1> $FILE_OUTPUT
      shift 2
    fi

  else
    echo "unknown argument"
    exit 0
  fi
done

exec 2> tree_errors.txt

echo $DIR
tree1 $DIR

if [[ $KEY_LIST_DIR == true ]]; then
  echo "$DIRNUM directories"
else
  echo "$DIRNUM directories, $FILENUM files"
fi

echo "KEY -d: $KEY_LIST_DIR, KEY -o: $KEY_FILE_OUTPUT"
