def get_max(arr):
    # Returns the maximum value in the array
    return max(arr)

def counting_sort(arr, exp):
    n = len(arr)
    output = [0] * n  # output array
    count = [0] * 10  # count array to store the count of occurrences of digits

    # Store count of occurrences in count[]
    for i in range(n):
        index = (arr[i] // exp) % 10
        count[index] += 1

    # Change count[i] so that count[i] contains the actual position of this digit in output[]
    for i in range(1, 10):
        count[i] += count[i - 1]

    # Build the output array
    i = n - 1
    while i >= 0:
        index = (arr[i] // exp) % 10
        output[count[index] - 1] = arr[i]
        count[index] -= 1
        i -= 1

    # Copy the output array to arr[], so that arr[] contains sorted numbers according to the current digit
    for i in range(n):
        arr[i] = output[i]

def radix_sort(arr):
    # Find the maximum number to know the number of digits
    max_val = get_max(arr)

    # Do counting sort for every digit. Note that instead of passing the digit number, exp is passed.
    # exp is 10^i where i is the current digit number
    exp = 1
    while max_val // exp > 0:
        counting_sort(arr, exp)
        exp *= 10

# Example array
arr = [170, 45, 75, 90, 802, 24, 2, 66]
# Sort the array using radix_sort
radix_sort(arr)
# Print the sorted array
print("Sorted array is:", arr)