let fibonacci = fn(x) {
   if (x < 2) {
        return x;
    } else {
        return fibonacci(x - 1) + fibonacci(x - 2);
    }
};

puts("Fibonacci sequence:");
let fibonacci_iter = fn(start, end) {
    let iter = fn(i) {
        if (i < end) {
            puts(fibonacci(i));
            iter(i+1);
        }
    }
    iter(start);
}

fibonacci_iter(1, 10);