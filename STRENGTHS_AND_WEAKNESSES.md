# Pontos Fortes e Oportunidades de Melhoria (Strengths & Weaknesses)

## 1. Pontos Fortes (Strengths)
- **Carga em Memória com Tabela Hash (`map[string]bool`):** O uso de `map` no Go garante busca instantânea $O(1)$. Não importa se a lista tem 10 ou 100.000 sites, a verificação ocorre em poucos nanossegundos.
- **Tratamento de Recursos com `defer arquivo.Close()`:** Garante que o descritor de arquivo seja fechado adequadamente, evitando vazamento de memória ou locks no sistema operacional.
- **Leitura em Streaming com `bufio.Scanner`:** Lê linha por linha em vez de carregar todo o arquivo de uma só vez na memória, tornando o consumo de RAM mínimo.
- **Resolução Híbrida Inteligente:** Devolve `0.0.0.0` com TTL customizado imediatamente para requisições proibidas e atua como resolver recursivo encaminhando requisições legítimas para a Cloudflare (`1.1.1.1`).
- **Resiliência e Zero Dependências Nativas:** O projeto compila para um binário único autônomo (sem necessidade de instalar Go, Python ou bibliotecas na máquina Linux de destino).

## 2. Oportunidades de Melhoria / Dívidas Técnicas (Weaknesses)
- **Bloqueio de Subdomínios:** Atualmente, se o usuário bloquear `exemplo.com`, uma consulta para `sub.exemplo.com` pode não ser detectada se não estiver explicitada na lista. Podemos implementar verificação hierárquica de sufixos.
- **Normalização de Domínios:** Se o usuário digitar domínios com maiúsculas ou espaços no `blocklist.txt`, o teste exato de igualdade pode falhar. Futuramente podemos limpar espaços e converter para minúsculas (`strings.ToLower` e `strings.TrimSpace`).
- **Recarga a Quente (Hot Reload):** Se um novo site for adicionado ao `blocklist.txt`, atualmente é necessário reiniciar o programa para que ele seja carregado.

## 3. Lições Aprendidas (Lessons Learned)
- **Comportamento do `nslookup` no Windows:** O utilitário nativo ignora portas customizadas passadas via CLI direta, exigindo clientes dedicados como nosso `testar.go` ou ferramentas como `dig` para testes em portas não padrão.
- **Formatação de Registros de Recursos (RR):** Respostas DNS exigem cabeçalhos de autoridade e tipos de registros condizentes com a query para evitar rejeições silenciosas por clientes de rede.

