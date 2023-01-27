# Challenge: Ganho de Capital

Programa de linha de comando (CLI) que calcula o imposto a
ser pago sobre lucros ou prejuízos de operações no mercado financeiro de ações

## Tecnologia

Projeto escrito em [Go](https://go.dev/) na versão 1.18

## Arquitetura

A arquitetura do projeto foi baseada no layout do repositório [golang-standards](https://github.com/golang-standards/project-layout)

## Requisitos para compilação

 - [Go](https://go.dev/doc/install) versão 1.18+

## Executando localmente

Com o Go instalado localmente, basta na raiz do repositório, executar o comando abaixo:

```
go run cmd/main.go
```

A aplicação irá aguardar o input ser passado, como no exemplo seguinte:

```
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000}]
[{"operation":"sell", "unit-cost":10.00, "quantity": 10000}]
```

Também é possível indicar um arquivo:

```
go run cmd/main.go < test/cases/case-1.txt
```

## Testes unitários

Para executar os testes unitários utilize o comando a seguir:

```
make unit-tests
```

## Testes de integração

Para executar os testes unitários utilize o comando a seguir:

```
make integration-tests
```
