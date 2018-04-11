Criação do indice de texto:
```
db.publication.createIndex({ titulo: "text", tecnologia: "text" }, { weights : { titulo: 3, conteudo: 1 }, default_language: "portuguese", name: "_tituloTecID" });
```
Exemplo de busca usando o indice:
```
db.publication.find({$text:{$search:"java"}}).pretty()
```
