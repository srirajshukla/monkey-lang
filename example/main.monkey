let map = fn(arr, f) {
    let iter = fn(arr, acc) {
        if (len(arr) == 0) {
            acc
        } else {
            iter(rest(arr), push(acc, f(first(arr))))
        }
    };

    iter(arr, [])
};

let a = [1, 2, 3];
puts(join("original array: ", a), "");
let square = fn (x) {x * x};

puts(join("mapped array: ", map(a, square)));

let reduce = fn(arr, init, f) {
    let iter = fn(arr, acc) {
        if (len(arr) == 0) {
            acc
        } else {
            iter(rest(arr), f(acc, first(arr)))
        }
    };

    iter(arr, init);
};

let reduced = reduce(map(a, square), 0, fn(s, v) {s + v});

puts(join("Reduced Value = ", reduced))



let numbers = [1, 2, 3, 4, 5];

let doubled = map(numbers, fn(x) { return x * 2; });
puts("Doubled: ");
puts(doubled);

let sum = reduce(numbers, 0, fn(acc, x) { return acc + x; });
puts("Sum: ");
puts(sum);

let filter = fn (arr, f_fn) {
    let iter = fn(arr, acc) {
        if (len(arr) == 0) {
            acc
        } else {
            if (f_fn(first(arr)) == true) {
                iter(rest(arr), push(acc, first(arr)))
            } else {
                iter(rest(arr), acc)
            }
        }
    }

    iter(arr, [])
}

let evens = filter(numbers, fn(x) { return rem(x,2) == 0; });
puts("Even numbers: ");
puts(evens);

