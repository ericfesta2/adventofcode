from heapq import heappush, heappop

if __name__ == '__main__':
    heap1 = []
    heap2 = []

    with open('input.txt', 'r') as inp:
        for line in inp:
            num1, num2 = line.split('   ')

            heappush(heap1, int(num1))
            heappush(heap2, int(num2))

    total_distance = 0

    while len(heap1) > 0:
        num1 = heappop(heap1)
        total_distance += abs(num1 - heappop(heap2))

    print(total_distance)
