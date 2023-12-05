function to_decimal(bits) {
    value=0
    for (i=0; i < length(bits); i++) {
        bit = int(bits[i])
        if (bit > 0) { value += 2 ^ (msb - i) }
    }
    return value;
}

function find_code(type)
{
    for (c=0; c < codes; ++c) { set[c] = 1 }
    for (b=0; b < msb; b++) {
        ones = 0; total = 0
        for (c=0; c < codes; c++) {
            value = substr(bitstrings[c], b + 1, 1)
            if (set[c] == 1) {
                total += 1
                column[c] = value
                if (value == 1) { ones += 1 }
            }
        }
        if (type == "O2") { target = ((ones >= (total / 2)) ? 1 : 0) }
        else              { target = ((ones < (total / 2)) ? 1 : 0) }
        count = 0
        for (c=0; c < codes; c++) {
            if (set[c] == 1 && column[c] != target) { set[c] = 0 }
            if (set[c] == 1) { count += 1; last = c}
        }
        if (count == 1) { return bitstrings[last] }
        delete column
    }
}

{
    bitstrings[NR-1] = $0
    codes += 1
}

END {
    msb = length(bitstrings[0])
    split(find_code("O2"), o2, "")
    split(find_code("CO2"), co2, "")
    print to_decimal(o2) * to_decimal(co2)
}


