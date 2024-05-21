# treinamento_go
Treinamento GO.

## Aplicação
Esta aplicação foi criada a partir do Getting Started do GO, com pequensa alterações para testes e melhor retorno dos dados.

## Como rodar a aplicação.
Para testar a aplicação, bastar fazer o build da aplicação, configurar o local de instalação e gerar o arquivo.

## Windows
Faça download da linguagem GO no site [Go Dev](https://go.dev/dl/)
<br>Baixe a última versão do Golang
<br>Faça a instação da mesma, é bem simples.
<br>Verifique o caminho onde os binários estão: ```go list -f '{{.Target}}'```
<br>Salve o caminho pois será necessário nos próximos passos.

### Set PATH
**No prompt de comando:**
<br> ```set PATH=%PATH%,C:\path\to\your\install\directory```
<br>Aqui você configura o caminho do diretório da sua instalação (Local onde você quer rodar sua aplicação.)

<br>```go install```
<br>Aqui fazemos a instalação no diretório informado.

Após isso, bastar digitar o nome do seu .exe que sua aplicação irá rodar.
