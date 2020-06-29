convert([],0).
convert([X|Xs],Num):-
    convert(Xs,Rest),
    length(Xs,Pow),
    Num is Rest + X * 2 ** (Pow).

binary(Str, Dec) :-
    string_chars(Str, Chars),
    maplist([Char]>>member(Char,['1','0']), Chars),
    maplist(atom_number,Chars,Nums),
    convert(Nums,Dec), !.
