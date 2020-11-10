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

Resolução do Desafio 2 (Microsserviços se comunicando utilizando filas)

Os três microsserviços da Aula 2 foram utilizados neste desafio. Não copiei e colei, assisti o vídeo inteiro reproduzindo os mesmo passos no meu computador. Segue em anexo as imagens solicitadas na pasta
"Segundo Desafio".

Resolução do Desafio 3 (Microsserviços utilizando Docker)

Todos os microsserviços do Desafio 2 foram adaptados com a tecnologia Docker. Por falta de tempo, eu não consegui saber porque o microsserviço B não se conecta ao RabbitMQ (connection refused).
Mas acredito que valeu a aplicação do conhecimento, apesar deste pequeno problema. Todos os outros serviços estão funcionando corretamente.

https://hub.docker.com/repository/docker/saulopinedo/microservice-a
https://hub.docker.com/repository/docker/saulopinedo/microservice-b
https://hub.docker.com/repository/docker/saulopinedo/microservice-c
