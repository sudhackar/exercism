factorial(0,1):- !. 
factorial(N,F) :-  
   N>0, 
   N1 is N-1, 
   factorial(N1,F1), 
   F is N * F1.

comb(N,K,Out) :-
    X is N-K,
    factorial(N,Fn),factorial(K,Fk),factorial(X,Fx),
    Out is Fn/(Fx*Fk).

row(N,Row) :-
    numlist(0, N, X),
    maplist(comb(N),X,Row).

pascal(0, []):- !.
pascal(Num,Lists) :-
    N is Num-1,
    numlist(0, N, X),
    maplist(row,X,Lists).
