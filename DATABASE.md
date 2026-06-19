# 📊 Simulator API - Integração com Banco de Dados

## ✅ Estrutura Criada

### 1. **Banco de Dados**
- Arquivo: `internal/database/db.go`
- Banco SQLite: `simulator.db`
- Tabela: `simulations` com 18 colunas

### 2. **Camada de Dados**
- Arquivo: `data/repository/simulation_repository_impl.go`
- Funções:
  - `SaveSimulation()` - Salva nova simulação
  - `GetSimulationByID()` - Busca por ID
  - `GetAllSimulations()` - Lista todas com limite

### 3. **Handlers**
- `simulation_handler.go` - POST /simulation (com persistência)
- `query_handler.go` - GET /simulations, GET /simulations/:id

### 4. **Rotas**
- POST `/simulation` - Calcula e salva
- GET `/simulations` - Lista todas (?limit=100)
- GET `/simulations/:id` - Busca por ID

## 📋 Estrutura da Tabela `simulations`

| Campo | Tipo | Descrição |
|-------|------|-----------|
| id | TEXT | UUID único |
| created_at | TIMESTAMP | Quando foi criada |
| updated_at | TIMESTAMP | Última atualização |
| current_assets | REAL | Ativos atuais |
| monthly_contribution | REAL | Contribuição mensal |
| annual_percentage | REAL | % anual |
| current_age | INTEGER | Idade atual |
| retirement_age | INTEGER | Idade de aposentadoria |
| time_in_years | INTEGER | Tempo em anos |
| inflation | REAL | Taxa de inflação |
| target_amount | REAL | Valor alvo |
| years_to_retirement | INTEGER | Anos até aposentadoria |
| final_amount | REAL | Valor final |
| monthly_income | REAL | Renda mensal |
| status | TEXT | success/error |
| error_message | TEXT | Mensagem de erro |
| user_ip | TEXT | IP do cliente |
| request_duration_ms | INTEGER | Duração em ms |

## 🚀 Como Usar

### 1. Build
```bash
go build -o simulator-api.exe
```

### 2. Run
```bash
./simulator-api.exe
```

### 3. Criar Simulação
```bash
curl -X POST http://localhost:8080/simulation \
  -H "Content-Type: application/json" \
  -d '{
    "currentAssets": 100000,
    "monthlyContribution": 2000,
    "annualPercentage": 0.08,
    "currentAge": 30,
    "retirementAge": 65,
    "timeInYears": 35,
    "inflation": 0.03
  }'
```

### 4. Listar Simulações
```bash
curl http://localhost:8080/simulations?limit=10
```

### 5. Buscar por ID
```bash
curl http://localhost:8080/simulations/{id}
```

## 📦 Dependências Adicionadas
- `github.com/mattn/go-sqlite3` - Driver SQLite
- `github.com/google/uuid` - Geração de UUIDs

## 📁 Arquivos Criados/Modificados
- ✅ `internal/database/db.go` (novo)
- ✅ `internal/handlers/query_handler.go` (novo)
- ✅ `main.go` (modificado - inicializa DB)
- ✅ `data/repository/simulation_repository_impl.go` (modificado)
- ✅ `internal/handlers/simulation_handler.go` (modificado)
- ✅ `internal/routes/routes.go` (modificado)
- ✅ `go.mod` (modificado)
