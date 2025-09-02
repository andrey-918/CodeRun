def is_it_correct(forma):
    if '<>' in forma or '<<' in forma or '>>' in forma or '<\>' in forma or chr(47) + chr(47) in forma:
        return False
    mass = []
    new_one = ''
    for i in forma:
        if i == '<':
            if new_one != '':
                return False
        new_one += i
        if i == '>':
            if new_one[0] != '<':
                return False
            mass.append(new_one)
            new_one = ''
    stack = []
    for i in mass:
        if stack != [] and stack[-1][:1] + '/' +  stack[-1][1:] == i:
            stack.pop()
        else:
            stack.append(i)
    if stack != []:
        return False
    return True
def for_answer(XML_fomra):
    alfabet = [_ for _ in range (97, 123)]
    alfabet.append(60)
    alfabet.append(62)
    alfabet.append(47)
    for i in range(len(XML_fomra)):
        for x in alfabet :
            simvol = chr(x)
            if is_it_correct(XML_fomra[:i] + simvol + XML_fomra[i + 1:]):
                return(XML_fomra[:i] + simvol + XML_fomra[i + 1:])
print(for_answer(input()))