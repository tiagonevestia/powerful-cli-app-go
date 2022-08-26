# Markdown Preview(mdp)

Uma das tarefas mais frequentes que precisamos executar em uma ferramenta de linha de comando é trabalhar com arquivos. Programas extraem dados de arquivos e salvam resultados neles. Particularmente de suma importância é quando trabalhamos com Linux ou Unix porque os recursos do sistema são representados como arquivos. Desenvolvemos um app CLI para se sentir confortável com manipulação de arquivos.

## Sobre

Vamos desenvolver uma ferramenta para visualizar arquivos Markdown localmente, usando um navegador da web. **Markdown** é uma linguagem de marcação leve que usa texto simples com sintaxe especial para representar a formatação compatível com HTML.

Em resumo, a ferramenta converte a fonte do Markdown em HTML que pode ser visualizado em um navegador.

### (V0)

A v0 da ferramenta aceitará o nome do arquivo em Markdown a ser visualizado como seu argumento. Com isso, ela executará quatro etapas principais:

1- Ler o conteúdo do arquivo Markdown de entrada.

2- Converter o conteúdo do arquivo Markdown para HTML(utilizando algumas lib externas).

3- Envolver os resultados com um cabeçalho e rodapé HTML.

4- Salvar o buffer em um arquivo HTML.

### (V1)

A v1 da ferramenta irá criar e usar arquivos temporários em vez de arquivos locais, para evitar o conflito de arquivos sendo visualizados simultaneamente.

1 - Criar um arquivos temporários em vez de arquivos locais.

2 - Vamos utilizar o pacote **ioutil** que oferece uma função de TempFile para criar arquivos temporários com nome aleatório.

3 - Permitindo que sejam executados simultaneamente com segurança.
