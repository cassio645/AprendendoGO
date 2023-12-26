#!/bin/bash

# EXECUTE: ./criarPasta.sh

nivel="nivel" # Adicionar numero do nivel ex 'nivel1' 'nivel5'


# Um for para a quantidade de exercicios daquele nivel
for i in {01..10}; do
    directory="exercicios/$nivel/ex$i"
    file="$directory/main.go"

    # Criando o diretório se não existir
    mkdir -p "$directory"

    # Criando o arquivo inicial GO dentro do diretório
    echo -e 'package main\n\nimport "fmt"\n\nfunc main() {\n\n}' > "$file"

    echo "Arquivo $file criado com sucesso!"
done
