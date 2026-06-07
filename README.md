# Projeto Korp

Desafio técnico desenvolvido para o processo seletivo da Korp.

## Objetivo

Desenvolver um serviço HTTP em Golang executando em containers Docker, utilizando NGINX como proxy reverso, Prometheus para coleta de métricas, Grafana para observabilidade e Ansible para automação da infraestrutura.

---

## Tecnologias Utilizadas

- Golang
- Docker
- Docker Compose
- NGINX
- Prometheus
- Grafana
- Ansible

---

## Estrutura do Projeto

```text
http-server-projeto-korp/
│
├── ansible/
├── grafana/
├── nginx/
├── prometheus/
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## Serviço HTTP

Endpoint:

```http
GET /projeto-korp
```

Resposta:

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-06-07T04:00:00Z"
}
```

O horário é gerado dinamicamente em UTC a cada requisição.

---

## Arquitetura

```text
Cliente
   │
   ▼
NGINX (porta 80)
   │
   ▼
http-server-projeto-korp (porta 8080)
   │
   ├── /projeto-korp
   └── /metrics
          │
          ▼
     Prometheus
          │
          ▼
       Grafana
```

---

## Como Executar

### Build da aplicação

```bash
docker build -t http-server-projeto-korp .
```

### Subir ambiente completo

```bash
docker compose up --build
```

---

## Testes

Consultar API:

```bash
curl http://localhost/projeto-korp
```

Consultar métricas:

```bash
curl http://localhost:8080/metrics
```

---

## Monitoramento

### Prometheus

Acesso:

```text
http://localhost:9090
```

Métrica implementada:

```text
http_requests_total
```

Responsável por contabilizar todas as requisições recebidas pelo endpoint da aplicação.

---

### Grafana

Acesso:

```text
http://localhost:3000
```

Credenciais padrão:

```text
Usuário: admin
Senha: admin
```

Dashboard configurado para exibir:

- Total de requisições
- Disponibilidade do serviço

---

## Containers

### http-server-projeto-korp

Serviço desenvolvido em Golang responsável por responder as requisições HTTP.

### nginx

Proxy reverso responsável por encaminhar as requisições da porta 80 para o serviço Golang na porta 8080.

### prometheus

Coleta métricas expostas pela aplicação.

### grafana

Visualização das métricas e dashboards.

---

## Automação com Ansible

O ambiente pode ser provisionado automaticamente utilizando Ansible.

Exemplo:

```bash
ansible-playbook playbook.yml
```

O playbook realiza:

- Instalação do Docker
- Criação da rede Docker
- Build da imagem
- Inicialização dos containers
- Configuração do NGINX
- Configuração do Prometheus
- Validação da aplicação

---

## Autor

Ivana Miranda
