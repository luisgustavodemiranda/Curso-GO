---
name: criar-modulo-go
description: "Cria um módulo completo de estudo no Curso-GO a partir de imagem ou lista de aulas: pastas numeradas, exemplos Go implementados, arquivos auxiliares, READMEs, validação e skill específica. Use quando o estudante pedir um novo módulo; para dúvidas de uma aula existente, use a skill do módulo."
---

# Criar um módulo de estudo Go

## Resultado solicitado

Neste repositório, criar um módulo significa entregar exemplos completos para estudar, não apenas pastas ou programas que imprimem o título da aula. Execute o fluxo inteiro quando o estudante enviar a imagem ou lista e pedir a criação. Se ele pedir explicitamente apenas uma estrutura inicial, respeite esse escopo.

As instruções abaixo são uma esteira conduzida pelo assistente na conversa; não dependem de serviço em segundo plano, OCR externo ou publicação.

## 1. Interpretar a entrada

- Leia visualmente a imagem anexada e extraia título, nomes e ordem das aulas. Ignore duração, progresso e estado de reprodução.
- Se houver várias imagens, una as partes na ordem e elimine somente repetições claramente causadas por sobreposição.
- Distinga aulas de código de links, apostilas e itens de documentação. Preserve esses itens no índice, mas não crie um programa artificial para eles.
- A imagem fornece o roteiro, não o conteúdo dos vídeos. Crie exemplos didáticos próprios coerentes com os títulos; não afirme reproduzir o código do professor.
- Informe brevemente o módulo e a quantidade de aulas identificada e continue. Não exija aprovação desse inventário quando estiver legível.
- Se faltar um título essencial ou a imagem estiver cortada/ilegível, peça somente a informação faltante. Continue o trabalho independente; não invente aulas escondidas.

## 2. Localizar e numerar

Leia o README da raiz, liste os diretórios de módulos e confira o estado do Git antes de editar. Use os módulos existentes como referência de formato, sem copiar seu conteúdo temático.

Use o próximo número após o maior prefixo existente no formato `NN - Nome`, a menos que o usuário determine outro. Numere as aulas dentro dele a partir de 01, mantendo a ordem da entrada. Preserve nomes em português e normalize apenas caracteres incompatíveis com caminhos Windows, espaços duplicados e espaços/pontos finais.

Confira se o nome ou destino já existe. Em uma retomada, compare os arquivos e complete o que falta sem duplicar o módulo nem sobrescrever alterações do estudante. Se o pedido puder significar substituir um módulo existente ou criar outro, esclareça essa diferença antes da substituição.

## 3. Criar exemplos completos

Para cada aula de código, implemente o conceito anunciado com o menor exemplo útil e idiomático. Inclua funções auxiliares, dados de exemplo, templates, HTML, configurações e demais arquivos necessários à execução.

- Preserve a progressão de assuntos. Uma aula avançada pode partir da anterior, mas seu diretório deve conter o necessário para ser estudada independentemente.
- Prefira a biblioteca padrão. Quando o tema exigir dependência, crie `go.mod` na aula, escolha uma versão explícita, gere `go.sum` quando aplicável e documente o download.
- Um arquivo independente pode usar `go run main.go`. Se houver vários arquivos do pacote, use um módulo de aula com `go run .` ou documente todos os arquivos explicitamente. Não crie um módulo na raiz do curso sem necessidade.
- Não entregue TODOs, funções vazias ou mensagens de identificação como implementação final do assunto.
- Trate erros e libere recursos. Mantenha comentários em português focados no conceito.
- Para HTTP, prefira exemplos locais reproduzíveis, documentando a diferença entre simulação e integração real. Se a integração for essencial ao tema, implemente e documente também o modo real.
- Servidores de estudo devem ouvir localmente. Documente porta, rotas, testes e encerramento. Para clientes, use limites de tempo; não deixe validações penduradas.
- Para banco de dados, mensageria ou outros serviços, inclua configuração e instruções necessárias, sem credenciais reais. Não substitua silenciosamente o assunto por uma simulação. Se o ambiente não permitir executar, registre a limitação sem anunciar validação completa.

