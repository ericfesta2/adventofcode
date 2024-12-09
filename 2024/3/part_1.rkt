#lang racket

(define (sum_multiplied matches sum_so_far)
    (match matches
        [(cons h t) (sum_multiplied t
            (+ sum_so_far (* (string->number (cadr h)) (string->number (caddr h)))))]
        ['() sum_so_far]))

(letrec (
    [inp (open-input-file "input.txt")]
    [matches
        (regexp-match* #rx"mul\\(([0-9]+),([0-9]+)\\)" #:match-select values (port->string inp))])
    (sum_multiplied matches 0))
