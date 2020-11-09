# Submissão de desafios da plataforma Avança Dev

Resolução do Desafio 1 (Crie o seu microsserviço)

Existem quatro microsserviços desenvolvidos. A ideia da atividade era criar um microsserviço D respondendo à chamada do microsserviço C (Cupons).
Porém, neste caso, criei um microsserviço D respondendo ao microsserviço B. Não creio que seja problema, até porque o que vale é a aplicação do conhecimento.
Os microsserviços são:

  A) Serviço de Front-end;
  B) Mid Server
  C) Serviço de Cupons
  D) Serviço de Cartão de Crédito
  
 Resumidamente, o A cuida da página Web, e envia dados para B. O B, por sua vez, chama C (ou não, pois o usuário pode não querer usar cupom) e D.
 C e D retornam para B "valid" ou "invalid". B retorna para A a resposta final.
