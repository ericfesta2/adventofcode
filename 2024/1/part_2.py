from collections import defaultdict

if __name__ == '__main__':
    list1 = []
    list2_freqs = defaultdict(int) # values of unset keys will be initialized to 0

    with open('input.txt') as inp:
        for line in inp:
            num1, num2 = line.split('   ')

            list1.append(int(num1))
            list2_freqs[int(num2)] += 1

    print(sum([num * list2_freqs[num] for num in list1]))
