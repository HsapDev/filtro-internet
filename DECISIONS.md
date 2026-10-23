# Registro de Decisões de Arquitetura (ADR)

## ADR 001: Abordagem de Filtro de Rede (DNS Sinkhole vs Proxy HTTPS)

### Contexto
O usuário deseja bloquear sites adultos para todos os dispositivos de sua rede doméstica através de uma máquina Linux rodando continuamente, construindo a solução em código para aprender as engrenagens de rede.

### Opções Consideradas
1. **Filtro DNS Customizado (DNS Sinkhole):** Interceptar requisições DNS na rede e responder `0.0.0.0` para domínios listados como proibidos.
2. **Servidor Proxy HTTP/HTTPS:** Interceptar as requisições web diretamente e analisar as URLs solicitadas.

### Decisão Tomada
**Opção 1 (Filtro DNS Customizado).**

### Justificativa
- A web moderna opera quase 100% sobre HTTPS (criptografado). Um proxy exigiria instalação manual de certificados raiz em cada dispositivo da casa (smartphones, TVs, computadores) para evitar alertas graves de segurança e certificado inválido.
- A abordagem DNS atua no nível do domínio (ex: `dominio.com`), sendo suficiente para impedir o carregamento total dos sites.
- Funciona de maneira uniforme para qualquer dispositivo na rede Wi-Fi via configuração centralizada de DNS no roteador.
- É ideal para fins educacionais, permitindo compreender a estrutura de pacotes UDP e o protocolo DNS na prática.

---

## ADR 002: Escolha da Linguagem de Programação (Go / Golang)

### Contexto
Necessidade de escolher a tecnologia para implementar o serviço de DNS Sinkhole, considerando aprendizado do usuário, facilidade de compilação para Linux e performance/consumo de recursos em execução 24/7.

### Opções Consideradas
1. **Python:** Familiar, porém requer interpretador e gerenciamento de dependências no ambiente final.
2. **Go (Golang):** Compilação para binário único estático, baixíssimo consumo de memória, padrão de mercado em infraestrutura e redes.
3. **Node.js:** Ecossistema web conhecido, mas maior pegada de memória e dependências externas.

### Decisão Tomada
**Go (Golang).**

### Justificativa
- Facilidade de deploy: gera um único executável estático que pode ser copiado e executado diretamente no Linux sem precisar instalar nada lá.
- Excelente suporte nativo a redes no pacote padrão (`net`).
- Linguagem simples, com poucas palavras-chave e tipagem explícita, ideal para o aprendizado prático de redes.

---

## ADR 003: Adoção da Biblioteca `github.com/miekg/dns`

### Contexto
Para processar e responder requisições DNS sobre UDP/TCP, é necessário escolher entre implementar o parser do protocolo binário RFC 1035 manualmente ou utilizar uma biblioteca especializada da comunidade Go.

### Opções Consideradas
1. **Parser de Pacotes Manual (Raw Sockets):** Decodificação bit a bit dos datagramas UDP. Alto risco de bugs, complexidade desnecessária para o MVP e perda de foco na lógica de filtragem.
2. **Biblioteca `github.com/miekg/dns`:** A biblioteca de DNS padrão e mais consolidada do ecossistema Go.

### Decisão Tomada
**Opção 2 (`github.com/miekg/dns`).**

### Justificativa
- **Confiabilidade e Segurança:** É a base do CoreDNS (servidor DNS oficial interno do Kubernetes) e adotada globalmente por empresas como Cloudflare e Google Cloud.
- **Foco Didático:** Elimina a complexidade de baixo nível de manipulação de bytes e permite ao usuário focar na lógica central: roteamento, consulta na blocklist e encaminhamento.
- **Pure Go:** Escrita 100% em Go nativo, sem dependências externas de C (CGO), mantendo a compilação cruzada e binário único estático para Linux.


