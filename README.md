# http-server-projeto-korp

Projeto desenvolvido como parte do desafio técnico da Korp.

## Objetivo

Implementar um serviço HTTP em Golang executado em containers Docker, utilizando NGINX como proxy reverso, Prometheus para monitoramento, Grafana para visualização de métricas e Ansible para automação do provisionamento.

---

## Tecnologias Utilizadas

* Golang
* Docker
* Docker Compose
* NGINX
* Prometheus
* Grafana
* Ansible

---

## Arquitetura

Cliente

↓

NGINX (porta 80)

↓

http-server-projeto-korp (porta 8080)

↓

Prometheus (coleta métricas)

↓

Grafana (visualização)

---

## Funcionalidades

### Endpoint principal

GET `/projeto-korp`

Exemplo de resposta:

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-06-08T11:29:12Z"
}
```

O campo `horario` é gerado dinamicamente em UTC a cada requisição.

### Endpoint de Health Check

GET `/health`

Resposta:

```json
{
  "status": "UP"
}
```

### Endpoint de Métricas

GET `/metrics`

Métricas expostas no padrão Prometheus.

---

## Métricas Implementadas

### Disponibilidade do Serviço

```promql
service_up
```

### Volume de Requisições

```promql
http_requests_total
```

---

## Estrutura do Projeto

```text
.
├── ansible
│   ├── inventory.ini
│   └── playbook.yml
├── grafana
│   └── dashboard.json
├── nginx
│   └── http-server-projeto-korp.conf
├── prometheus
│   └── prometheus.yml
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## Executando com Docker Compose

### Build e inicialização

```bash
docker compose up -d --build
```

### Verificar containers

```bash
docker ps
```

### Testar aplicação

```bash
curl http://localhost/projeto-korp
```

### Testar Health Check

```bash
curl http://localhost/health
```

### Testar métricas

```bash
curl http://localhost/metrics
```

---

## Monitoramento

### Prometheus

Acessar:

```text
http://localhost:9090
```

### Grafana

Acessar:

```text
http://localhost:3000
```

Login padrão:

```text
Usuário: admin
Senha: admin
```

Dashboard configurado para exibir:

* Disponibilidade do serviço
* Volume de requisições

---

## Automação com Ansible

Para provisionar o ambiente:

```bash
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml
```

O playbook realiza:

* Verificação do Docker
* Criação da rede Docker
* Build da aplicação
* Inicialização dos containers
* Configuração do proxy reverso
* Inicialização do Prometheus
* Inicialização do Grafana
* Validação do endpoint HTTP
* Exibição da resposta no terminal

---

## Teste Esperado

Após a execução do ambiente:

```bash
curl http://localhost/projeto-korp
```

Resposta esperada:

```json
{
  "nome": "Projeto Korp",
  "horario": "<horário UTC atual>"
}
```
