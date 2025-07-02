#!/bin/bash
echo "Привет. Имя: "
read name
mkdir "$name"
touch "$name"/welcome.txt
echo "Привет, "$name"!  Это твоя первая папка" > $name/welcome.txt
