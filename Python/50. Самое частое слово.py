def get_answer(a):
    dictionary = dict()
    answer = -1
    answer_word = 'œ∑´®†¥¨ƒˆøß∂ƒ'
    for i in a:
        if i not in dictionary:
            dictionary[i] = 0
        dictionary[i] += 1
        if dictionary[i] > answer or (dictionary[i] == answer and answer_word > i):
            answer = dictionary[i]
            answer_word = i
    return answer_word

    
with open('input.txt', 'r') as f:
    a = f.read().split()
    print(get_answer(a))