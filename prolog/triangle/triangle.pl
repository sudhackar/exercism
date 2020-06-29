fit(X, Y, Z) :-
    X>0,
    Y>0,
    Z>0,
    msort([X,Y,Z],[A,B,C]), C < A + B.

triangle(X, X, X, "equilateral") :-
    fit(X, X, X),
    !.

triangle(X, Y, Z, "isosceles") :-
    fit(X,Y,Z),
    (sort([X,Y,Z],[_,_]);sort([X,Y,Z],[_])), !.

triangle(X, Y, Z, "scalene") :-
    fit(X,Y,Z), sort([X,Y,Z],[_,_,_]).