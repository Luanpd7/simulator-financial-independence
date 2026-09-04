# 📈 Simulador de Aposentadoria — Back-end

API REST desenvolvida em **Go (Golang)** responsável pelos cálculos financeiros utilizados no Simulador de Aposentadoria.

A API recebe os dados financeiros informados pelo usuário, processa a simulação e retorna indicadores relacionados à evolução patrimonial e ao planejamento para aposentadoria.

- **[Ver publicação no LinkedIn](https://www.linkedin.com/posts/luan-pereira-dias-a327a0253_conclu%C3%AD-mais-um-projeto-pessoal-com-o-activity-7478299174644506624-a626?utm_source=share&utm_medium=member_desktop&rcm=ACoAAD6WM8YBZpZBcOFJu9kQqXyrUpV5Vaoz8g4)**

## 🚀 Tecnologias

- **Go (Golang):** utilizado para desenvolvimento da API e implementação dos cálculos financeiros.
- **Gin:** framework utilizado para criação do servidor HTTP e gerenciamento das rotas da API REST.
- **REST API:** utilizada para comunicação entre o back-end e o front-end Flutter.
- **Clean Architecture:** utilizada para separar as responsabilidades entre as camadas da aplicação.
- **Git / GitHub:** utilizados para controle de versão e gerenciamento do código.
- **JIRA:** Para organização nas demandas.

## 📋 Funcionalidades

A API é responsável por realizar os cálculos da simulação financeira a partir das informações fornecidas pelo usuário.

Entre os dados utilizados na simulação estão:

- Patrimônio atual
- Aporte mensal
- Taxa de juros anual
- Inflação
- Idade atual
- Idade desejada para aposentadoria
- Período da simulação

A partir desses dados, o back-end calcula:

- Patrimônio final
- Patrimônio final ajustado pela inflação
- Taxa de juros real anual
- Taxa de juros real mensal
- Total contribuído
- Total de rendimentos obtidos
- Tempo até a aposentadoria
- Evolução do patrimônio ao longo dos anos

## 🧮 Cálculos Financeiros

A API utiliza conceitos financeiros para realizar as projeções de longo prazo.

### Juros Compostos

Os juros compostos são utilizados para calcular a evolução dos investimentos ao longo do período da simulação.

```text
Patrimônio atual
      +
Aportes mensais
      +
Rentabilidade
      │
      ▼
Evolução do patrimônio
```

### Inflação

A inflação é utilizada para calcular o valor real do patrimônio futuro, permitindo comparar o patrimônio nominal com seu poder de compra.

```text
Patrimônio Futuro
      │
      ├── Valor Nominal
      │
      └── Valor Ajustado pela Inflação
```

### Taxa de Juros Real

A aplicação calcula a rentabilidade real do investimento considerando o efeito da inflação sobre a taxa de retorno.

Isso permite analisar o crescimento real do patrimônio ao longo do período da simulação.

## 🏗️ Arquitetura

O projeto utiliza **Clean Architecture** para separar as responsabilidades da aplicação.

A arquitetura busca manter separadas as camadas responsáveis por:

- Entidades e modelos de domínio
- Regras de negócio
- Acesso aos dados
- Tratamento das requisições HTTP
- Rotas da API

## 🔄 Fluxo da Simulação

```text
Flutter Web
     │
     │ REST API
     ▼
Go + Gin
     │
     ▼
Validação dos Dados
     │
     ▼
Regras de Negócio
     │
     ▼
Cálculos Financeiros
     │
     ├── Juros Compostos
     ├── Inflação
     ├── Juros Reais
     └── Evolução Patrimonial
     │
     ▼
Resultado da Simulação
     │
     ▼
Flutter Web
```

## 📊 Resultado da API

O back-end retorna os principais indicadores da simulação:

```text
Resultado
│
├── Patrimônio Final
├── Patrimônio Ajustado pela Inflação
├── Taxa Real Anual
├── Taxa Real Mensal
├── Total Contribuído
├── Total de Rendimentos
├── Tempo até a Aposentadoria
│
└── Evolução Patrimonial
    │
    ├── Ano
    ├── Idade
    ├── Patrimônio Nominal
    └── Patrimônio Ajustado
```

Os dados de evolução patrimonial retornados pela API são utilizados pelo front-end para gerar o gráfico da simulação.

## 💰 Conceitos Financeiros Aplicados

Durante o desenvolvimento foram aplicados conceitos relacionados ao mercado financeiro, incluindo:

- **Juros compostos**
- **Taxa de juros real**
- **Inflação**
- **Rentabilidade**
- **Aportes mensais**
- **Evolução patrimonial**
- **Planejamento financeiro de longo prazo**
- **Planejamento para aposentadoria**

## 🔗 Front-end

O front-end da aplicação foi desenvolvido utilizando **Flutter**.

Repositório:

https://lnkd.in/d_YrJkMw

## 🎯 Objetivo do Projeto

O projeto foi desenvolvido com o objetivo de aprimorar conhecimentos em desenvolvimento **Full Stack**, utilizando **Flutter e Go (Golang)** em uma aplicação completa.

O desenvolvimento do back-end permitiu colocar em prática conhecimentos como:

- Desenvolvimento de APIs REST com Go
- Utilização do framework Gin
- Clean Architecture
- Separação de responsabilidades
- Implementação de regras de negócio
- Integração entre Flutter e Go
- Processamento de cálculos financeiros

Além dos conhecimentos técnicos, o projeto permitiu aplicar conceitos do **mercado financeiro** em um cenário prático, utilizando programação para realizar projeções de patrimônio e planejamento para aposentadoria.
