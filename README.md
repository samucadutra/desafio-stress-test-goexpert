# Desafio Stress Test GoExpert

Este projeto é uma ferramenta de linha de comando (CLI) para realizar testes de carga em serviços web. Ele permite configurar o número total de requisições e o nível de concorrência, gerando um relatório ao final do teste.

## Como criar a imagem Docker

1. Certifique-se de que o Docker está instalado e configurado em sua máquina.
2. Navegue até o diretório do projeto.
3. Execute o seguinte comando para criar a imagem Docker:

   ```bash
   docker build -t stress-test .
   ```

## Como executar a aplicação

Após criar a imagem Docker, você pode executar a aplicação com o seguinte comando:

```bash
docker run stress-test --url=<URL_DO_SERVIÇO> --requests=<TOTAL_DE_REQUISIÇÕES> --concurrency=<NÍVEL_DE_CONCORRÊNCIA>
```

ou apenas:

```bash
docker run stress-test
```

- Dessa maneira, a aplicação irá utilizar os valores padrão para `url`, `requests` e `concurrency`, que são `http://google.com`, `100` e `10`, respectivamente.

```bash

### Exemplos

1. Testar o serviço `http://google.com` com 1000 requisições e 10 chamadas simultâneas:

   ```bash
   docker run stress-test-goexpert --url=http://google.com --requests=1000 --concurrency=10
   ```

2. Testar o serviço `http://meusite.com` com 500 requisições e 5 chamadas simultâneas:

   ```bash
   docker run stress-test-goexpert --url=http://meusite.com --requests=500 --concurrency=5
   ```

## Relatório Gerado

Ao final do teste, a aplicação exibirá um relatório com as seguintes informações:

- Tempo total de execução.
- Total de requisições realizadas.
- Número de requisições com status HTTP 200.
- Distribuição de outros códigos de status HTTP (ex.: 404, 500, etc.).

## Observação

Certifique-se de que o serviço que você está testando está acessível e preparado para receber o número de requisições configurado.
