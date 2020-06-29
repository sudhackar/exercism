anagram(Str1,Str2) :-
    string_lower(Str1,LStr1),string_lower(Str2,LStr2),
    LStr1 \= LStr2,
    string_chars(LStr1,List1),string_chars(LStr2,List2),
    permutation(List1, List2).

anagram(Word, Options, Matching) :-
    include(anagram(Word),Options,Matching).
