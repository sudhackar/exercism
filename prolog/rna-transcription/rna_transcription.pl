convert_one('C','G').
convert_one('G','C').
convert_one('A','U').
convert_one('T','A').
convert(Char,D0,Dn):- convert_one(Char,Dchar), Dn = [Dchar | D0]. 

rna_transcription(Rna, Dna) :-
    string_chars(Rna, RChars),
    maplist([Char]>>member(Char,['C','G','A','T']), RChars),
    foldl(convert,RChars,[],RDChars),
    reverse(RDChars, DChars),
    string_chars(Dna,DChars),!.
