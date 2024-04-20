import functools

def accumulate(doc, sentence):
    return doc + ". " + sentence


def accumulate_first_sentences(sentences, n):
    if n <= 0:
        return ""
    elif len(sentences) == 0:
        return ""
    return functools.reduce(accumulate, sentences[:n]) + "."