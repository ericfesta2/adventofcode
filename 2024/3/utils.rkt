#lang racket

(define (extract_values_to_mult str)
    (regexp-match* #rx"mul\\(([0-9]+),([0-9]+)\\)" #:match-select values str))

; matches is a list of '(<ORIGINAL_STRING> <CAPTURED_NUM_STR_1> <CAPTURED_NUM_STR_2>)
; based on the regex in extract_values_to_mult
(define (sum_multiplied matches sum_so_far)
    (match matches
        [(cons (list _ op1 op2) t) (sum_multiplied t
            (+ sum_so_far (* (string->number op1) (string->number op2))))]
        ['() sum_so_far]))

(provide extract_values_to_mult)
(provide sum_multiplied)
