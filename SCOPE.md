# Escopo do Projeto: Filtro de Internet Doméstico (DNS Sinkhole)

## 1. Problema e Objetivo
Proteger os dispositivos conectados à rede doméstica contra o acesso a sites de conteúdo adulto (+18). O objetivo é criar um serviço em código que rode em uma máquina Linux 24h por dia e intercepte as requisições de resolução de nomes da rede.

## 2. O que É o Projeto (MVP)
- **Serviço DNS Customizado:** Um servidor DNS simples desenvolvido em código.
- **Mecanismo de Bloqueio (Blocklist):** Leitura de um arquivo de texto com domínios proibidos (`blocklist.txt`).
- **Resolução Condicional:**
  - Se o domínio solicitado estiver na lista negra: responder com endereço nulo (`0.0.0.0`), impedindo a conexão.
  - Se o domínio for legítimo: repassar a consulta (forwarding) para um servidor DNS público confiável (ex: Cloudflare `1.1.1.1` ou Google `8.8.8.8`) e devolver a resposta ao dispositivo solicitante.
- **Registro de Log Básico no Console:** Exibir no terminal quais domínios foram consultados e se foram liberados ou bloqueados.

## 3. Critérios de Pronto (Definition of Done)
1. O servidor inicia e escuta na porta UDP padrão de DNS (porta 53 ou porta alternativa para testes locais).
2. Uma requisição para um domínio liberado (ex: `wikipedia.org`) retorna o IP real do site.
3. Uma requisição para um domínio na lista (ex: `site-bloqueado.com`) retorna `0.0.0.0`.
4. Os testes funcionam via ferramentas de diagnóstico padrão (ex: `nslookup` ou `dig`).

## 4. O que NÃO será feito no MVP (Fora de Escopo Inicial)
- Interface gráfica web complexa (dashboard).
- Quebra/inspeção profunda de tráfego HTTPS via proxy.
- Categorização dinâmica por inteligência artificial em tempo real.
- Sistema de autenticação de usuários por dispositivo.
