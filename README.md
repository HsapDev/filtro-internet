# Filtro de Internet Doméstico (DNS Sinkhole em Go)

Um servidor DNS leve e de alta performance desenvolvido em Go para bloquear o acesso a sites de conteúdo adulto (+18) na rede local.

---

## 🚀 Como Funciona
Toda vez que um dispositivo da rede tenta acessar a internet:
1. Ele faz uma requisição DNS para este servidor.
2. O servidor consulta a lista de bloqueios em memória (`blocklist.txt`):
   - **Se o domínio estiver na lista:** O servidor responde instantaneamente com o IP nulo `0.0.0.0`, impedindo a conexão.
   - **Se o domínio for liberado:** O servidor repassa a requisição para a Cloudflare (`1.1.1.1:53`) e devolve o endereço IP real.

---

## 📁 Estrutura do Projeto
- `main.go`: Código do servidor DNS, carregamento da blocklist e regras de roteamento.
- `blocklist.txt`: Lista de domínios proibidos (um por linha).
- `testar.go`: Utilitário de linha de comando para testar consultas locais.
- `go.mod` / `go.sum`: Gerenciamento de dependências (`github.com/miekg/dns`).

---

## 🛠️ Como Executar Localmente (Ambiente de Testes)

1. **Iniciar o Servidor:**
   ```powershell
   go run main.go
   ```

2. **Testar com o Cliente de Testes:**
   * Testando um domínio bloqueado:
     ```powershell
     go run testar.go xvideos.com
     ```
   * Testando um domínio liberado:
     ```powershell
     go run testar.go google.com
     ```

---

## 🐧 Como Gerar o Executável para o Linux (Deploy 24/7)
O Go permite gerar o executável final para o seu Linux diretamente do seu Windows:

```powershell
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o filtro-dns main.go
```

Isso gera um arquivo único binário chamado `filtro-dns`. Basta copiá-lo para a sua máquina Linux junto com o arquivo `blocklist.txt` e executar como serviço de segundo plano na porta padrão `53`.
