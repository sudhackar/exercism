isogram("") :-!.
isogram(S) :-
    string_lower(S, LowerS),
    split_string(LowerS, " -", "", BreakS),
    atomic_list_concat(BreakS, OnlyChars),
    string_codes(OnlyChars, ListChars),
    sort(ListChars, SortedListChars),
    same_length(ListChars, SortedListChars).
