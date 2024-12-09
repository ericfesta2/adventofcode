FROM ocaml/opam:alpine

WORKDIR /app
COPY . /app

RUN ocamlc -o day_3 part_1.ml

ENTRYPOINT ["/app/day_3"]
