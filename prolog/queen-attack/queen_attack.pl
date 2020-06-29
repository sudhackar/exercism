% Please visit https://exercism.io/tracks/prolog/installation
% for instructions on setting up prolog.
% Visit https://exercism.io/tracks/prolog/tests
% for help running the tests for prolog exercises.

% Replace the goal below with
% your implementation.
create((X, X)) :-
    between(1, 8, X).    

attack((X, _),  (X, _)) :-
    !.
attack((_, Y),  (_, Y)) :-
    !.
attack((X1, Y1),  (X2, Y2)) :-
    abs(X1-X2)=:=abs(Y1-Y2).
