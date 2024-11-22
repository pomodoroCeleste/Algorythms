def bucket_sort(arr):
    if len(arr) == 0:
        return arr

    # Create n empty buckets
    n = len(arr)
    buckets = [[] for _ in range(n)]

    # Put array elements in different buckets
    for num in arr:
        index = int(n * num)  # Index in bucket
        buckets[index].append(num)

    # Sort individual buckets
    for bucket in buckets:
        bucket.sort()

    # Concatenate all buckets into arr
    sorted_arr = []
    for bucket in buckets:
        sorted_arr.extend(bucket)

    return sorted_arr

# Example usage
if __name__ == "__main__":
    arr = [0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68]
    print("Given array is:", arr)
    sorted_arr = bucket_sort(arr)
    print("Sorted array is:", sorted_arr)