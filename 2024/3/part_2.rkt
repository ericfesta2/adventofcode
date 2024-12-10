#lang racket

(require "utils.rkt")

(define (filter_disabled split_on_disabled enabled_list ind)
    (match split_on_disabled
        [(cons h t)
            (filter_disabled t (append enabled_list
                (if (= ind 0) ; Every mul() operation before the first don't() should be counted
                    (extract_values_to_mult h)
                    (match (regexp-match-positions #rx"do\\(\\)" h)
                        [#f '()]
                        [(list (cons _ i)) (extract_values_to_mult (substring h i))])))
                (+ ind 1))]
        ['() enabled_list]))

(letrec (
    [inp (open-input-file "input.txt")]
    [split_on_disabled (string-split (port->string inp) #rx"don't\\(\\)")])
    (sum_multiplied (filter_disabled split_on_disabled '() 0) 0))
