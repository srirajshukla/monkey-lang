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