#lang racket

(define (extract_values_to_mult str)
    (regexp-match* #rx"mul\\(([0-9]+),([0-9]+)\\)" #:match-select values str))

(define (sum_multiplied matches sum_so_far)
    (match matches
        [(cons h t) (sum_multiplied t
            (+ sum_so_far (* (string->number (cadr h)) (string->number (caddr h)))))]
        ['() sum_so_far]))

(provide extract_values_to_mult)
(provide sum_multiplied)
