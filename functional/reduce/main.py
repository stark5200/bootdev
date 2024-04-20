import functools

def accumulate(doc, sentence):
    return doc + ". " + sentence


def accumulate_first_sentences(sentences, n):
    if n <= 0:
        return ""
    elif len(sentences) == 0:
        return ""
    return functools.reduce(accumulate, sentences[:n]) + "."

#zip    
a = [1, 2, 3]
b = [4, 5, 6]

c = list(zip(a, b))
print(c)
# [(1, 4), (2, 5), (3, 6)]

# 

valid_formats = [
    "docx",
    "pdf",
    "txt",
    "pptx",
    "ppt",
    "md",
]
def pair_document_with_format(doc_names, doc_formats):
    paired = list(zip(doc_names, doc_formats))
    is_valid = lambda x: x[1] in valid_formats
    return filter(is_valid, paired)