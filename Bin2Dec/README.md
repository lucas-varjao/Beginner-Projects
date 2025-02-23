# Bin2Dec

Uma aplicação simples em Go que converte números binários em decimais.

## Funcionalidades

- Converte números binários de até 8 dígitos para decimal
- Validação de entrada para garantir que apenas 0's e 1's são aceitos
- Processamento concorrente usando goroutines para melhor performance

## Como Usar

1. Execute o programa:

```bash
go run main.go
```

2. Digite um número binário (máximo 8 dígitos)
3. O programa mostrará o valor decimal equivalente
4. Digite 'q' para sair

## Exemplos

```
Entrada: 1010
Saída: 10

Entrada: 11111111
Saída: 255
```

## Tecnologias

- Go 1.23.5
- Testes unitários
- Concorrência com goroutines e channels

## Estrutura do Projeto

```
Bin2Dec/
  ├── main.go       # Código principal
  ├── main_test.go  # Testes unitários
  └── go.mod        # Dependências
```
