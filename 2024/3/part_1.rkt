#lang racket

(require "utils.rkt")

(letrec (
    [inp (open-input-file "input.txt")]
    [matches (extract_values_to_mult (port->string inp))])
    (sum_multiplied matches 0))