## 4. Documentar para estudar

Crie esta organização, ajustando arquivos auxiliares ao tema:

```text
NN - Nome do módulo/
  README.md
  00 - Documentos/
    README.md
  01 - Primeira aula/
    main.go
    README.md
  02 - Próxima aula/
    ...
```

Cada README de aula deve explicar objetivo, funcionamento do código, requisitos, comando exato a partir da pasta correta, saída ou resposta esperada e um desafio adicional com critério de conferência. Informe efeitos observáveis: arquivos gerados, dados gravados, portas abertas, internet ou serviços necessários.

O índice do módulo deve ligar todas as aulas, indicar quais são apenas documentação e mostrar como começar. Em `00 - Documentos`, relacione materiais fornecidos e referências oficiais verificadas; não invente apostilas nem duplique PDFs existentes sem motivo.

Atualize o README da raiz com o módulo e seus requisitos. Não substitua instruções específicas dos módulos anteriores.

## 5. Criar a skill do módulo

Crie `.github/skills/nome-do-modulo-go/SKILL.md`, seguindo a convenção local. Use nome único, curto, sem acentos, em minúsculas e com hífens. Se já existir, atualize o arquivo após lê-lo.

O frontmatter deve conter `name` e uma `description` que identifique o módulo e os pedidos de estudo que ela atende. O corpo deve registrar:

- caminho e índice do módulo;
- convenções reais de execução e validação;
- particularidades dos exemplos, serviços e recursos;
- orientação para explicar, depurar e ampliar uma aula sem refatorar as demais;
- necessidade de manter código, resultado esperado e índice sincronizados.

Escreva instruções específicas ao tema implementado. Não copie regras de HTTP para um módulo sem HTTP, nem replique todo este fluxo de criação. Vincule a skill no README do módulo.

## 6. Validar e corrigir

Formate os arquivos com gofmt e execute os exemplos finitos, comparando o resultado real com o README. Para servidores, faça uma requisição local ou teste o handler e encerre qualquer processo que tenha iniciado.

Use testes de comportamento quando trouxerem confiança relevante: casos de erro, serialização, rotas ou concorrência. Não crie testes apenas para conferir a existência de arquivos ou repetir a implementação. Execute comandos de módulo no diretório que contém seu go.mod.

Confira se todos os itens legíveis da entrada estão no índice e têm o material correspondente. Verifique links locais, arquivos auxiliares, comandos documentados, formatação e diff. Valide o frontmatter e o conteúdo da skill; use um validador de skills disponível e, se ele estiver indisponível, registre a checagem alternativa.

Falha de download, serviço ausente ou permissão do ambiente não equivale a erro de código. Resolva o que for possível e identifique precisamente qualquer parte não validada. Não marque um exemplo como aprovado só porque compila.

## Entrega

Encerre com o caminho do módulo e da skill, número de aulas de código e de referência, validações realizadas e eventuais limitações. Informe um comando para começar.

## Encerramento com commit

Depois de concluir a implementação e as validações, confira o diff e o estado do Git para identificar os arquivos relacionados à tarefa.

- Se o estudante já autorizou o commit desta entrega, faça-o sem pedir confirmação novamente. Inclua apenas os arquivos relacionados ao trabalho, use uma mensagem descritiva em português e informe o hash e o estado final do diretório de trabalho.
- Se ainda não houver autorização, termine a entrega perguntando: **“Quer que eu faça o commit dessas alterações agora?”** A pergunta serve para obter autorização para registrar o trabalho concluído; não interrompa a implementação para fazê-la antecipadamente.
- Se houver falhas de validação ou partes incompletas, informe-as antes de oferecer o commit e não descreva o módulo como totalmente validado.
- Não inclua alterações alheias à tarefa nem faça push sem autorização específica.
